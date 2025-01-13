package main

import (
	"fmt"
	"math/big"
)

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

	fmt.Println("prod ", n)
	fmt.Println("ed-1/l ", ed_1_l)
	sum := new(big.Int).Sub(n, ed_1_l)
	sum = sum.Add(sum, one)
	fmt.Println("suma ", sum)
	delta := new(big.Int).Mul(sum, sum)
	prod4 := new(big.Int).Mul(n, four)
	delta = delta.Sub(delta, prod4)
	fmt.Println("delta ", delta)
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

		fmt.Println("q ", q_curent)
		fmt.Println("r ", r_curent)
		fmt.Println("alpha ", alpha_i)
		fmt.Println("betha ", betha_i)

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
}
