package services

import (
	"fizzbuzz-api/entities"
	"fizzbuzz-api/facades"
	"fizzbuzz-api/models"
	"log"
	"strconv"
)

// FizzbuzzService a service dedicated to compute everything in relation to Fizzbuzz
type FizzbuzzService struct {}

// ComputeFizzbuzz compute the FizzBuzz algorithm (https://en.wikipedia.org/wiki/Fizz_buzz) with a GetFizzbuzzFacade to
// determine which number to replace by which string in which limit.
// It stores the GetFizzbuzzFacade as a Request entity to the database asynchronously
// It returns a FizzbuzzModel with all values to the controller that will render it
//
// Note that any problem with recording the Request object won't prevent this method to compute the Fizzbuzz result
// and so won't interfere with the consuming application
func (f *FizzbuzzService) ComputeFizzbuzz(facade facades.GetFizzbuzzFacade) models.FizzbuzzModel {
	go func(facade facades.GetFizzbuzzFacade) {
		request := entities.NewRequest(facade.Int1, facade.Int2, facade.Limit, facade.Str1, facade.Str2)
		err := request.Save()

		if err != nil {
			log.Println(err)
		}
	}(facade)

	var values []string

	for i := uint64(1); i <= facade.Limit; i++ {
		var value string

		if i % facade.Int1 == 0 {
			value += facade.Str1
		}

		if i % facade.Int2 == 0 {
			value += facade.Str2
		}

		if len(value) == 0 {
			value = strconv.FormatUint(i, 10)
		}

		values = append(values, value)
	}

	return models.FizzbuzzModel{Data: values}
}
