package controllers

import (
	"fizzbuzz-api/facades"
	"fizzbuzz-api/models"
	"fizzbuzz-api/services"
	"log"
	"net/http"
)

// FizzbuzzController controller that handles every queries related to Fizzbuzz algorithm
// Composed with controllers.controller
type FizzbuzzController struct {
	// controller access to generic methods
	*controller
	// fizzbuzzService Service that handles Fizzbuzz related computations
	fizzbuzzService services.FizzbuzzService
}

// GetFizzbuzz tries to cast the FacadeInterface in Request's Context as GetFizzbuzzFacade, makes fizzbuzzService compute
// the data needed for these params and renders it
// Route /
func (f *FizzbuzzController) GetFizzbuzz() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)  {
		facade, ok := r.Context().Value(facades.FacadeContextKey).(facades.GetFizzbuzzFacade)

		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			f.render(models.ErrorModel{Message: "Couldn't read arguments"}, w)
			log.Println("GetFizzbuzzFacade not found in ctx")

			return
		}

		fizzbuzzData := f.fizzbuzzService.ComputeFizzbuzz(facade)

		f.render(fizzbuzzData, w)
	})
}

// GetFizzbuzzCtrl creates an instance of FizzbuzzController with FizzbuzzService loaded
func GetFizzbuzzCtrl() FizzbuzzController {
	return FizzbuzzController{
		fizzbuzzService: services.FizzbuzzService{},
	}
}
