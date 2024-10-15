package cprime

import (
	"slices"
	"testing"
)

type scenario struct {
	param   uint64
	isPrime bool
}

var scenarios []scenario = []scenario{
	{1, false},
	{2, true},
	{3, true},
	{5, true},
	{8, false},
	{7, true},
	{11, true},
	{37027414657, true},
	{46810093859, true},
	{(1 << 61) - 1, true},
	{18446744073709551557, true},
	{18446744073709551559, false},
}

func printMessage(t *testing.T, scene scenario) {
	if scene.isPrime {
		t.Errorf("é falso que o valor %v é primo", scene.param)
	} else {
		t.Errorf("é falso que o valor %v é composto", scene.param)
	}
}

func TestIsPrime(t *testing.T) {
	for _, value := range scenarios {
		if IsPrime(value.param) != value.isPrime {
			printMessage(t, value)
		}
	}
}

func TestIsPrimeConcurrent(t *testing.T) {
	for _, value := range scenarios {
		if IsPrimeConcurrent(value.param) != value.isPrime {
			printMessage(t, value)
		}
	}
}

func TestGenerateInterval(t *testing.T) {
	param1 := faixa{3, 9}
	param2 := uint64(6)
	esperado := []faixa{{3, 4}, {4, 5}, {5, 6}, {6, 7}, {7, 8}, {8, 9}}
	retorno := generateIntervals(param1, param2)

	if !slices.Equal(retorno, esperado) {
		t.Errorf("precisava de %v, retornou: %v\n", esperado, retorno)
	}

	param1 = faixa{3, 112}
	param2 = uint64(10)
	esperado = []faixa{{3, 22}, {22, 32}, {32, 42}, {42, 52}, {52, 62}, {62, 72}, {72, 82}, {82, 92}, {92, 102}, {102, 112}}
	retorno = generateIntervals(param1, param2)

	if !slices.Equal(retorno, esperado) {
		t.Errorf("precisava de %v, retornou: %v\n", esperado, retorno)
	}
}
