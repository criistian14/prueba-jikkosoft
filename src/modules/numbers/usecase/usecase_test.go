package usecase_test

import (
	projectErrors "github.com/criistian14/prueba-jikkosoft/src/errors"
	numberUseCase "github.com/criistian14/prueba-jikkosoft/src/modules/numbers/usecase"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberUsecase_SortArrayNumbers(t *testing.T) {
	numbers := []int{123, 4, 5, 10, 5, 14, 4}
	numberUsecase := numberUseCase.NewNumberUsecase()

	t.Run("success", func(t *testing.T) {
		expectedSorted := []int{4, 5, 10, 14, 123, 4, 5}

		sorted, err := numberUsecase.SortArrayNumbers(numbers)
		if err != nil {
			t.Fatalf("Error Validate Numbers: %s", err)
		}

		assert.Equal(t, expectedSorted, sorted)
	})

	// ! Errors
	t.Run("error-numbers-empty", func(t *testing.T) {
		numbers := make([]int, 0)
		_, err := numberUsecase.SortArrayNumbers(numbers)

		assert.EqualError(t, err, projectErrors.ErrNumbersEmpty.Error())
	})

	t.Run("error-numbers-null", func(t *testing.T) {
		_, err := numberUsecase.SortArrayNumbers(nil)

		assert.EqualError(t, err, projectErrors.ErrNumbersNull.Error())
	})
}
