package main

import (
	"fmt"
	"payload-generator-go/utils"
)

func main() {
	payload := "<script>alert(1)</script>"

	fmt.Println("Base64:", utils.EncodeBase64(payload))
	fmt.Println("URL Encoded:", utils.EncodeURL(payload))
	fmt.Println("Hex Encoded:", utils.EncodeHex(payload))
	fmt.Println("Unicode Encoded:", utils.EncodeUnicode(payload))

	fmt.Println("Random Spacing:", utils.ObfuscateRandomSpacing(payload))
	fmt.Println("With Comments:", utils.ObfuscateWithComments(payload))
	fmt.Println("Case Flipped:", utils.ObfuscateCaseFlip(payload))
	fmt.Println("Homoglyphs:", utils.ObfuscateHomoglyphs(payload))
}
