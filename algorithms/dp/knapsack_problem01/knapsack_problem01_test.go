package knapsackproblem01

import "testing"

func TestKnapsackProblem01(t *testing.T) {
	type Knapsack struct {
		w []int
		v []int
		c int
	}
	testData := []Knapsack{
		Knapsack{
			w: []int{1, 2, 3},
			v: []int{6, 10, 12},
			c: 5,
		},
		Knapsack{
			w: []int{},
			v: []int{},
			c: 5,
		},
	}

	expectedData := []int{22, 0}

	for index, data := range testData {
		if res := knapsack01(data.w, data.v, data.c); res != expectedData[index] {
			t.Errorf("expected %d, got %d", expectedData[index], res)
		}
	}

	knapsack := Knapsack{
		w: []int{1, 2, 3, 3},
		v: []int{6, 10, 12},
		c: 5,
	}

	defer func() {
		if err := recover(); err == nil {
			t.Error("Failed, panic error can't catch")
		}
	}()
	knapsack01(knapsack.w, knapsack.v, knapsack.c)
}
