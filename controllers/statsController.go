package controllers

import (
	"fizzbuzz-api/models"
	"fizzbuzz-api/services"
	"log"
	"net/http"
)

// StatsController controller that handles every queries related to app statistics
// Composed with controllers.controller
type StatsController struct {
	// controller access to generic methods
	*controller
	// statsService Service that handles statistics related computation
	statsService services.StatsService
}

// GetMostAskedRequest compute most asked Fizzbuzz Request and renders it
// Route /stats/most-frequent-request
func (s *StatsController) GetMostAskedRequest() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)  {
		statsData, err := s.statsService.GetMostAskedRequest()

		if err != nil {
			log.Println(err)
			s.render(models.ErrorModel{Message: "Something went wrong"}, w)

			return
		}

		s.render(statsData, w)
	})
}

// GetStatsController creates an instance of StatsController with StatsService loaded
func GetStatsController() StatsController {
	return StatsController{
		statsService: services.StatsService{},
	}
}
