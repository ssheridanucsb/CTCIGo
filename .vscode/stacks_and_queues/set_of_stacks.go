package stacksandqueues

import "errors"

type SetOfStacks struct {
	stacks       []*IntStack
	subStackSize int
}

func NewSetOfStacks(subStackSize int) *SetOfStacks {
	stackList := make([]*IntStack, 0, 1)
	return &SetOfStacks{stacks: stackList, subStackSize: subStackSize}
}

func (s *SetOfStacks) isEmpty() bool {
	return len(s.stacks) == 0
}

func (s *SetOfStacks) push(data int) {

	if len(s.stacks) == 0 {
		subStack := NewIntStack()
		subStack.push(data)
		s.stacks = append(s.stacks, subStack)
		return
	}

	added := false
	for _, stack := range s.stacks {
		if stack.size < s.subStackSize {
			stack.push(data)
			added = true
		}
	}

	if !added {
		subStack := NewIntStack()
		subStack.push(data)
		s.stacks = append(s.stacks, subStack)
	}
}

func (s *SetOfStacks) pop() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("empty stack set")
	}

	stack := s.stacks[len(s.stacks)-1]
	p, err := stack.pop()
	if err != nil {
		return 0, err
	}
	if stack.isEmpty() {
		s.stacks = s.stacks[:len(s.stacks)-1]
	}

	return p, nil
}

func (s *SetOfStacks) popAt(index int) (int, error) {
	if s.isEmpty() {
		return 0, errors.New("empty set of stacks")
	}

	if index >= len(s.stacks) {
		return 0, errors.New("substack does not exist")
	}

	stack := s.stacks[index]

	p, err := stack.pop()
	if err != nil {
		return 0, err
	}
	// keep ordering when removing
	// from total stack set without shifting everything over
	stack.size = s.subStackSize
	return p, nil
}
