
package shortenpkg

import (
    "database/sql"
    "net/http"
    _ "github.com/lib/pq"
)

var db *sql.DB

func Set(shortKey, url string) error {
    _, err := db.Exec("INSERT INTO url_shortener (shortKey, url) VALUES ($1, $2)", shortKey, url)
    return err
}

func Get(shortKey string) (string, error) {
    var url string
    err := db.QueryRow("SELECT url FROM url_shortener WHERE shortKey = $1", shortKey).Scan(&url)
    if err != nil {
        return "", err
    }
    return url, nil
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
        return
    }
    longURL := r.FormValue("url")
    if longURL == "" {
        http.Error(w, "URL is required", http.StatusBadRequest)
        return
    }
    shortKey := GenerateShortenedKey()
    err := Set(shortKey, longURL)
    if err != nil {
        http.Error(w, "Failed to save URL", http.StatusInternalServerError)
        return
    }
    w.Write([]byte(shortKey))
}

func ResolveHandler(w http.ResponseWriter, r *http.Request) {
    shortKey := r.FormValue("key")
    if shortKey == "" {
        http.Error(w, "Key is required", http.StatusBadRequest)
        return
    }
    longURL, err := Get(shortKey)
    if err != nil {
        http.Error(w, "Failed to get URL", http.StatusInternalServerError)
        return
    }
    w.Write([]byte(longURL))
}
