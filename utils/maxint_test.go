package utils

import "testing"

func TestCalcMaxInt(t *testing.T) {
	testData := [][]int{
		{},
		{3, 4, 67, 8},
	}
	expectedData := []int{-1, 67}

	for index, data := range testData {
		if res, _ := CalcMaxInt(data...); res != expectedData[index] {
			t.Errorf("expected %d, got %d", expectedData[index], res)
		}
	}
}
