package pushswap

type Stacks struct {
	A []int
	B []int
}

func NewStacks(a []int) *Stacks {
	stackA := make([]int, len(a))
	copy(stackA, a)
	return &Stacks{
		A: stackA,
		B: []int{},
	}
}

func (s *Stacks) IsSorted() bool {
	if len(s.B) != 0 {
		return false
	}
	for i := 0; i < len(s.A)-1; i++ {
		if s.A[i] > s.A[i+1] {
			return false
		}
	}
	return true
}
