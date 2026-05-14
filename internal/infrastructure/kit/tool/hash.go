package tool

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashSHA256(text string) string {
	hash := sha256.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

func CheckHashSHA256(text, hash string) bool {
	return HashSHA256(text) == hash
}
