package bitmexws

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

const (
	defaultExpiry = 5
)

// Create the authentication signature for this request
func generateAuthSignature(apiSecret string, expires int64) string {
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write([]byte(fmt.Sprintf("GET/realtime%d", expires)))
	return hex.EncodeToString(h.Sum(nil))
}

// Generate authentication object and sends it on channel `send`
func Authenticate(send chan *Command, apiKey, apiSecret string) {
	expires := time.Now().Unix() + defaultExpiry
	c := CreateCommand("authKeyExpires", apiKey, expires, generateAuthSignature(apiSecret, expires))
	send <- c
}
