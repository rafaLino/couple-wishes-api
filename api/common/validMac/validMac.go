package validMac

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func ValidMAC(body, secret []byte, header string) bool {
	signature := strings.Split(header, "=")[1]
	sig, _ := hex.DecodeString(signature)
	mac := hmac.New(sha256.New, secret)
	mac.Write(body)

	return hmac.Equal(sig, mac.Sum(nil))
}
