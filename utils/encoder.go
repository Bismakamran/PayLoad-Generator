package utils

import (
	"encoding/base64"
	"fmt"
	"net/url"
)

// Base64 Encoding
func EncodeBase64(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

// URL Encoding
func EncodeURL(input string) string {
	return url.QueryEscape(input)
}

// Hex Encoding
func EncodeHex(input string) string {
	hex := ""
	for _, c := range input {
		hex += fmt.Sprintf("\\x%02x", c)
	}
	return hex
}

// Unicode Encoding
func EncodeUnicode(input string) string {
	unicode := ""
	for _, c := range input {
		unicode += fmt.Sprintf("\\u%04x", c)
	}
	return unicode
}
