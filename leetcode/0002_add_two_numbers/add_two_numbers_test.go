package addtwonumbers

import (
	"os"
	"reflect"
	"testing"
)

func TestAddTwoNumbers1(t *testing.T) {
	testDatas := []struct {
		name     string
		arg1     *ListNode
		arg2     *ListNode
		expected *ListNode
	}{
		{
			name: "one",
			arg1: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
			arg2: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
			expected: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 8,
					},
				},
			},
		},
		{
			name: "two",
			arg1: &ListNode{
				Val:  5,
				Next: nil,
			},
			arg2: &ListNode{
				Val:  5,
				Next: nil,
			},
			expected: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
		{
			name: "three",
			arg1: &ListNode{
				Val:  1,
				Next: nil,
			},
			arg2: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
					},
				},
			},
			expected: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 1,
						},
					},
				},
			},
		},
		{
			name: "four",
			arg1: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
					},
				},
			},
			arg2: &ListNode{
				Val:  1,
				Next: nil,
			},
			expected: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 1,
						},
					},
				},
			},
		},
		{
			name: "five",
			arg1: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
			arg2: &ListNode{
				Val:  1,
				Next: nil,
			},
			expected: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
		{
			name: "six",
			arg1: &ListNode{
				Val:  1,
				Next: nil,
			},
			arg2: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
			expected: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	}

	for _, testData := range testDatas {
		t.Run(testData.name, func(t *testing.T) {
			if result := addTwoNumbers1(testData.arg1, testData.arg2); !reflect.DeepEqual(result, testData.expected) {
				t.Errorf("expected %v, got %v", testData.expected, result)
			}
		})
	}
}

func TestAddTwoNumbers2(t *testing.T) {
	testDatas := []struct {
		name     string
		arg1     *ListNode
		arg2     *ListNode
		expected *ListNode
	}{
		{
			name: "one",
			arg1: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
			arg2: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
			expected: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 8,
					},
				},
			},
		},
		{
			name: "two",
			arg1: &ListNode{
				Val:  5,
				Next: nil,
			},
			arg2: &ListNode{
				Val:  5,
				Next: nil,
			},
			expected: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
		{
			name: "three",
			arg1: &ListNode{
				Val:  1,
				Next: nil,
			},
			arg2: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
					},
				},
			},
			expected: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 1,
						},
					},
				},
			},
		},
		{
			name: "four",
			arg1: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
					},
				},
			},
			arg2: &ListNode{
				Val:  1,
				Next: nil,
			},
			expected: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 1,
						},
					},
				},
			},
		},
		{
			name: "five",
			arg1: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
			arg2: &ListNode{
				Val:  1,
				Next: nil,
			},
			expected: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
		{
			name: "six",
			arg1: &ListNode{
				Val:  1,
				Next: nil,
			},
			arg2: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
			expected: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	}

	for _, testData := range testDatas {
		t.Run(testData.name, func(t *testing.T) {
			if result := addTwoNumbers2(testData.arg1, testData.arg2); !reflect.DeepEqual(result, testData.expected) {
				t.Errorf("expected %v, got %v", testData.expected, result)
			}
		})
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
