package main

import "fmt"

func main() {
	fmt.Println(ordenacaoPorSelecao([]int{5, 3, 6, 2, 10}))
}

func buscaMenor(array []int) int {
	menor := array[0]
	menorIndice := 0

	for i := 0; i < len(array); i++ {
		if array[i] < menor {
			menor = array[i]
			menorIndice = i
		}
	}

	return menorIndice
}

func ordenacaoPorSelecao(array []int) []int {
	var novoArray []int

	for len(array) > 0 {
		menor := buscaMenor(array)
		novoArray = append(novoArray, array[menor])
		array = append(array[:menor], array[menor+1:]...)
	}

	return novoArray
}
