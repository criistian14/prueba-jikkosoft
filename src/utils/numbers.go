package utils

import (
	projectErrors "github.com/criistian14/prueba-jikkosoft/src/errors"
)

// * Sort by algorithm QuickSort
func SortNumbers(numbers []int) ([]int, error) {
	// ! Not pass numbers
	if numbers == nil {
		return nil, projectErrors.ErrNumbersNull
	}

	// ! Empty numbers
	if len(numbers) < 1 {
		return nil, projectErrors.ErrNumbersEmpty
	}

	newArray := make([]int, len(numbers))
	for i, v := range numbers {
		newArray[i] = v
	}

	sort(newArray, 0, len(numbers)-1)
	newArray = separateDuplicates(newArray)

	return newArray, nil
}

// * Algorithm quicksort
func sort(arr []int, start int, end int) {
	if (end - start) < 1 {
		return
	}

	pivot := arr[end]
	splitIndex := start

	for i := start; i < end; i++ {
		if arr[i] < pivot {
			temp := arr[splitIndex]

			arr[splitIndex] = arr[i]
			arr[i] = temp

			splitIndex++
		}
	}

	arr[end] = arr[splitIndex]
	arr[splitIndex] = pivot

	sort(arr, start, splitIndex-1)
	sort(arr, splitIndex+1, end)
}

// * Separate the repeated ones and put them at the end
func separateDuplicates(arr []int) []int {
	isDuplicates := map[int]bool{}
	var duplicates []int
	var result []int

	for _, v := range arr {
		if isDuplicates[v] != true {
			isDuplicates[v] = true
			result = append(result, v)
		} else {
			duplicates = append(duplicates, v)
		}
	}

	result = append(result, duplicates...)

	return result
}
