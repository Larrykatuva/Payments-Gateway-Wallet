package services

import (
	"encoding/base64"
	"regexp"
	"strings"
	"time"
)

func generateTransactionNumber() string {
	ticks := time.Now().UnixNano()
	bytes := make([]byte, 8)
	for i := 0; i < 8; i++ {
		bytes[i] = byte(ticks >> (i * 8))
	}
	id := base64.StdEncoding.EncodeToString(bytes)
	id = strings.ReplaceAll(id, "+", " ")
	id = strings.ReplaceAll(id, "/", " ")
	id = strings.TrimRight(id, "=")
	return id
}

func removeSpecialChars(refNo string) string {
	reg := regexp.MustCompile("[!@#$%\\^&*\\(\\)_+=\\/\\\\{\\}\\[\\]\\|/:;/\"'<>,.\\?\\~`;]")
	return reg.ReplaceAllString(refNo, "")
}

func GenerateRrn() string {
	return removeSpecialChars(generateTransactionNumber())
}
