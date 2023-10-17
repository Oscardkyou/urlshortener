
package shortenpkg

import (
    "database/sql"
    "net/http"
    "log"
    _ "github.com/lib/pq"
    "GoUrlShortener/internal/shortener"
)

var db *sql.DB

func main() {
    var err error
    db, err = sql.Open("postgres", "user=username password=password dbname=mydb sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }

    http.HandleFunc("/api/shorten", shortener.ShortenHandler)
    http.HandleFunc("/api/resolve", shortener.ResolveHandler)

    log.Fatal(http.ListenAndServe(":8080", nil))
}
