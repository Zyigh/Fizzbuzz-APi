package services_test

import (
	"fizzbuzz-api/facades"
	"fizzbuzz-api/services"
	"strings"
	"testing"
)

// TestFizzbuzzService_ComputeFizzbuzzLengthIsLimit asserts ComputeFizzbuzz returns an array of length defined by limit
func TestFizzbuzzService_ComputeFizzbuzzLengthIsLimit(t *testing.T) {
	service := &services.FizzbuzzService{}
	facade := facades.GetFizzbuzzFacade{
		Int1:  facades.DefaultInt1,
		Int2:  facades.DefaultInt2,
		Limit: 300,
		Str1:  facades.DefaultStr1,
		Str2:  facades.DefaultStr2,
	}

	model := service.ComputeFizzbuzz(facade)

	if len(model.Data) != int(facade.Limit) {
		t.Errorf("Invalid Fizzbuzz length, expected %d, got %d", facade.Limit, len(model.Data))
	}
}

// TestFizzbuzzService_ComputeFizzbuzzInt1IsReplacedByStr1 asserts ComputeFizzbuzz replace int1 multiples with str1
func TestFizzbuzzService_ComputeFizzbuzzInt1IsReplacedByStr1(t *testing.T) {
	service := &services.FizzbuzzService{}
	facade := facades.GetFizzbuzzFacade{
		Int1:  2,
		Int2:  facades.DefaultInt2,
		Limit: facades.DefaultLimit,
		Str1:  "test",
		Str2:  facades.DefaultStr2,
	}

	model := service.ComputeFizzbuzz(facade)

	for i, str := range model.Data {
		if (i + 1) % int(facade.Int1) == 0 {
			if !strings.HasPrefix(str, facade.Str1) {
				t.Errorf("Expected string starting with %s, got %s", facade.Str1, str)
			}
		}
	}
}

// TestFizzbuzzService_ComputeFizzbuzzInt2IsReplacedByStr2 asserts ComputeFizzbuzz replace int2 multiples with str2
func TestFizzbuzzService_ComputeFizzbuzzInt2IsReplacedByStr2(t *testing.T) {
	service := &services.FizzbuzzService{}
	facade := facades.GetFizzbuzzFacade{
		Int1:  facades.DefaultInt1,
		Int2:  8,
		Limit: facades.DefaultLimit,
		Str1:  facades.DefaultStr1,
		Str2:  "test",
	}

	model := service.ComputeFizzbuzz(facade)

	for i, str := range model.Data {
		if (i + 1) % int(facade.Int2) == 0 {
			if !strings.HasSuffix(str, facade.Str2) {
				t.Errorf("Expected string starting with %s, got %s", facade.Str2, str)
			}
		}
	}
}
