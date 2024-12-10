package main

import (
	"fmt"
	"math/big"
)

func modularInverse(x, p *big.Int) *big.Int {
	inv := new(big.Int).ModInverse(x, p)
	if inv == nil {
		panic("Inversul modular nu exista")
	}
	return inv
}
func shank(g, p, x *big.Int) int {

	one := big.NewInt(1)
	m := new(big.Int).Sqrt(new(big.Int).Sub(p, one))
	table := make([]*big.Int, m.Int64())

	g_inv := modularInverse(g, p)
	for i := int64(0); i < m.Int64(); i++ {

		im := new(big.Int).Mul(big.NewInt(i), m)

		g_la_im := new(big.Int).Exp(g_inv, im, p)

		x_g_la_im := new(big.Int).Mul(x, g_la_im)
		x_g_la_im.Mod(x_g_la_im, p)

		fmt.Println(x_g_la_im)
		table[int(i)] = x_g_la_im
	}

	for j := int64(0); j < m.Int64(); j++ {

		g_la_j := new(big.Int).Exp(g, big.NewInt(j), p)

		for i := int64(0); i < m.Int64(); i++ {
			if (g_la_j.Cmp(table[i])) == 0 {

				fmt.Println("i= ", i)
				fmt.Println("j = ", j)
				a := new(big.Int).Add(new(big.Int).Mul(big.NewInt(int64(i)), m), big.NewInt(j))
				return int(a.Int64())
			}

		}

	}

	return -1
}

func main() {

	g := big.NewInt(3)
	p := big.NewInt(101)
	x := big.NewInt(37)

	a := shank(g, p, x)
	fmt.Println("a = ", a)

}
