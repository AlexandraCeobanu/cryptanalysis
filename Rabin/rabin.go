package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// var prime_factors = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409, 419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523, 541}
// var prime_factors = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

var prime_factors = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}

// var prime_factors = []int{2, 3, 5, 7}

func isPrime(p *big.Int) bool {
	return p.ProbablyPrime(20)
}
func generateP() *big.Int {

	for {

		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(prime_factors))))
		initialP := big.NewInt(int64(prime_factors[index.Int64()]))
		for {

			if initialP.BitLen() >= 10 {
				break
			}
			index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(prime_factors))))
			factor := big.NewInt(int64(prime_factors[index.Int64()]))
			if new(big.Int).Mul(initialP, factor).BitLen() <= 512 {

				initialP.Mul(initialP, factor)
			}

		}
		p := new(big.Int).Add(initialP, big.NewInt(1))

		if isPrime(p) {

			return p
		}
	}
}

func cmmdc(a, b *big.Int) *big.Int {
	a_cpy := big.NewInt(a.Int64())
	b_cpy := big.NewInt(b.Int64())
	zero := big.NewInt(0)
	for b_cpy.Cmp(zero) != 0 {
		a_cpy, b_cpy = b_cpy, new(big.Int).Mod(a_cpy, b_cpy)
	}
	return a_cpy
}
func findX(x *big.Int, array []*big.Int) *big.Int {
	for i := 0; i < len(array); i++ {
		if array[i].Cmp(x) == 0 {
			if i != 0 {
				return array[i-1]
			} else {
				return big.NewInt(1)
			}
		}
	}
	return big.NewInt(-1)
}
func Pollard(n *big.Int) (*big.Int, *big.Int) {
	a_curent := big.NewInt(2)
	rezultateF := []*big.Int{}
	one := big.NewInt(1)
	minusOne := big.NewInt(-1)
	rezultateF = append(rezultateF, a_curent)

	for {

		a_curent = new(big.Int).Mul(a_curent, a_curent)
		a_curent = a_curent.Add(a_curent, one)
		a_curent.Mod(a_curent, n)

		col := findX(a_curent, rezultateF)
		if col.Cmp(minusOne) == 0 {
			rezultateF = append(rezultateF, a_curent)
		} else {
			precendent := rezultateF[len(rezultateF)-1]
			max := new(big.Int)
			min := new(big.Int)
			if precendent.Cmp(col) == -1 {
				max = col
				min = precendent
			} else {
				max = precendent
				min = col
			}

			p := cmmdc(new(big.Int).Sub(max, min), n)
			q := cmmdc(new(big.Int).Add(max, min), n)
			return p, q
		}
	}

}

func jacobi(x, n *big.Int) *big.Int {
	one := big.NewInt(1)
	minusone := big.NewInt(-1)
	two := big.NewInt(2)
	eight := big.NewInt(8)
	seven := big.NewInt(7)
	three := big.NewInt(3)
	five := big.NewInt(5)
	four := big.NewInt(4)
	minus := false
	for {
		if x.Cmp(one) == 0 {
			if minus == false {
				return one
			}
			return minusone
		} else if x.Cmp(two) == 0 {
			if new(big.Int).Mod(n, eight).Cmp(one) == 0 || new(big.Int).Mod(n, eight).Cmp(seven) == 0 {
				if minus == false {
					return one
				}
				return minusone

			}
			if new(big.Int).Mod(n, eight).Cmp(three) == 0 || new(big.Int).Mod(n, eight).Cmp(five) == 0 {
				if minus == false {
					return minusone
				}

				return one
			}
		} else if x.Cmp(n) == 1 {
			x = x.Mod(x, n)
		} else if new(big.Int).Mod(n, four).Cmp(one) == 0 || new(big.Int).Mod(x, four).Cmp(one) == 0 {
			temp := big.NewInt(x.Int64())
			x = big.NewInt(n.Int64())
			n = big.NewInt(temp.Int64())
		} else if new(big.Int).Mod(n, four).Cmp(three) == 0 && new(big.Int).Mod(x, four).Cmp(three) == 0 {
			temp := big.NewInt(x.Int64())
			x = big.NewInt(n.Int64())
			n = big.NewInt(temp.Int64())
			if minus == false {
				minus = true
			} else {
				minus = false
			}

		}

	}
}
func encrypt(m, n *big.Int) []*big.Int {

	cryptotext := []*big.Int{}
	two := big.NewInt(2)
	cryptotext1 := new(big.Int).Exp(m, two, n)
	cryptotext = append(cryptotext, cryptotext1)

	x_2 := new(big.Int).Mod(m, two)

	cryptotext = append(cryptotext, x_2)

	jacobi_x_n := jacobi(m, n)

	cryptotext = append(cryptotext, jacobi_x_n)
	return cryptotext
}
func reziduu(x *big.Int, n *big.Int) (*big.Int, *big.Int) {
	sol := new(big.Int).ModSqrt(x, n)
	sol2 := new(big.Int)
	if sol == nil {
		fmt.Printf("Nu există soluiii")
		return nil, nil
	} else {

		sol2 = new(big.Int).Sub(n, sol)
	}
	return sol, sol2
}
func CRT(c1, p1, c2, p2 *big.Int) *big.Int {
	M := new(big.Int).Mul(p1, p2)
	M1 := new(big.Int).Div(M, p1)
	M2 := new(big.Int).Div(M, p2)

	y1 := new(big.Int).ModInverse(M1, p1)
	y2 := new(big.Int).ModInverse(M2, p2)

	minusone := big.NewInt(-1)
	zero := big.NewInt(0)
	if y1.Cmp(minusone) == 0 || y2.Cmp(minusone) == 0 {
		fmt.Println("Inversul modular nu exista")
		return big.NewInt(-1)
	}

	x := new(big.Int).Mul(c1, M1)
	x = x.Mul(x, y1)
	c2_M2_y2 := new(big.Int).Mul(c2, M2)
	c2_M2_y2 = c2_M2_y2.Mul(c2_M2_y2, y2)
	x = x.Add(x, c2_M2_y2)
	x = x.Mod(x, M)

	if x.Cmp(zero) == -1 {
		x.Add(x, M)
	}

	return x
}
func decrypt(cryptotext []*big.Int, n *big.Int) *big.Int {
	two := big.NewInt(2)
	y := cryptotext[0]
	mod := cryptotext[1]
	jacobis := cryptotext[2]
	p, q := Pollard(n)
	x1, x2 := reziduu(y, p)
	x3, x4 := reziduu(y, q)

	// fmt.Println("x1 = ", x1)
	// fmt.Println("x2= ", x2)
	// fmt.Println("x3= ", x3)
	// fmt.Println(" x4= ", x4)

	x1x3 := CRT(x1, p, x3, q)
	x1x4 := CRT(x1, p, x4, q)

	x2x3 := CRT(x2, p, x3, q)
	x2x4 := CRT(x2, p, x4, q)

	fmt.Println("x1 = ", x1x3)
	fmt.Println("x2= ", x1x4)
	fmt.Println("x3= ", x2x3)
	fmt.Println(" x4= ", x2x4)

	if new(big.Int).Mod(x1x3, two).Cmp(mod) == 0 && jacobi(x1x3, n).Cmp(jacobis) == 0 {
		return x1x3
	}
	if new(big.Int).Mod(x1x4, two).Cmp(mod) == 0 && jacobi(x1x4, n).Cmp(jacobis) == 0 {
		return x1x4
	}
	if new(big.Int).Mod(x2x3, two).Cmp(mod) == 0 && jacobi(x2x3, n).Cmp(jacobis) == 0 {
		return x2x3
	}
	if new(big.Int).Mod(x2x4, two).Cmp(mod) == 0 && jacobi(x2x4, n).Cmp(jacobis) == 0 {
		return x2x4
	}
	return nil
}

func main() {

	// p := big.NewInt(7)
	// q := big.NewInt(3)
	// N := new(big.Int).Mul(p, q)
	// m := big.NewInt(5)
	// fmt.Printf("Mesajul pentru criptare este: %d", m)
	// fmt.Println()
	// cryptotext := encrypt(m, N)
	// fmt.Printf("Criptotextul este: %d", cryptotext)
	// fmt.Println()

	// mDecrypted := decrypt(cryptotext, N)
	// fmt.Println("Mesajul decriptat este: ", mDecrypted)

	p := generateP()
	q := generateP()
	N := new(big.Int).Mul(p, q)
	m := big.NewInt(5)
	fmt.Printf("Mesajul pentru criptare este: %d", m)
	fmt.Println()
	cryptotext := encrypt(m, N)
	fmt.Printf("Criptotextul este: %d", cryptotext)
	fmt.Println()

	mDecrypted := decrypt(cryptotext, N)
	fmt.Println("Mesajul decriptat este: ", mDecrypted)
}
