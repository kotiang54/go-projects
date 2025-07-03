package intermediate

import (
	"encoding/base64"
	"fmt"
)

func main() {
	/*
		Text Encoding (ASCII, UTF-8, UTF-16)
		Data Encoding (Base64, URL Encoding)
		File Encoding (Binary Text)
	*/
	// Base64 - converts binary data into a text representation using a set of 64-ASCII characters

	data := []byte("He~llo, Base64 Encoding")

	// Encode
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println(encoded)

	// Decode from base64
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error in decoding", err)
		return
	}
	fmt.Println("Decoded:", decoded)
	fmt.Println("Decoded - String:", string(decoded))

	// URL safe, avoid '/' and '+'
	urlSafeEncoded := base64.URLEncoding.EncodeToString(data)
	fmt.Println("URL Safe Encode:", urlSafeEncoded)
}
