package utils_test

import (
	projectErrors "github.com/criistian14/prueba-jikkosoft/src/errors"
	"github.com/criistian14/prueba-jikkosoft/src/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortNumbers(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		numbers := []int{123, 4, 5, 10, 5, 14, 4}
		expectedSorted := []int{4, 5, 10, 14, 123, 4, 5}

		sorted, err := utils.SortNumbers(numbers)
		if err != nil {
			t.Fatalf("Error Sorting Numbers: %s", err)
		}

		assert.Equal(t, expectedSorted, sorted)
	})

	// ! Errors
	t.Run("error-numbers-empty", func(t *testing.T) {
		numbers := make([]int, 0)
		_, err := utils.SortNumbers(numbers)

		assert.EqualError(t, err, projectErrors.ErrNumbersEmpty.Error())
	})

	t.Run("error-numbers-null", func(t *testing.T) {
		_, err := utils.SortNumbers(nil)

		assert.EqualError(t, err, projectErrors.ErrNumbersNull.Error())
	})
}
