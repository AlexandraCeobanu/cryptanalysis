package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

var epsilon = 0.005

func readCryptotext() string {
	cryptotext, error := os.ReadFile("cryptotextVigenere.txt")
	if error != nil {
		panic(error)
	}

	return string(cryptotext)
}
func computeIC(text string) float64 {

	letterFreq := make(map[rune]int)
	for letter := 'A'; letter <= 'Z'; letter++ {
		letterFreq[letter] = 0
	}

	for _, letter := range text {
		letterFreq[letter] = letterFreq[letter] + 1
	}
	var sum float64 = 0
	textLength := len(text)
	for _, letter := range letterFreq {

		if textLength > 1 {
			firstP := float64(letter) / float64(textLength)
			secondP := float64(letter-1) / float64(textLength-1)
			sum += firstP * secondP
		}
	}
	return sum
}
func extractSubtext(text string, m int, j int) string {

	var builder strings.Builder
	for i := j; i < len(text); i = i + m {
		builder.WriteRune(rune(text[i]))
	}
	return builder.String()

}
func findKeyLength(cryptotext string) int {

	m := 1

	for {
		m = m + 1
		// lengthFound := true
		var sum float64
		for i := 1; i <= m; i++ {
			subtext := extractSubtext(cryptotext, m, i-1)
			ic := float64(computeIC(subtext))
			sum = sum + ic
			// if math.Abs(ic-0.065) > epsilon {
			// 	lengthFound = false
			// 	break
			// }
		}
		avg := float64(sum / (float64(m)))
		if math.Abs(avg-0.065) < epsilon {
			// lengthFound = false
			break
		}
		// if lengthFound {
		// 	break
		// }
	}
	return m
}
func computeMIC(beta string) float64 {

	var realFreq = []float64{0.08167, 0.01492, 0.02782, 0.04253, 0.12702, 0.02228, 0.02015, 0.06094, 0.06966, 0.00153, 0.00772, 0.04025, 0.02406, 0.06749, 0.07507, 0.01929, 0.00095, 0.05987, 0.06327, 0.09056, 0.02758, 0.00978, 0.02360, 0.00150, 0.01974, 0.00074}

	letterFreq := make(map[rune]int)
	for letter := 'A'; letter <= 'Z'; letter++ {
		letterFreq[letter] = 0
	}

	for _, letter := range beta {
		letterFreq[letter] = letterFreq[letter] + 1
	}
	var sum float64 = 0
	betaLength := len(beta)
	for letter, freq := range letterFreq {

		firstP := float64(freq) / float64(betaLength)
		sum += firstP * float64(realFreq[letter-'A'])

	}
	return sum
}
func shift(text string, s int) string {
	var builder strings.Builder
	for _, letter := range text {

		newLetter := int((int(letter-'A') + s) % 26)
		builder.WriteRune(rune('A' + newLetter))
	}
	return builder.String()
}

func findKey(cryptotext string, m int) string {
	var keyBuilder strings.Builder

	for j := 1; j <= m; j++ {
		s := -1
		for {
			s = s + 1
			subtext := extractSubtext(cryptotext, m, j-1)
			newText := shift(subtext, s)
			mic := computeMIC(newText)
			if math.Abs(mic-0.065) < epsilon {
				break
			}

		}
		keyBuilder.WriteRune(rune('A' + ((26 - s) % 26)))
	}
	return keyBuilder.String()

}
func main() {

	// var alphabet = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	cryptotext := readCryptotext()
	// cryptotext := "LYIOPVIRNBZ"
	keyLength := findKeyLength(cryptotext)
	fmt.Println(keyLength)
	key := findKey(cryptotext, keyLength)
	fmt.Println(key)
}
