package main

import (
	"fmt"
	"math/rand"
	"time"
)

func PRGA(s []byte, k []byte, nr int) []byte {
	keystream := make([]byte, nr)
	j := 0
	for i := 0; i < nr; i++ {

		j = (j + int(s[i+1])) % 256

		temp := s[i+1]
		s[i+1] = s[j]
		s[j] = temp

		t := int(s[i+1]+s[j]) % 256
		keystream[i] = s[t]
	}
	return keystream
}

func initialization(key []byte) ([]byte, []byte) {

	s := make([]byte, 256)
	k := make([]byte, 256)
	N := len(key)
	for i := 0; i <= 255; i++ {
		s[i] = byte(i)
		k[i] = key[i%N]

	}

	j := 0
	for i := 0; i <= 255; i++ {
		j = (j + int(s[i]) + int(k[i])) % 256
		temp := s[i]
		s[i] = s[j]
		s[j] = temp
	}

	return s, k

}

func generateIVS() [][3]byte {

	ivs := make([][3]byte, 200)
	for i := 0; i < 200; i++ {
		ivs[i] = [3]byte{3, 255, byte(rand.Intn(256))}
	}
	return ivs
}

func pairs(key []byte) []map[byte][]byte {

	ivs := generateIVS()

	list := make([]map[byte][]byte, len(ivs))

	for i := 0; i < len(ivs); i++ {

		pairs := make(map[byte][]byte)
		newKey := append(ivs[i][:], key...)

		s, k := initialization(newKey)
		keystreamFirstByte := PRGA(s, k, 1)[0]

		pairs[keystreamFirstByte] = ivs[i][:]
		list[i] = pairs
	}

	return list
}

func max(list map[byte]int) byte {

	max := 0
	var keyByte byte
	for key, value := range list {

		if value >= max {
			max = value
			keyByte = key
		}
	}
	return keyByte

}

func fms(pairs []map[byte][]byte) map[byte]int {

	freqK3 := make(map[byte]int)
	for i := 0; i < len(pairs); i++ {

		for key, value := range pairs[i] {

			K3 := byte((int(key) - 6 - int(value[2])) % 256)
			if freqK3[K3] == 0 {
				freqK3[K3] = 1
			} else {
				freqK3[K3] += 1
			}
		}
	}

	return freqK3
}
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func main() {

	key := []byte("aykey")
	// iv := [3]byte{3, 255, byte(rand.Intn(256))}
	// newKey := append(iv[:], key...)
	// s, k := initialization(newKey)

	// keystream := PRGA(s, k, 64)
	// fmt.Println("KeyStream : ", keystream[1])

	pairs := pairs(key)

	freq := fms(pairs)

	fmt.Println("K3: ", max(freq))
	fmt.Println(freq)

	var count int
	count = 0
	for i := 0; i < 10001; i++ {

		key = []byte(randomString(10))
		s, k := initialization(key)
		keystream := PRGA(s, k, 64)
		secondByte := keystream[1]
		if secondByte == 0 {
			count = count + 1
		}

	}

	fmt.Println("second byte = 0 : ", float64(count)/float64(1000))

}
