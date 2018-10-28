package addTwoNumbers

import "testing"

func TestAddTwoNumbers1(t *testing.T) {
	var (
		liCase   map[string][]int
		liCases  []map[string][]int
		testFunc = [2]func(l1 *ListNode, l2 *ListNode) *ListNode{addTwoNumbers1, addTwoNumbers2}
	)

	// case 1
	liCase = make(map[string][]int)
	liCase["l1"] = []int{2, 4, 3}
	liCase["l2"] = []int{5, 6, 4}
	liCase["expected"] = []int{7, 0, 8}
	liCases = append(liCases, liCase)

	// case 2
	liCase["l1"] = []int{5}
	liCase["l2"] = []int{5}
	liCase["expected"] = []int{0, 1}
	liCases = append(liCases, liCase)

	// case 3
	liCase["l1"] = []int{1}
	liCase["l2"] = []int{9, 9, 9}
	liCase["expected"] = []int{0, 0, 0, 1}
	liCases = append(liCases, liCase)

	// case 4
	liCase["l1"] = []int{9, 9}
	liCase["l2"] = []int{1}
	liCase["expected"] = []int{0, 0, 1}
	liCases = append(liCases, liCase)

	for i := 0; i < len(liCases); i++ {
		l1 := sliToList(liCases[i]["l1"])
		l2 := sliToList(liCases[i]["l2"])
		for j := 0; j < len(testFunc); j++ {
			newL := testFunc[j](l1, l2)
			for k := 0; k < len(liCases[i]["expected"]); k++ {
				if liCases[i]["expected"][k] != newL.Val {
					t.Error("测试不通过")
				}
				newL = newL.Next
			}
		}
	}
}

func sliToList(sli []int) *ListNode {
	li := ListNode{Val: -1}
	cur := &li
	for i := 0; i < len(sli); i++ {
		cur.Next = &ListNode{Val: sli[i]}
		cur = cur.Next
	}
	return li.Next
}
