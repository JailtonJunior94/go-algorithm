package main

import "fmt"

func main() {
	// searchItem := 4
	// listItems := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	// found, count := searchBinary(listItems, searchItem)
	// if found {
	// 	fmt.Printf("O item %d foi encontrado e a quantidade de tentativas foi %d\n", searchItem, count)
	// 	return
	// }
	// fmt.Printf("O item NÃO foi encontrado e a quantidade de tentativas foi %d\n", count)

	name := "Lucas"
	listNames := []string{
		"Alice", "Bob", "Charlie", "David", "Eve", "Frank", "Grace", "Hannah", "Isaac", "Julia",
		"Liam", "Olivia", "Noah", "Emma", "Jackson", "Sophia", "Aiden", "Oliver", "Lucas", "Mia",
		"Daniel", "Ella", "Matthew", "Lily", "Michael", "Elizabeth", "Henry", "Grace", "Samuel", "Chloe",
		"James", "Ava", "Benjamin", "Amelia", "William", "Harper", "John", "Evelyn", "Joseph", "Ella",
		"Andrew", "Mila", "Logan", "Avery", "David", "Scarlett", "Nicholas", "Abigail", "Nathan", "Sofia",
		"Ryan", "Emily", "Alexander", "Charlotte", "Daniel", "Madison", "Ethan", "Harper", "Matthew", "Emma",
		"Christopher", "Luna", "Mason", "Lily", "Anthony", "Aria", "Jacob", "Eleanor", "Nicholas", "Hazel",
		"Jonathan", "Victoria", "Thomas", "Penelope", "Samuel", "Madelyn", "Nathan", "Aubrey", "Luke", "Addison",
		"Benjamin", "Stella", "Jackson", "Natalie", "William", "Zoe", "Liam", "Brooklyn", "Oliver", "Leah",
		"Lucas", "Audrey", "Carter", "Savannah", "Henry", "Skylar", "James", "Madeline", "Owen", "Ruby",
		"Sebastian", "Alice", "Joseph", "Zoey", "Matthew", "Isabella", "Wyatt", "Alyssa", "Aiden", "Alexa",
		"David", "Samantha", "Leo", "Gabriella", "Daniel", "Aria", "Gabriel", "Lucy", "Anthony", "Nora",
		"Jonathan", "Riley", "Nicholas", "Eva", "Jackson", "Scarlet", "Robert", "Anna", "Henry", "Ariel",
	}

	foundName, countName := searchBinaryWithName(listNames, name)
	if foundName {
		fmt.Printf("O item %s foi encontrado e a quantidade de tentativas foi %d\n", name, countName)
		return
	}
	fmt.Printf("O item NÃO foi encontrado e a quantidade de tentativas foi %d\n", countName)
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

func searchBinaryWithName(list []string, searchItem string) (bool, int) {
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
