package services

import (
	"fizzbuzz-api/entities"
	"fizzbuzz-api/models"
)

// StatsService a service dedicated to compute everything in relation to app's statistics
type StatsService struct {}

// GetMostAskedRequest gets the most asked request, casts it as a RequestModel and returns it
func (s *StatsService) GetMostAskedRequest() (*models.RequestModel, error) {
	request, err := (entities.Request{}).FindMostAskedRequest()

	if err != nil {
		return nil, err
	}

	return &models.RequestModel{
		Int1:  request.Int1,
		Int2: request.Int2,
		Limit: request.Limit,
		Str1: request.Str1,
		Str2: request.Str2,
	}, nil
}
