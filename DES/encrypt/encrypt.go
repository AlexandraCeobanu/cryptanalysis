package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
)

func readPlaintext() string {
	plaintext, error := os.ReadFile("DES/encrypt/plaintext.txt")
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

func toBytes(plaintext string) []byte {
	// var bitString strings.Builder
	bitString, err := hex.DecodeString(plaintext)
	if err != nil {
		log.Fatal("Error decoding hex:", err)
	}
	// for _, character := range plaintext {
	// 	bitString.WriteString(fmt.Sprintf("%08b", character))
	// }
	return bitString
}

func initialPermutation(bitString string, p [][]int) string {
	var stringAfterP strings.Builder
	for i, row := range p {
		for j, _ := range row {

			stringAfterP.WriteRune(rune(bitString[p[i][j]-1]))
		}
	}
	return stringAfterP.String()
}

// func createBlocks(bitString string) []string {
// 	var blocks []string
// 	for i := 0; i < len(bitString); i += 64 {
// 		end := i + 64
// 		if end > len(bitString) {
// 			end = len(bitString)
// 		}
// 		block := string(bitString[i:end])

// 		blocks = append(blocks, block)
// 	}
// 	return blocks

// }

// func encrypt(plaintext string, key []string) string {

// }

func toBits(arrayOfBytes []byte) string {

	var stringOfBits strings.Builder
	for _, byteElement := range arrayOfBytes {
		stringOfBits.WriteString(fmt.Sprintf("%08b", byteElement))

	}
	return stringOfBits.String()
}

func main() {
	// var key = []string{"1", "3", "3", "4", "5", "7", "7", "9", "9", "B", "B", "C", "D", "F", "F", "1"}
	initialP := [][]int{
		{58, 50, 42, 34, 26, 18, 10, 2},
		{60, 52, 44, 36, 28, 20, 12, 4},
		{62, 54, 46, 38, 30, 22, 14, 6},
		{64, 56, 48, 40, 32, 24, 16, 8},
		{57, 49, 41, 33, 25, 17, 9, 1},
		{59, 51, 43, 35, 27, 19, 11, 3},
		{61, 53, 45, 37, 29, 21, 13, 5},
		{63, 55, 47, 39, 31, 23, 15, 7},
	}

	plaintext := readPlaintext()
	arrayOfBytes := toBytes(plaintext)
	fmt.Println(arrayOfBytes)

	stringOfBits := toBits(arrayOfBytes)
	fmt.Println(stringOfBits)

	bitsAfterP := initialPermutation(stringOfBits, initialP)
	fmt.Println(bitsAfterP)

	L0 := bitsAfterP[0:32]
	R0 := bitsAfterP[32:64]

	fmt.Println("L0 = ", L0)
	fmt.Println("L1 = ", R0)

	// blocks := createBlocks(stringOfBits)
	// for _, block := range blocks {
	// 	fmt.Println(block)
	// 	blockAfterP := initialPermutation(stringOfBits, initialP)
	// 	fmt.Println(blockAfterP)
	// 	fmt.Println("------------------------------------------------")
	// }

	// cryptotext := encrypt(newPlaintext, key)
	// writeCryptotext(cryptotext)
}
