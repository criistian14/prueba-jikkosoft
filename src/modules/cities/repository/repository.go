package repository

import cityDomain "github.com/criistian14/prueba-jikkosoft/src/modules/cities/domain"

type CityRepository interface {
	FindCity(city cityDomain.City, isDeleted bool) (cityDomain.City, error)
	SaveCity(city cityDomain.City) (cityDomain.City, error)
}

type cityRepository struct {
}

// * Create and return new instance
func NewCityRepository() *cityRepository {
	return &cityRepository{}
}
