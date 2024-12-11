package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var prime_factors = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409, 419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523, 541}

func isPrime(p *big.Int) bool {
	return p.ProbablyPrime(20)
}
func generateP() *big.Int {

	for {

		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(prime_factors))))
		initialP := big.NewInt(int64(prime_factors[index.Int64()]))
		for {

			if initialP.BitLen() >= 1024 {
				break
			}
			index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(prime_factors))))
			factor := big.NewInt(int64(prime_factors[index.Int64()]))
			initialP.Mul(initialP, factor)
		}
		p := new(big.Int).Add(initialP, big.NewInt(1))
		if isPrime(p) {
			return p
		}
	}

}
func modularInverse(x, p *big.Int) *big.Int {
	inv := new(big.Int).ModInverse(x, p)
	if inv == nil {
		panic("Inversul modular nu exista")
	}
	return inv
}
func shank(alpha, p, beta *big.Int) int {

	one := big.NewInt(1)
	n := new(big.Int).Sub(p, one)
	m := new(big.Int).Sqrt(n)
	table := make(map[int]*big.Int)

	for j := 0; j < int(m.Int64()); j++ {

		alpha_la_j := new(big.Int).Exp(alpha, big.NewInt(int64(j)), p)
		table[j] = alpha_la_j
	}

	alpha_inv := modularInverse(alpha, p)
	alpha_la_m := new(big.Int).Exp(alpha_inv, m, p)
	gamma := beta

	for i := 0; i < int(m.Int64()); i++ {

		for j := 0; j < int(m.Int64()); j++ {

			if (gamma.Cmp(table[j])) == 0 {

				x := new(big.Int).Add(new(big.Int).Mul(big.NewInt(int64(i)), m), big.NewInt(int64(j)))
				return int(x.Int64())
			}
		}

		gamma.Mul(gamma, alpha_la_m)

	}
	return -1
}

func primeFactorization(n *big.Int) []map[int]int {

	factorization := make(map[int]int)
	index_factor := 0
	factor := big.NewInt(int64(prime_factors[index_factor]))
	for {
		if n.Cmp(big.NewInt(1)) == 0 || index_factor >= len(prime_factors) {
			break
		}
		power := 0
		for {

			if new(big.Int).Mod(n, factor).Cmp(big.NewInt(0)) != 0 {
				break
			}

			power = power + 1
			n.Div(n, factor)
		}
		if power != 0 {

			factorization[int(factor.Int64())] = power
		}

		index_factor = index_factor + 1
		if index_factor < len(prime_factors) {
			factor = big.NewInt(int64(prime_factors[index_factor]))
		}

	}

	list_factorization := make([]map[int]int, len(factorization))
	for i := range list_factorization {
		list_factorization[i] = make(map[int]int)
	}

	i := 0
	for key, value := range factorization {
		list_factorization[i][key] = value
		i = i + 1
	}

	return list_factorization
}

func returnKeyValue(factorization []map[int]int, i int) (int, int) {

	for key, value := range factorization[i] {
		return key, value
	}
	return -1, -1
}
func gauss(xi []*big.Int, primeFactors []map[int]int) *big.Int {

	r := len(primeFactors)
	m := big.NewInt(1)
	for i := 0; i < r; i++ {

		pi, ei := returnKeyValue(primeFactors, i)
		pi_ei := new(big.Int).Exp(big.NewInt(int64(pi)), big.NewInt(int64(ei)), nil)

		m.Mul(m, pi_ei)
	}

	fmt.Println("m = ", m)

	x := big.NewInt(0)
	for i := 0; i < r; i++ {

		p_i, e_i := returnKeyValue(primeFactors, i)

		m_i := new(big.Int).Exp(big.NewInt(int64(p_i)), big.NewInt(int64(e_i)), nil)
		c_i := new(big.Int).Div(m, m_i)

		c_i_inv := modularInverse(c_i, m_i)

		x_i := new(big.Int).Mul(c_i_inv, xi[i])
		m_x_i := new(big.Int).Mul(x_i, c_i)

		x.Add(x, m_x_i)

	}
	return x.Mod(x, m)
}
func silverPohligHellman(alpha *big.Int, n *big.Int, beta *big.Int, p *big.Int) *big.Int {

	// primeFactorization(10)
	n_cpy := big.NewInt(n.Int64())
	factorization := primeFactorization(n_cpy)

	r := len(factorization)
	var xs []*big.Int
	for i := 0; i < r; i++ {

		q, e := returnKeyValue(factorization, i)

		gamma := big.NewInt(1)
		ant_l := big.NewInt(0)
		n_div_q := new(big.Int)
		n_div_q.Div(n, big.NewInt(int64(q)))

		alpha_b := new(big.Int).Exp(alpha, n_div_q, p)
		var ls []*big.Int

		for j := 0; j < e; j++ {
			j_m1 := new(big.Int).Sub(big.NewInt(int64(j)), big.NewInt(1))

			q_la_j := new(big.Int)
			if j_m1.Cmp(big.NewInt(-1)) == 0 {
				q_la_j = modularInverse(big.NewInt(int64(q)), p)
			} else {
				q_la_j = new(big.Int).Exp(big.NewInt(int64(q)), j_m1, p)
			}

			l_j_q_j := new(big.Int)
			l_j_q_j.Mul(ant_l, q_la_j)

			alpha2 := new(big.Int).Exp(alpha, l_j_q_j, p)

			gamma = new(big.Int).Mul(gamma, alpha2)
			gamma.Mod(gamma, p)

			j_p1 := new(big.Int).Add(big.NewInt(int64(j)), big.NewInt(1))
			q_la_j_p1 := new(big.Int).Exp(big.NewInt(int64(q)), j_p1, p)
			n_q_la_j_p1 := new(big.Int).Div(n, q_la_j_p1)
			gamma_inv := modularInverse(gamma, p)

			beta_gamma_inv := new(big.Int).Mul(beta, gamma_inv)
			beta_b := new(big.Int).Exp(beta_gamma_inv, n_q_la_j_p1, p)

			l_j := big.NewInt(int64(shank(alpha_b, p, beta_b)))

			ls = append(ls, l_j)

			ant_l = big.NewInt(l_j.Int64())

		}

		x_i := new(big.Int)

		for index, element := range ls {
			if index == 0 {
				x_i = element

			} else {

				q_e := new(big.Int).Exp(big.NewInt(int64(q)), big.NewInt(int64(index)), p)

				j_q := new(big.Int).Mul(element, q_e)

				x_i.Add(x_i, j_q)
			}

		}
		xs = append(xs, x_i)

	}
	fmt.Println("xs: ", xs)
	x := gauss(xs, factorization)
	return x

}

func DiffieHellman() {
	p := generateP()
	n := new(big.Int).Sub(p, big.NewInt(1))
	g := big.NewInt(2)
	a, _ := rand.Int(rand.Reader, new(big.Int).Sub(p, big.NewInt(1)))
	fmt.Println("before Pohlig Hellman a= ", a)
	g_la_a := new(big.Int).Exp(g, a, p)
	findA := silverPohligHellman(g, n, g_la_a, p)

	fmt.Println("after Pohlig Hellman a= ", findA)
}

func main() {

	// g := big.NewInt(3)
	// p := big.NewInt(101)
	// x := big.NewInt(37)

	// a := shank(g, p, x)
	// fmt.Println("a = ", a)

	// p_prime := generateP()
	// fmt.Println("Numarul prim generat p: ", p_prime)

	alpha := big.NewInt(71)
	n := big.NewInt(250)
	beta := big.NewInt(210)
	p := big.NewInt(251)

	x := silverPohligHellman(alpha, n, beta, p)

	fmt.Println("x = ", x)
	DiffieHellman()
	// xi := []*big.Int{
	// 	big.NewInt(2),
	// 	big.NewInt(3),
	// 	big.NewInt(1),
	// }
	// primeFactors := []map[int]int{
	// 	{3: 1}, // p1 = 3^1
	// 	{5: 1}, // p2 = 5^1
	// 	{7: 1}, // p3 = 7^1
	// }

	// result := gauss(xi, primeFactors)
	// fmt.Println("Rezultat:", result)
}
