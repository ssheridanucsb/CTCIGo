package stacksandqueues

import "errors"

type AnimalQueueIntNode struct {
	name            string //name of animal
	dog             bool   //true if dog, false if cat
	next            *AnimalQueueIntNode
	nextSameSpecies *AnimalQueueIntNode
	prev            *AnimalQueueIntNode
}

func NewAnimalQueueIntNode(name string, dog bool) *AnimalQueueIntNode {
	return &AnimalQueueIntNode{name: name, dog: dog, next: nil, nextSameSpecies: nil, prev: nil}
}

type AnimalQ struct {
	first *AnimalQueueIntNode
	last  *AnimalQueueIntNode

	firstDog *AnimalQueueIntNode
	lastDog  *AnimalQueueIntNode

	firstCat *AnimalQueueIntNode
	lastCat  *AnimalQueueIntNode
}

func NewAnimalQ() *AnimalQ {
	return &AnimalQ{first: nil, last: nil, firstDog: nil, lastDog: nil, firstCat: nil, lastCat: nil}
}

func (q *AnimalQ) isEmpy() bool {
	return q.first == nil
}

func (q *AnimalQ) enqueueAnimal(name string, dog bool) {
	node := NewAnimalQueueIntNode(name, dog)

	//add to back of q
	if q.last != nil {
		node.prev = q.last
		q.last.next = node
	}

	q.last = node

	if q.first == nil {
		q.first = q.last
	}

	if dog {
		if q.lastDog != nil {
			q.lastDog.nextSameSpecies = node
		}
		q.lastDog = node
		if q.firstDog == nil {
			q.firstDog = q.lastDog
		}

	} else {
		if q.lastCat != nil {
			q.lastCat.nextSameSpecies = node
		}
		q.lastCat = node
		if q.firstCat == nil {
			q.firstCat = q.lastCat
		}
	}
}

func (q *AnimalQ) dequeue() (string, bool, error) {
	if q.isEmpy() {
		return "", false, errors.New("empty q")
	}
	dq := q.first
	q.first = q.first.next
	q.first.prev = nil
	if dq.dog {
		q.firstDog = dq.nextSameSpecies
	} else {
		q.firstCat = dq.nextSameSpecies
	}

	if q.first == nil {
		q.last = nil
		q.lastDog = nil
		q.lastCat = nil
	}

	return dq.name, dq.dog, nil
}

func (q *AnimalQ) dequeueDog() (string, error) {
	d := q.firstDog
	if d == nil {
		return "", errors.New("no dogs in q")
	}

	q.firstDog = d.nextSameSpecies
	if q.firstDog == nil {
		q.lastDog = nil
	}

	if d.prev != nil {
		d.prev.next = d.next
		if d.next != nil {
			d.next.prev = d.prev
		}
	} else {
		// first element
		q.first = d.next
		if q.first == nil {
			q.last = nil
		}

	}

	return d.name, nil
}

func (q *AnimalQ) dequeueCat() (string, error) {
	c := q.firstCat
	if c == nil {
		return "", errors.New("no cats in q")
	}

	q.firstCat = c.nextSameSpecies
	if q.firstCat == nil {
		q.lastCat = nil
	}

	if c.prev != nil {
		c.prev.next = c.next
		if c.next != nil {
			c.next.prev = c.prev
		}
	} else {
		// first element
		q.first = c.next
		if q.first == nil {
			q.last = nil
		}
	}

	return c.name, nil
}
