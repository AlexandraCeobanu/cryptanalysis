package main

import (
	"fmt"
	"math/big"
)

func cmmdc(a, b *big.Int) *big.Int {
	a_cpy := big.NewInt(a.Int64())
	b_cpy := big.NewInt(b.Int64())
	zero := big.NewInt(0)
	for b_cpy.Cmp(zero) != 0 {
		a_cpy, b_cpy = b_cpy, new(big.Int).Mod(a_cpy, b_cpy)
	}
	return a_cpy
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

func sign(p, q, n, m, d, e *big.Int) *big.Int {

	m_mod_p := new(big.Int).Mod(m, p)
	p_1 := new(big.Int).Sub(p, big.NewInt(1))
	d_mod_p_1 := new(big.Int).Mod(d, p_1)
	sp := new(big.Int).Exp(m_mod_p, d_mod_p_1, p)

	m_mod_q := new(big.Int).Mod(m, q)
	q_1 := new(big.Int).Sub(q, big.NewInt(1))
	d_mod_q_1 := new(big.Int).Mod(d, q_1)
	sq := new(big.Int).Exp(m_mod_q, d_mod_q_1, q)

	sp = big.NewInt(5)

	s := CRT(sp, p, sq, q)

	s_e := new(big.Int).Exp(s, e, n)
	s_e_m := new(big.Int).Sub(s_e, m)
	qF := cmmdc(s_e_m, n)
	fmt.Println("Am aflat q: ", qF)

	return s

}

func main() {
	p := big.NewInt(7)
	q := big.NewInt(11)
	n := new(big.Int).Mul(p, q)
	m := big.NewInt(62)
	e := big.NewInt(7)
	phi := new(big.Int).Mul(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))
	d := new(big.Int).ModInverse(e, phi)
	if d == nil {
		fmt.Printf("nu exista invers multiplicativ pentru %d modulo %d", e, phi)
	}
	s := sign(p, q, n, m, d, e)

	fmt.Println("Semnatura ", s)
}
