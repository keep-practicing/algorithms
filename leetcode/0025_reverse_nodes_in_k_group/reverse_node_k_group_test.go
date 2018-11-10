package reversenodesinkgroup

import (
	"reflect"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name     string
		args     args
		expected *ListNode
	}{
		{
			name: "one",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  5,
									Next: nil,
								},
							},
						},
					},
				},
				k: 2,
			},
			expected: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
		},
		{
			name: "two",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  5,
									Next: nil,
								},
							},
						},
					},
				},
				k: 3,
			},
			expected: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
		},

		{
			name: "three",
			args: args{
				head: &ListNode{
					Val:  1,
					Next: nil,
				},
				k: 2,
			},
			expected: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}
	for _, data := range tests {
		t.Run(data.name, func(t *testing.T) {
			if result := reverseKGroup(data.args.head, data.args.k); !reflect.DeepEqual(result, data.expected) {
				t.Errorf("reverseKGroup() = %v, expected %v", result, data.expected)
			}
		})
	}
}
