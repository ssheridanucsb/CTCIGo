package stacksandqueues

import "testing"

func TestAnimalQ(t *testing.T) {
	q := NewAnimalQ()

	q.enqueueAnimal("Finn", true)
	q.enqueueAnimal("Charlotte", false)
	q.enqueueAnimal("Zeke", false)
	q.enqueueAnimal("Cody", true)

	c, err := q.dequeueCat()
	if err != nil {
		t.Error("expected no error")
	}

	if c != "Charlotte" {
		t.Errorf("expected Charlotte but recieved %s", c)
	}

	d, _, _ := q.dequeue()
	if d != "Finn" {
		t.Errorf("expected Finn but recieved %s", d)
	}

	d, _ = q.dequeueDog()
	if d != "Cody" {
		t.Errorf("expected Cody but recieved %s", d)
	}

	c, _ = q.dequeueCat()
	if c != "Zeke" {
		t.Errorf("expected Zeke but recieved %s", c)
	}

	if !q.isEmpy() {
		t.Error("expected empty queue")
	}

}
