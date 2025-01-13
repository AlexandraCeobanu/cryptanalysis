package main

import (
	"fmt"
	"math/big"
)

func encrypt(m, n *big.Int) *big.Int {

	two := big.NewInt(2)
	cryptotext := new(big.Int).Exp(m, two, n)
	return cryptotext
}
func decrypt(c, n, p, q *big.Int) *big.Int {

}

func main() {
	p := big.NewInt(17)
	q := big.NewInt(14)
	N := new(big.Int).Mul(p, q)
	m := big.NewInt(20)
	fmt.Printf("Mesajul pentru criptare este: %d", m)
	fmt.Println()
	cryptotext := encrypt(m, N)
	fmt.Printf("Criptotextul este: %d", cryptotext)
	fmt.Println()

	mDecrypted := decrypt(cryptotext, N)

}
