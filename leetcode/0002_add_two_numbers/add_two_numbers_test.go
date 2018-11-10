package addtwonumbers

import "testing"

func TestAddTwoNumbers1(t *testing.T) {
	var (
		liCase  map[string][]int
		liCases []map[string][]int
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

	// case 5
	liCase["l1"] = []int{7}
	liCase["l2"] = []int{8, 9}
	liCase["expected"] = []int{5, 0, 1}
	liCases = append(liCases, liCase)

	// case 6
	liCase["l1"] = []int{1, 9, 5, 6, 7}
	liCase["l2"] = []int{7}
	liCase["expected"] = []int{8, 9, 5, 6, 7}
	liCases = append(liCases, liCase)

	// case 7
	liCase["l1"] = []int{7}
	liCase["l2"] = []int{1, 9, 5, 6, 7}
	liCase["expected"] = []int{8, 9, 5, 6, 7}
	liCases = append(liCases, liCase)

	// case 8
	liCase["l1"] = []int{3, 9, 5, 6, 7}
	liCase["l2"] = []int{7}
	liCase["expected"] = []int{0, 0, 6, 6, 7}
	liCases = append(liCases, liCase)

	for i := 0; i < len(liCases); i++ {
		l1 := sliToList(liCases[i]["l1"])
		l2 := sliToList(liCases[i]["l2"])
		newL1 := addTwoNumbers1(l1, l2)
		for k := 0; k < len(liCases[i]["expected"]); k++ {
			if liCases[i]["expected"][k] != newL1.Val {
				t.Error("测试不通过")
			}
			newL1 = newL1.Next
		}
		l1 = sliToList(liCases[i]["l1"])
		l2 = sliToList(liCases[i]["l2"])
		newL2 := addTwoNumbers2(l1, l2)
		for k := 0; k < len(liCases[i]["expected"]); k++ {
			if liCases[i]["expected"][k] != newL2.Val {
				t.Error("测试不通过")
			}
			newL2 = newL2.Next
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
