package main

import (
	"log"
	"net/http"
	"urlshortener/api/shortenpkg"
	// убедитесь, что путь к пакету верный
	// "urlshortener/shortener" // если вам потребуется этот импорт
)

func main() {
	// Так как обработчики теперь находятся в пакете shortenpkg, нет необходимости создавать экземпляр MemoryStorage
	// и shortenerService в этом файле.

	// Обработчики URL
	http.HandleFunc("/api/shorten", shortenpkg.ShortenHandler)
	http.HandleFunc("/api/resolve", shortenpkg.ResolveHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
