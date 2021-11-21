package gameanalytics

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
)

// const HOST = "https://api.gameanalytics.com/v2/"
const HOST = "https://api.gameanalytics.com/"

func authHash(body string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(body))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func headers(body string, secret string) http.Header {
	return http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{authHash(body, secret)},
	}
}
