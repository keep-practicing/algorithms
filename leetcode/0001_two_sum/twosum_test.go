package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	if res := twoSum(nums, target); !reflect.DeepEqual(res, []int{0, 1}) {
		t.Error("Failed, two sum")
	}
}
