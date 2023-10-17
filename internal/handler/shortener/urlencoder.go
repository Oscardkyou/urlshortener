package shortenpkg

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// UrlShortener - структура, представляющая собой сократитель URL, который хранит URL в текстовом файле
type UrlShortener struct {
	store string // Путь к текстовому файлу, где хранятся пары короткий URL - длинный URL
}

// New создает новый экземпляр UrlShortener
func New() *UrlShortener {
	rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел
	return &UrlShortener{store: "/path/to/your/storage.txt"}
}

// generateRandomString генерирует случайную строку заданной длины
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// Encode принимает длинный URL и возвращает короткий URL, сохраняя эту пару в файле
func (us *UrlShortener) Encode(longUrl string) string {
	shortUrl := generateRandomString(6)
	// Проверяем на уникальность короткого URL
	for us.loadFromStorage(shortUrl) != "" {
		shortUrl = generateRandomString(6)
	}
	us.saveToStorage(shortUrl, longUrl)
	return shortUrl
}

// Decode принимает короткий URL и возвращает соответствующий длинный URL из файла
func (us *UrlShortener) Decode(shortUrl string) (string, bool) {
	longUrl := us.loadFromStorage(shortUrl)
	return longUrl, longUrl != ""
}

// saveToStorage сохраняет пару короткий URL - длинный URL в файл
func (us *UrlShortener) saveToStorage(shortUrl string, longUrl string) {
	file, err := os.OpenFile(us.store, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := file.WriteString(shortUrl + "::::" + longUrl + "\n"); err != nil {
		log.Fatal(err)
	}
}

// loadFromStorage загружает длинный URL из файла по заданному короткому URL
func (us *UrlShortener) loadFromStorage(shortUrl string) string {
	file, err := os.Open(us.store)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "::::")
		if parts[0] == shortUrl {
			return parts[1]
		}
	}
	return ""
}
