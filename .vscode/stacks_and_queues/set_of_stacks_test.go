package stacksandqueues

import "testing"

func TestSetOfStacks(t *testing.T) {
	set := NewSetOfStacks(2)
	set.push(1)
	set.push(2)
	set.push(3)
	set.push(4)
	set.push(5)

	for i := 5; i > 0; i-- {
		p, err := set.pop()
		if err != nil {
			t.Error("Threw unexpected error")
		}
		if p != i {
			t.Errorf("expected %d, recieved %d", i, p)
		}
	}
}
