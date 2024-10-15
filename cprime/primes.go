package cprime

import (
	"math"
	"runtime"
)

var numCores int

func init() {
	numCores = runtime.NumCPU()
}

type faixa struct {
	start uint64
	lim   uint64
}

// trivialNonPrime verifica se natural é um não-primo trivial
func trivialNonPrime(natural uint64) bool {
	if natural == 1 {
		return true
	}
	if natural%2 == 0 && natural != 2 {
		return true
	}
	return false
}

// IsPrime versão não concorrente objetivada para comparação de performance
func IsPrime(natural uint64) bool {
	if trivialNonPrime(natural) {
		return false
	}

	squareRoot := uint64(math.Sqrt(float64(natural)))
	i := uint64(3)
	for ; i <= squareRoot; i += 2 {
		if natural%i == 0 {
			break
		}
	}

	return i > squareRoot
}

// IsPrimeConcurrent dispara goroutines que testaram os divisores da sua parte do intervalo de divisores
func IsPrimeConcurrent(natural uint64) bool {
	if trivialNonPrime(natural) {
		return false
	}
	squareRoot := uint64(math.Sqrt(float64(natural)))
	intervalos := generateIntervals(faixa{3, squareRoot}, uint64(numCores))
	canal := make(chan bool)
	for _, inter := range intervalos {
		go isPrimeInInterval(natural, inter.start, inter.lim, canal)
	}

	for i := 0; i < len(intervalos); i++ {
		if !<-canal {
			return false
		}
	}

	return true
}

// generateIntervals divide o intervalo dos divisores em n partes
func generateIntervals(interval faixa, times uint64) []faixa {

	if interval.lim < interval.start {
		return []faixa{interval}
	}
	sizeInterval := (interval.lim - interval.start) / times
	if sizeInterval == 0 {
		return []faixa{interval}
	}
	remainder := (interval.lim - interval.start) % times

	intervals := make([]faixa, times)
	intervals[0].start = interval.start
	intervals[0].lim = interval.start + sizeInterval + remainder
	for i := uint64(1); i < times; i++ {
		intervals[i].start = intervals[i-1].lim
		intervals[i].lim = intervals[i].start + sizeInterval
	}

	return intervals
}

// isPrimeInInterval testa se o número não tem divisor dentro de um determinado intervalo
func isPrimeInInterval(natural, start, lim uint64, ret chan bool) {
	if start%2 == 0 {
		start--
	}
	for i := start; i < lim; i += 2 {
		if natural%i == 0 {
			ret <- false
			return
		}
	}
	ret <- true
}
