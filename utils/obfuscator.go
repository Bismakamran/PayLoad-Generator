package utils

import (
	"math/rand"
	"strings"
	"time"
)

// Add random spacing between characters (e.g., <s c r i p t>)
func ObfuscateRandomSpacing(input string) string {
	var result strings.Builder
	for i, c := range input {
		result.WriteRune(c)
		if i != len(input)-1 {
			result.WriteString(" ")
		}
	}
	return result.String()
}

// Insert HTML comments between characters (e.g., <script><!-- -->alert(1)</script>)
func ObfuscateWithComments(input string) string {
	return strings.ReplaceAll(input, ">", "><!-- -->")
}

// Flip character cases randomly (e.g., <ScRiPt>)
func ObfuscateCaseFlip(input string) string {
	rand.Seed(time.Now().UnixNano())
	var result strings.Builder
	for _, c := range input {
		if rand.Intn(2) == 0 {
			result.WriteString(strings.ToLower(string(c)))
		} else {
			result.WriteString(strings.ToUpper(string(c)))
		}
	}
	return result.String()
}

// Replace characters with homoglyphs (basic version)
func ObfuscateHomoglyphs(input string) string {
	homoglyphs := map[rune]string{
		'a': "а", // Cyrillic a
		'e': "е", // Cyrillic e
		'i': "і", // Ukrainian i
		'o': "о", // Cyrillic o
		'c': "с", // Cyrillic s
		's': "ѕ", // Cyrillic s (different)
		'l': "ⅼ", // Roman numeral 50
	}
	var result strings.Builder
	for _, c := range input {
		if replacement, ok := homoglyphs[c]; ok {
			result.WriteString(replacement)
		} else {
			result.WriteRune(c)
		}
	}
	return result.String()
}
