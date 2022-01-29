package main

import (
	"fmt"
	"math"
)

func main() {
	var p, q int = 37, 53
	x, r := Alice(p, q)
	y, s := Bob(p, q)
	k1 := int(math.Pow(float64(x), float64(s))) % q
	k2 := int(math.Pow(float64(y), float64(r))) % q
	fmt.Println("Key1:", k1)
	fmt.Println("Key2:", k2)
}

func Alice(p, q int) (int, int) {
	r := 5
	x := int(math.Pow(float64(p), float64(r))) % q
	return x, r
}

func Bob(p, q int) (int, int) {
	s := 3
	y := int(math.Pow(float64(p), float64(s))) % q
	return y, s
}
