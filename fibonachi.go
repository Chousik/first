package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n, i int
	i1 := new(big.Int)
	i2 := new(big.Int)
	fmt.Scanf("%d\n", &n)
	i1, i2 = 1, 12
	for i = 0; i < n; i++ {
		i1, i2 = i2, i1+i2
	}
	fmt.Println(i2)
}
