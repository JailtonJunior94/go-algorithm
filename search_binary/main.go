package main

import "fmt"

func main() {
	searchItem := 4
	listItems := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	found, count := searchBinary(listItems, searchItem)
	if found {
		fmt.Printf("O item %d foi encontrado e a quantidade de tentativas foi %d\n", searchItem, count)
		return
	}
	fmt.Printf("O item NÃO foi encontrado e a quantidade de tentativas foi %d\n", count)
}

func searchBinary(list []int, searchItem int) (bool, int) {
	begin := 0
	end := len(list) - 1

	count := 0
	current := 0

	for begin <= end {
		var middle int = (begin + end) / 2
		current := list[middle]

		if current == searchItem {
			return true, count
		}

		if current > searchItem {
			end = middle - 1
		}

		if current < searchItem {
			begin = middle + 1
		}
		count++
	}
	return false, current
}
