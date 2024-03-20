// Разработать программу, которая
// перемножает, делит, складывает, вычитает две числовых переменных a,b,
// значение которых > 2^20
package main

import (
	"fmt"
	"math/big"
)

func Mul(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func Div(a, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b)
}

func Add(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func Sub(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func main() {
	var a, b big.Int
	a.Lsh(big.NewInt(2), 21)
	b.Lsh(big.NewInt(2), 22)
	fmt.Println("Mul", a, b, Mul(&a, &b))
	fmt.Println("Div", a, b, Div(&a, &b))
	fmt.Println("Add", a, b, Add(&a, &b))
	fmt.Println("Sub", a, b, Sub(&a, &b))
}
