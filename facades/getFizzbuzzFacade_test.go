package facades_test

import (
	"fizzbuzz-api/facades"
	"fmt"
	"net/http"
	"testing"
)

// TestFizzbuzzFacade_DeserializeWithoutParams asserts Deserialize returns GetFizzbuzzFacade with default values when call
// with http.Request without params
func TestGetFizzbuzzFacade_DeserializeWithoutParams(t *testing.T) {
	facade := facades.GetFizzbuzzFacade{}

	request, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Errorf("test is invalid, should have *http.Request got err %s", err)
	}

	fi1, err := facade.Deserialize(request)

	if err != nil {
		t.Errorf("Request without params should make facade with default values")
	}

	f, ok := fi1.(facades.GetFizzbuzzFacade)

	if !ok {
		t.Errorf("GetFizzbuzzFacade.Deserialize method should return a GetFizzbuzzFacade instance")
	}

	if f.Int1 != facades.DefaultInt1 ||
		f.Int2 != facades.DefaultInt2 ||
		f.Limit != facades.DefaultLimit ||
		f.Str1 != facades.DefaultStr1 ||
		f.Str2 != facades.DefaultStr2 {
		t.Errorf("GetFizzbuzzFacade instanciated without params should have default values")
	}
}

// TestFizzbuzzFacade_DeserializeWithPartialParams asserts Deserialize returns GetFizzbuzzFacade with values specified in
// http.Request
func TestGetFizzbuzzFacade_DeserializeWithPartialParams(t *testing.T) {
	facade := facades.GetFizzbuzzFacade{}

	definedInt1  := uint64(8)
	definedLimit := uint64(42)
	definedStr1  := "Marvin"

	url := fmt.Sprintf("/?int1=%d&limit=%d&str1=%s", definedInt1, definedLimit, definedStr1)
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		t.Errorf("test is invalid, should have *http.Request got err %s", err)
	}

	facadeInterface, err := facade.Deserialize(request)

	if err != nil {
		t.Errorf("Request without params should make facade with default values")
	}

	f, ok := facadeInterface.(facades.GetFizzbuzzFacade)

	if !ok {
		t.Errorf("GetFizzbuzzFacade.Deserialize method should return a GetFizzbuzzFacade instance")
	}

	if f.Int1 != definedInt1 ||
		f.Limit != definedLimit ||
		f.Str1 != definedStr1 {
		t.Errorf(
			"params in url %s should be found in GetFizzbuzzFacade, expected %d, %d, %s, got %d, %d, %s",
			url,
			definedInt1,
			definedLimit,
			definedStr1,
			f.Int1,
			f.Limit,
			f.Str1,
			)
	}

	if f.Int2 != facades.DefaultInt2 ||
		f.Str2 != facades.DefaultStr2 {
		t.Errorf("params not specified in url should lead to default values in GetFizzbuzzFacade")
	}
}

// TestFizzbuzzFacade_DeserializeWithWrongParams asserts Deserialize fails when http.Request params aren't correct
func TestGetFizzbuzzFacade_DeserializeWithWrongParams(t *testing.T) {
	facade := facades.GetFizzbuzzFacade{}

	definedInt1  := -12
	definedLimit := "fourty two"
	definedStr1  := "Arthur Dent"

	url := fmt.Sprintf("/?int1=%d&limit=%s&str1=%s", definedInt1, definedLimit, definedStr1)
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		t.Errorf("test is invalid, should have *http.Request got err %s", err)
	}

	_, err = facade.Deserialize(request)

	if err == nil {
		t.Errorf("Non castable params should make deserialize returns error")
	}
}
