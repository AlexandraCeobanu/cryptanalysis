package main

import (
	"fmt"
	"os"
	"strings"
)

type letter struct {
	Key   string
	Value float32
}

var sortedFrequencies []letter

func readPlaintext() string {
	plaintext, error := os.ReadFile("plaintext.txt")
	if error != nil {
		panic(error)
	}
	return string(plaintext)
}
func writeCryptotext(cryptotext string) {

	file, error := os.Create("cryptotext.txt")
	if error != nil {
		fmt.Println("Failed to create file:", error)
		return
	}
	defer file.Close()

	_, error = file.WriteString(cryptotext)
	if error != nil {
		fmt.Println("Failed to write to file:", error)
		return
	}

}
func processPlaintext(plaintext string) string {

	lowerString := strings.ToLower(plaintext)
	var builder strings.Builder
	for _, character := range lowerString {
		if character >= 'a' && character <= 'z' {
			builder.WriteRune(character)
		}
	}
	newString := builder.String()
	return newString
}
func encrypt(plaintext string, key []string) string {

	var cryptotext string
	var builder strings.Builder
	for _, character := range plaintext {
		encryptedCharacter := key[character-'a']
		builder.WriteString(encryptedCharacter)
	}
	cryptotext = builder.String()
	return cryptotext
}

func main() {
	var key = []string{"O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N"}
	plaintext := readPlaintext()
	newPlaintext := processPlaintext(plaintext)
	//fmt.Println(newPlaintext)
	cryptotext := encrypt(newPlaintext, key)
	writeCryptotext(cryptotext)
}
