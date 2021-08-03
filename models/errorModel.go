package models

// ErrorModel jsonable struct representing the shape of error messages sent to the program consuming the API
type ErrorModel struct {
	// Message the json key that will contain any error message that will be consumed by the program consuming the API
	Message string `json:"message"`
}
