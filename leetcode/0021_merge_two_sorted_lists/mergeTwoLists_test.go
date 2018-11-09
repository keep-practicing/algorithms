package mergetwolists

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	testDatas := []struct {
		name     string
		agr1     *ListNode
		arg2     *ListNode
		expected *ListNode
	}{
		{
			name: "one",
			agr1: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: nil}}},
			arg2: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 6, Next: nil}}},
			expected: &ListNode{
				Val: 1, Next: &ListNode{
					Val: 2, Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
								Next: &ListNode{
									Val:  6,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "two",
			agr1: &ListNode{Val: 1, Next: nil},
			arg2: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 6, Next: nil}}},
			expected: &ListNode{
				Val: 1, Next: &ListNode{
					Val: 2, Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
		{
			name:     "three",
			agr1:     nil,
			arg2:     nil,
			expected: nil,
		},
	}

	for _, testData := range testDatas {
		t.Run(testData.name, func(t *testing.T) {
			if result := mergeTwoLists(testData.agr1, testData.arg2); !reflect.DeepEqual(result, testData.expected) {
				t.Errorf("mergeTwoLists() = %v, expected %v", result, testData.expected)
			}
		})
	}
}
