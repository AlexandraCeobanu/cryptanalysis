package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var initialPInv = [][]int{
	{40, 8, 48, 16, 56, 24, 64, 32},
	{39, 7, 47, 15, 55, 23, 63, 31},
	{38, 6, 46, 14, 54, 22, 62, 30},
	{37, 5, 45, 13, 53, 21, 61, 29},
	{36, 4, 44, 12, 52, 20, 60, 28},
	{35, 3, 43, 11, 51, 19, 59, 27},
	{34, 2, 42, 10, 50, 18, 58, 26},
	{33, 1, 41, 9, 49, 17, 57, 25},
}

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

var P = [][]int{
	{16, 7, 20, 21},
	{29, 12, 28, 17},
	{1, 15, 23, 26},
	{5, 18, 31, 10},
	{2, 8, 24, 14},
	{32, 27, 3, 9},
	{19, 13, 30, 6},
	{22, 11, 4, 25},
}

var SBoxes = [][][]int{
	{
		{14, 4, 13, 1, 2, 15, 11, 8, 3, 10, 6, 12, 5, 9, 0, 7},
		{0, 15, 7, 4, 14, 2, 13, 1, 10, 6, 12, 11, 9, 5, 3, 8},
		{4, 1, 14, 8, 13, 6, 2, 11, 15, 12, 9, 7, 3, 10, 5, 0},
		{15, 12, 8, 2, 4, 9, 1, 7, 5, 11, 3, 14, 10, 0, 6, 13},
	},
	{
		{15, 1, 8, 14, 6, 11, 3, 4, 9, 7, 2, 13, 12, 0, 5, 10},
		{3, 13, 4, 7, 15, 2, 8, 14, 12, 0, 1, 10, 6, 9, 11, 5},
		{0, 14, 7, 11, 10, 4, 13, 1, 5, 8, 12, 6, 9, 3, 2, 15},
		{13, 8, 10, 1, 3, 15, 4, 2, 11, 6, 7, 12, 0, 5, 14, 9},
	},
	{
		{10, 0, 9, 14, 6, 3, 15, 5, 1, 13, 12, 7, 11, 4, 2, 8},
		{13, 7, 0, 9, 3, 4, 6, 10, 2, 8, 5, 14, 12, 11, 15, 1},
		{13, 6, 4, 9, 8, 15, 3, 0, 11, 1, 2, 12, 5, 10, 14, 7},
		{1, 10, 13, 0, 6, 9, 8, 7, 4, 15, 14, 3, 11, 5, 2, 12},
	},
	{
		{7, 13, 14, 3, 0, 6, 9, 10, 1, 2, 8, 5, 11, 12, 4, 15},
		{13, 8, 11, 5, 6, 15, 0, 3, 4, 7, 2, 12, 1, 10, 14, 9},
		{10, 6, 9, 0, 12, 11, 7, 13, 15, 1, 3, 14, 5, 2, 8, 4},
		{3, 15, 0, 6, 10, 1, 13, 8, 9, 4, 5, 11, 12, 7, 2, 14},
	},
	{
		{2, 12, 4, 1, 7, 10, 11, 6, 8, 5, 3, 15, 13, 0, 14, 9},
		{14, 11, 2, 12, 4, 7, 13, 1, 5, 0, 15, 10, 3, 9, 8, 6},
		{4, 2, 1, 11, 10, 13, 7, 8, 15, 9, 12, 5, 6, 3, 0, 14},
		{11, 8, 12, 7, 1, 14, 2, 13, 6, 15, 0, 9, 10, 4, 5, 3},
	},
	{
		{12, 1, 10, 15, 9, 2, 6, 8, 0, 13, 3, 4, 14, 7, 5, 11},
		{10, 15, 4, 2, 7, 12, 9, 5, 6, 1, 13, 14, 0, 11, 3, 8},
		{9, 14, 15, 5, 2, 8, 12, 3, 7, 0, 4, 10, 1, 13, 11, 6},
		{4, 3, 2, 12, 9, 5, 15, 10, 11, 14, 1, 7, 6, 0, 8, 13},
	},
	{
		{4, 11, 2, 14, 15, 0, 8, 13, 3, 12, 9, 7, 5, 10, 6, 1},
		{13, 0, 11, 7, 4, 9, 1, 10, 14, 3, 5, 12, 2, 15, 8, 6},
		{1, 4, 11, 13, 12, 3, 7, 14, 10, 15, 6, 8, 0, 5, 9, 2},
		{6, 11, 13, 8, 1, 4, 10, 7, 9, 5, 0, 15, 14, 2, 3, 12},
	},
	{
		{13, 2, 8, 4, 6, 15, 11, 1, 10, 9, 3, 14, 5, 0, 12, 7},
		{1, 15, 13, 8, 10, 3, 7, 4, 12, 5, 6, 11, 0, 14, 9, 2},
		{7, 11, 4, 1, 9, 12, 14, 2, 0, 6, 10, 13, 15, 3, 5, 8},
		{2, 1, 14, 7, 4, 10, 8, 13, 15, 12, 9, 0, 3, 5, 6, 11},
	},
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

func initialPermutationInv(bitString string) string {
	var stringAfterP strings.Builder
	for i, row := range initialPInv {
		for j, _ := range row {

			stringAfterP.WriteRune(rune(bitString[initialPInv[i][j]-1]))
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

// func deleteBits(bitString string) string {

// 	var newString strings.Builder
// 	for ind := 0; ind < len(bitString); ind = ind + 8 {

// 		substring := string(bitString[ind : ind+8-1])
// 		newString.WriteString(substring)
// 	}
// 	return newString.String()
// }

func E(A string) string {
	var extendedA strings.Builder
	for i, row := range Etable {
		for j, _ := range row {

			extendedA.WriteRune(rune(A[Etable[i][j]-1]))
		}
	}
	return extendedA.String()
}
func computeS_i(B_i string, i int) int {

	binaryRow := string(B_i[0]) + string(B_i[5])
	row, _ := strconv.ParseInt(binaryRow, 2, 64)

	binaryCol := string(B_i[1]) + string(B_i[2]) + string(B_i[3]) + string(B_i[4])
	col, _ := strconv.ParseInt(binaryCol, 2, 64)

	// fmt.Println("Row  ", row)
	// fmt.Println("Col ", col)
	value := SBoxes[i][row][col]
	// fmt.Println("Value ", value)
	return value

}
func f(R string, key string) string {
	expendedR := E(R)
	fmt.Println("Expended R: ")
	fmt.Println(expendedR)
	fmt.Println("---------------------------------------")

	var B strings.Builder
	var C strings.Builder
	for i := 0; i < len(expendedR); i++ {
		if expendedR[i] == key[i] {
			B.WriteString("0")
		} else {
			B.WriteString("1")
		}
	}

	stringB := B.String()
	fmt.Println("XOR R si K: ")
	fmt.Println(stringB)
	fmt.Println("---------------------------------------")

	for i := 0; i < len(stringB); i = i + 6 {

		value := computeS_i(stringB[i:i+6], i/6)
		C_i := fmt.Sprintf("%04b", value)
		C.WriteString(C_i)
	}

	fmt.Println("S_box: ")
	fmt.Println(C.String())
	fmt.Println("---------------------------------------")

	permutedC := permutationP(C.String())

	fmt.Println("f: ")
	fmt.Println(permutedC)
	fmt.Println("---------------------------------------")

	return permutedC

}
func permutationP(C string) string {

	var newC strings.Builder
	for i, row := range P {
		for j, _ := range row {

			newC.WriteRune(rune(C[P[i][j]-1]))
		}
	}
	return newC.String()
}
func permutationPC1(k string) string {

	var newK strings.Builder
	for i, row := range PC_1 {
		for j, _ := range row {

			newK.WriteRune(rune(k[PC_1[i][j]-1]))
		}
	}
	return newK.String()
}
func permutationPC2(C string, D string) string {

	var newK strings.Builder
	var k = C + D
	for i, row := range PC_2 {
		for j, _ := range row {

			newK.WriteRune(rune(k[PC_2[i][j]-1]))
		}
	}
	return newK.String()

}
func LS(A string, i int) string {
	var newA strings.Builder
	if i == 1 || i == 2 || i == 9 || i == 16 {
		newA.WriteString(A[1:])
		newA.WriteString(A[0:1])
	} else {
		newA.WriteString(A[2:])
		newA.WriteString(A[0:2])
	}
	return newA.String()

}
func XOR(a string, b string) string {
	var B strings.Builder

	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			B.WriteString("0")
		} else {
			B.WriteString("1")
		}
	}
	return B.String()
}
func encrypt(L0 string, R0 string, key string) string {

	previousL := L0
	previousR := R0

	var currentR string
	var currentL string
	initialKeyPermutation := permutationPC1(key)
	// fmt.Println("permutare initiala pe cheie:  ", initialKeyPermutation)
	// fmt.Println("----------------------------------------------------")
	previousC := initialKeyPermutation[0:28]
	previousD := initialKeyPermutation[28:56]

	for round := 1; round <= 16; round++ {

		fmt.Println("Round ", round)
		fmt.Println("-------------------------------------------------------------------------")
		currentC := LS(previousC, round)
		currentD := LS(previousD, round)
		K_i := permutationPC2(currentC, currentD)
		fmt.Println("Key :  ")
		fmt.Println(K_i)
		fmt.Println("----------------------------------")
		previousC = currentC
		previousD = currentD

		currentL = previousR
		currentR = XOR(f(previousR, K_i), previousL)

		fmt.Println("L = R: ")
		fmt.Println(currentR)
		fmt.Println("---------------------------------------")

		previousL = currentL
		previousR = currentR
	}

	fmt.Println("R16  ", currentR)
	fmt.Println("L16  ", currentL)

	inv := string(currentR) + string(currentL)
	invPermuted := initialPermutationInv(inv)

	fmt.Println("Binary : ", invPermuted)

	// var cryptotext strings.Builder
	// for i := 0; i < len(invPermuted); i += 4 {
	// 	chunk := invPermuted[i : i+4]
	// 	num, _ := strconv.ParseUint(chunk, 2, 4)
	// 	cryptotext.WriteString(fmt.Sprintf("%X", num))
	// }

	number, _ := strconv.ParseUint(invPermuted, 2, 64)
	cryptotext := fmt.Sprintf("%X", number)
	return cryptotext

}
func main() {
	key := "133457799BBCDFF1"
	// 00010010011010010101101111001001101101111011011111111000
	// 00010010011010010101101111001001101101111011011111111000
	keyArrayOfBytes := toBytes(key)
	keyStringOfBits := toBits(keyArrayOfBytes)
	// keyStringOfBits = deleteBits(keyStringOfBits)
	// fmt.Println("key bits: ", keyStringOfBits)
	// fmt.Println("-----------------------------------------------")
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
	// fmt.Println(arrayOfBytes)

	stringOfBits := toBits(arrayOfBytes)
	// fmt.Println(stringOfBits)

	bitsAfterP := initialPermutation(stringOfBits, initialP)
	// fmt.Println(bitsAfterP)

	L0 := bitsAfterP[0:32]
	R0 := bitsAfterP[32:64]

	// fmt.Println("L0 = ", L0)
	// fmt.Println("R0 = ", R0)

	cryptotext := encrypt(L0, R0, keyStringOfBits)
	fmt.Println(cryptotext)

	// 000110110000001011101111111111000111000001110010
	// 00011011OOOO0O1011101111111111000111OOOOO1110010

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
