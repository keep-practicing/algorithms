package lettercombinationsofaphonenumber

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	digits := "23"
	expectedData := []string{
		"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf",
	}

	if res := letterCombinations(digits); !reflect.DeepEqual(res, expectedData) {
		t.Errorf("expected %v, got %v", expectedData, res)
	}
}
