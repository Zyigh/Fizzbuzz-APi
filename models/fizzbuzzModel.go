package models

// FizzbuzzModel jsonable struct representing the shape of Fizzbuzz algorithm results sent to the program consuming the
// API
type FizzbuzzModel struct {
	// Data the json key that will contain the result of the Fizzbuzz algorithm computed with user parameters that will
	// be consumed by the program consuming the API
	Data []string `json:"data"`
}
