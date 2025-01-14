package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func gcd(a, b *big.Int) *big.Int {
	zero := big.NewInt(0)
	for b.Cmp(zero) != 0 {
		a, b = b, new(big.Int).Mod(a, b)
	}
	return a
}
func selectE(euler *big.Int) *big.Int {
	one := big.NewInt(1)
	for {
		e, err := rand.Int(rand.Reader, euler)
		if err != nil {
			fmt.Println("Eroare:", err)
			return nil
		}

		if e.Cmp(one) > 0 && gcd(euler, e).Cmp(one) == 0 {
			return e
		}
	}
}
func chooseM(n *big.Int) *big.Int {

	one := big.NewInt(1)
	for {
		m, err := rand.Int(rand.Reader, n)
		if err != nil {
			fmt.Println("Eroare:", err)
			return nil
		}

		if m.Cmp(one) >= 0 && gcd(m, n).Cmp(one) == 0 {
			return m
		}
	}

}

func encrypt(m, n, e *big.Int) *big.Int {

	cryptotext := new(big.Int).Exp(m, e, n)
	return cryptotext

}
func decrypt(cryptotext, n, d *big.Int) *big.Int {

	messageDecrypted := new(big.Int).Exp(cryptotext, d, n)
	return messageDecrypted

}
func main() {

	p := 61
	q := 53
	n := big.NewInt(int64(p) * int64(q))
	phi := big.NewInt(int64(p-1) * int64(q-1))
	e := selectE(phi)

	d := new(big.Int).ModInverse(e, phi)
	if d == nil {
		fmt.Printf("nu exista invers multiplicativ pentru %d modulo %d", e, phi)
	}

	// m := big.NewInt(65)
	m := chooseM(n)
	fmt.Printf("Criptam mesajul %d ", m)
	cryptotext := encrypt(m, n, e)
	fmt.Println()
	fmt.Printf("Criptarea este: %d", cryptotext)

	messageDecrypted := decrypt(cryptotext, n, d)
	fmt.Println()
	fmt.Printf("Mesajul decriptat este: %d", messageDecrypted)
}
