package helpers

import (
	"math/rand"
	"os"
	"strings"
)

func EnforceHttp(url string) string {
	if url[:4] != "http" {
		return "http://" + url
	}

	return url
}

func CheckDomainError(url string) bool {

	if url == os.Getenv("DOMAIN") {
		return false
	}

	newUrl := strings.Replace(url, "http://", "", 1)
	newUrl = strings.Replace(newUrl, "https://", "", 1)
	newUrl = strings.Replace(newUrl, "www.", "", 1)
	newUrl = strings.Split(newUrl, "/")[0]

	return newUrl != os.Getenv("DOMAIN")
}

func GenerateLinkId(length int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
