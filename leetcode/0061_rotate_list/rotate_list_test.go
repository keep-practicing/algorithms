package rotatelist

import (
	"reflect"
	"testing"
)

func TestRotateList(t *testing.T) {
	testDatas := []struct {
		name     string
		arg      *ListNode
		k        int
		expected *ListNode
	}{
		{
			name: "one",
			arg: &ListNode{Val: 1, Next: &ListNode{
				Val: 2, Next: &ListNode{
					Val: 3, Next: nil,
				},
			}},
			k: 2,
			expected: &ListNode{Val: 2, Next: &ListNode{
				Val: 3, Next: &ListNode{
					Val: 1, Next: nil,
				},
			}},
		},
	}

	for _, testData := range testDatas {
		t.Run(testData.name, func(t *testing.T) {
			if result := rotateRight(testData.arg, testData.k); !reflect.DeepEqual(result, testData.expected) {
				t.Errorf("expected %v, got %v", testData.expected, result)
			}
		})
	}
}
