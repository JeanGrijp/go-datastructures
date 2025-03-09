package stack

type Node struct {
	Value any
	Next  *Node
}

type Stack struct {
	first *Node
}

func (s *Stack) Push(value any) {
	s.first = &Node{Value: value, Next: s.first}
}

func (s *Stack) Pop() (any, bool) {
	if s.first == nil {
		return nil, false
	} else {
		value := s.first.Value
		s.first = s.first.Next
		return value, true
	}
}

func (s *Stack) Values() any {
	if s.first == nil {
		return nil
	}
	var values []any
	for aux := s.first; aux != nil; aux = aux.Next {
		values = append(values, aux.Value)
	}
	return values
}
