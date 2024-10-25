package main

import (
	"fmt"
	"os"
	"strings"
)

func readPlaintext() string {
	plaintext, error := os.ReadFile("plaintextVigenere.txt")
	if error != nil {
		panic(error)
	}
	return string(plaintext)
}
func writeCryptotext(cryptotext string) {

	file, error := os.Create("cryptotextVigenere.txt")
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

	lowerString := strings.ToUpper(plaintext)
	var builder strings.Builder
	for _, character := range lowerString {
		if character >= 'A' && character <= 'Z' {
			builder.WriteRune(character)
		}
	}
	newString := builder.String()
	return newString
}
func encrypt(plaintext string, key []rune) string {

	var cryptotext string
	var builder strings.Builder
	for index, character := range plaintext {

		key_i := key[int(index%len(key))]
		encryptedCharacter := ((character - 'A') + (key_i - 'A')) % 26
		builder.WriteRune('A' + encryptedCharacter)
	}
	cryptotext = builder.String()
	return cryptotext
}

func main() {
	// var alphabet = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	//                        0    1    2    3    4    5    6    7    8    9    10   11   12   13   14   15   16   17   18   19   20   21
	var key = []rune{'A', 'B', 'A', 'B', 'A', 'B', 'B'}
	// var key = []rune{'A', 'B', 'A', 'B', 'A', 'B', 'A', 'A'} caz pe care nu merge
	plaintext := readPlaintext()
	newPlaintext := processPlaintext(plaintext)
	//fmt.Println(newPlaintext)
	cryptotext := encrypt(newPlaintext, key)
	writeCryptotext(cryptotext)
}
