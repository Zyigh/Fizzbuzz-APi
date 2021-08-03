package models

// RequestModel jsonable struct representing the shape of a Request made to compute a FizzBuzz sent to the program that
// consumes the API for stats
type RequestModel struct {
	// Int1 Int1 the first int which multiples will be replaced by Str1
	Int1  uint64 `json:"int1"`
	// Int2 Int2 the second int which multiples will be replaced by Str2
	Int2  uint64 `json:"int2"`
	// Limit Limit the greatest int on which to compute FizzBuzz algortihm
	Limit uint64 `json:"limit"`
	// Str1 Str1 the string to replace multiples of Int1
	Str1  string `json:"str1"`
	// Str2 Str2 the string to replace multiples of Int2
	Str2  string `json:"str2"`

}
