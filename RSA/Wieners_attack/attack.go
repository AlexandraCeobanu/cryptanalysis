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

func check(e, l, d *big.Int) bool {
	two := big.NewInt(2)
	zero := big.NewInt(0)
	one := big.NewInt(1)
	if new(big.Int).Mod(d, two).Cmp(zero) == 0 {
		return false
	}

	ed := new(big.Int).Mul(e, d)
	ed_1 := new(big.Int).Sub(ed, one)

	if l.Cmp(zero) != 0 && new(big.Int).Mod(ed_1, l).Cmp(zero) != 0 {
		return false
	}
	return true
}
func resolveSystem(e, n, l, d *big.Int) (*big.Int, *big.Int) {

	one := big.NewInt(1)
	two := big.NewInt(2)
	four := big.NewInt(4)
	zero := big.NewInt(0)
	ed := new(big.Int).Mul(e, d)
	ed_1 := new(big.Int).Sub(ed, one)
	ed_1_l := new(big.Int).Div(ed_1, l)

	sum := new(big.Int).Sub(n, ed_1_l)
	sum = sum.Add(sum, one)

	delta := new(big.Int).Mul(sum, sum)
	prod4 := new(big.Int).Mul(n, four)
	delta = delta.Sub(delta, prod4)

	if delta.Cmp(zero) == -1 {
		fmt.Println("Nu exista solutii pentru sistem")
		return nil, nil
	}
	rad_delta := new(big.Int).Sqrt(delta)
	p := new(big.Int).Add(rad_delta, sum)
	p = p.Div(p, two)
	q := new(big.Int).Sub(sum, rad_delta)
	q = q.Div(q, two)

	return p, q
}
func findDPQ(e, n *big.Int) (*big.Int, *big.Int, *big.Int) {
	zero := big.NewInt(0)
	one := big.NewInt(1)
	q_curent := new(big.Int).Div(e, n)
	r_curent := new(big.Int).Mod(e, n)
	n_curent := big.NewInt(e.Int64())
	x := big.NewInt(n.Int64())
	it := 1
	alpha_i := new(big.Int)
	betha_i := new(big.Int)
	alpha_i_2 := new(big.Int)
	betha_i_2 := new(big.Int)
	q_prec := new(big.Int)
	for {
		if r_curent.Cmp(zero) == 0 {
			break
		}

		if it == 1 {

			alpha_i = big.NewInt(int64(q_curent.Int64()))
			betha_i = big.NewInt(1)
			alpha_i_2 = big.NewInt(int64(alpha_i.Int64()))
			betha_i_2 = big.NewInt(int64(betha_i.Int64()))

		} else if it == 2 {
			alpha_i = new(big.Int).Mul(q_prec, q_curent)
			alpha_i = alpha_i.Add(alpha_i, one)
			betha_i = big.NewInt(int64(q_curent.Int64()))
		} else {
			alpha_i_2 = big.NewInt(int64(alpha_i.Int64()))
			betha_i_2 = big.NewInt(int64(betha_i.Int64()))
			alpha_i = new(big.Int).Mul(q_curent, alpha_i)
			alpha_i = new(big.Int).Add(alpha_i, alpha_i_2)
			betha_i = new(big.Int).Mul(q_curent, betha_i)
			betha_i = new(big.Int).Add(betha_i, betha_i_2)
		}

		l := big.NewInt(alpha_i.Int64())
		d := big.NewInt(betha_i.Int64())
		result := check(e, l, d)

		if result && l.Cmp(zero) != 0 {
			p, q := resolveSystem(e, n_curent, l, d)
			if p != nil && q != nil {
				return d, p, q
			}
		}

		n_curent = big.NewInt(x.Int64())
		q_prec = big.NewInt(int64(q_curent.Int64()))
		q_curent = new(big.Int).Div(x, r_curent)
		r_curent = new(big.Int).Mod(x, r_curent)
		x = big.NewInt(r_curent.Int64())
		it = it + 1
	}
	return nil, nil, nil
}
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
func main() {

	e := big.NewInt(3467)
	n := big.NewInt(10605)
	d, p, q := findDPQ(e, n)
	if d == nil || p == nil || q == nil {
		fmt.Println("Nu s-au gasit d,p,q")
		fmt.Printf("d = %d , p = %d, q = %d ", d, p, q)
	} else {
		fmt.Printf("d = %d , p = %d, q = %d ", d, p, q)
	}

	// p := generateP()
	// q := generateP()
	// n := big.NewInt(new(big.Int).Mul(p, q).Int64())
	// p_1 := new(big.Int).Sub(p, big.NewInt(1))
	// q_1 := new(big.Int).Sub(q, big.NewInt(1))
	// euler := big.NewInt(new(big.Int).Mul(p_1, q_1).Int64())
	// e := selectE(euler)

	// d, p, q := findDPQ(e, n)
	// if d == nil || p == nil || q == nil {
	// 	fmt.Println("Nu s-au gasit d,p,q")
	// 	fmt.Printf("d = %d , p = %d, q = %d ", d, p, q)
	// } else {
	// 	fmt.Printf("d = %d , p = %d, q = %d ", d, p, q)
	// }
}
