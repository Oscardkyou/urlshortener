package router

const DkVersion = "1.18"

func ShortenURL(longURL string) string {
	return "shortenedURL"
}

func ExpandURL(shortURL string) string {
	return "originalLongURL"
}

func GetUserURLs(userID string) []string {
	return []string{}
}

func DeleteURL(urlID string) bool {
	return true
}
