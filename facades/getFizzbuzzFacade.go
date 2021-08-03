package facades

import (
	"net/http"
	"strconv"
)

// GetFizzbuzzFacade facade that represents Request's params to generate Fizzbuzz algortihm values
type GetFizzbuzzFacade struct {
	// Int1 the first int which multiples will be replaced by Str1
	Int1 uint64
	// Int2 the second int which multiples will be replaced by Str2
	Int2 uint64
	// Limit the greatest int on which to compute FizzBuzz algortihm
	Limit uint64
	// Str1 the string to replace multiples of Int1
	Str1 string
	// Str2 the string to replace multiples of Int2
	Str2 string
}

const (
	// DefaultInt1 default value for Int1
	DefaultInt1 = 3
	// DefaultInt2 default value for Int2
	DefaultInt2 = 5
	// DefaultLimit default value for Limit
	DefaultLimit = 100
	// DefaultStr1 default value for Str1
	DefaultStr1 = "Fizz"
	// DefaultStr2 default value for Str2
	DefaultStr2 = "Buzz"
)

// parseUint simple parser to convert Request int params that replaces empty values by 0
func (f *GetFizzbuzzFacade) parseUint(str string) (uint64, error) {
	if len(str) == 0 {
		return 0, nil
	}

	return strconv.ParseUint(str, 10, 64)
}

// Deserialize populate instance with Request params if possible, returns an error if params are invalid
// Non specified values or 0 values for int are replaced by default values defined as consts
func (f GetFizzbuzzFacade) Deserialize(r *http.Request) (FacadeInterface, error) {
	int1, err := f.parseUint(r.URL.Query().Get("int1"))
	if err != nil {
		return f, err
	}

	int2, err := f.parseUint(r.URL.Query().Get("int2"))
	if err != nil {
		return f, err
	}

	limit, err := f.parseUint(r.URL.Query().Get("limit"))
	if err != nil {
		return f, err
	}

	str1 := r.URL.Query().Get("str1")
	str2 := r.URL.Query().Get("str2")

	if int1 == 0 {
		int1 = DefaultInt1
	}

	if int2 == 0 {
		int2 = DefaultInt2
	}

	if limit == 0 {
		limit = DefaultLimit
	}

	if str1 == "" {
		str1 = DefaultStr1
	}

	if str2 == "" {
		str2 = DefaultStr2
	}

	f.Int1 = int1
	f.Int2 = int2
	f.Limit = limit
	f.Str1 = str1
	f.Str2 = str2

	return f, nil
}
