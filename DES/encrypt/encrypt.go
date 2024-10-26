package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
)

var Etable = [][]int{
	{32, 1, 2, 3, 4, 5},
	{4, 5, 6, 7, 8, 9},
	{8, 9, 10, 11, 12, 13},
	{12, 13, 14, 15, 16, 17},
	{16, 17, 18, 19, 20, 21},
	{20, 21, 22, 23, 24, 25},
	{24, 25, 26, 27, 28, 29},
	{28, 29, 30, 31, 32, 1},
}

var PC_1 = [][]int{
	{57, 49, 41, 33, 25, 17, 9},
	{1, 58, 50, 42, 34, 26, 18},
	{10, 2, 59, 51, 43, 35, 27},
	{19, 11, 3, 60, 52, 44, 36},
	{63, 55, 47, 39, 31, 23, 15},
	{7, 62, 54, 46, 38, 30, 22},
	{14, 6, 61, 53, 45, 37, 29},
	{21, 13, 5, 28, 20, 12, 4},
}

var PC_2 = [][]int{
	{14, 17, 11, 24, 1, 5},
	{3, 28, 15, 6, 21, 10},
	{23, 19, 12, 4, 26, 8},
	{16, 7, 27, 20, 13, 2},
	{41, 52, 31, 37, 47, 55},
	{30, 40, 51, 45, 33, 48},
	{44, 49, 39, 56, 34, 53},
	{46, 42, 50, 36, 29, 32},
}

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

func toBits(arrayOfBytes []byte) string {

	var stringOfBits strings.Builder
	for _, byteElement := range arrayOfBytes {
		stringOfBits.WriteString(fmt.Sprintf("%08b", byteElement))

	}
	return stringOfBits.String()
}

func deleteBits(bitString string) string {

	var newString strings.Builder
	for ind := 0; ind < len(bitString); ind = ind + 8 {

		substring := string(bitString[ind : ind+8-1])
		newString.WriteString(substring)
	}
	return newString.String()
}
func E(A string) string {
	var extendedA strings.Builder
	for i, row := range Etable {
		for j, _ := range row {

			extendedA.WriteRune(rune(A[Etable[i][j]-1]))
		}
	}
	return extendedA.String()
}
func f(R string) {
	expendedR := E(R)
	fmt.Println("Expended R: ")
	fmt.Println(expendedR)
}
func encrypt(L0 string, R0 string) {

	f(R0)
	// previousL := L0
	// previousR := R0
	// for round := 1; round <= 16; round++ {

	// 	currentL := previousR
	// 	currentR := f(previousR)

	// }

}
func main() {
	key := "133457799BBCDFF1"
	// 00010010011010010101101111001001101101111011011111111000
	// 00010010011010010101101111001001101101111011011111111000
	keyArrayOfBytes := toBytes(key)
	keyStringOfBits := toBits(keyArrayOfBytes)
	keyStringOfBits = deleteBits(keyStringOfBits)
	fmt.Println("key bits: ", keyStringOfBits)
	fmt.Println("-----------------------------------------------")
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
	fmt.Println("R0 = ", R0)

	encrypt(L0, R0)
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
