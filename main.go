package main

import (
	"concurrentPrimes/cprime"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	decimals := convert(os.Args[1:])
	fmt.Println("\nNon concurrent")
	iterator(cprime.IsPrime, decimals)
	fmt.Println("\nConcurrent version ")
	iterator(cprime.IsPrimeConcurrent, decimals)
}

func convert(strInts []string) []uint64 {

	out := make([]uint64, 0, len(strInts))

	for _, v := range strInts {
		decimal, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			fmt.Printf("erro ao converter %v, %v\n", v, err)
		} else {
			out = append(out, decimal)
		}

	}
	return out
}

func iterator(f func(uint64) bool, naturals []uint64) {
	start := time.Now()
	var b bool
	for _, d := range naturals {
		b = f(d)
		if b {
			fmt.Printf("\n%v é primo", d)
		} else {
			fmt.Printf("\n%v não é primo", d)
		}
	}

	spendTime := time.Since(start).Microseconds()
	fmt.Printf("\nLevou %v microsegundos\n", spendTime)
}
