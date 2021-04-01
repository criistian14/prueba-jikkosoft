package usecase

import "github.com/criistian14/prueba-jikkosoft/src/utils"

func (usecase *numberUsecase) SortArrayNumbers(numbers []int) ([]int, error) {
	return utils.SortNumbers(numbers)
}
