package usecase

type NumberUsecase interface {
	SortArrayNumbers(numbers []int) ([]int, error)
}

type numberUsecase struct {
}

// * Create and return new instance
func NewNumberUsecase() *numberUsecase {
	return &numberUsecase{}
}
