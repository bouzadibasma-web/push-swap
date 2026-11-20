package pushswap

import "testing"

func applyOps(initial []int, ops []string) *Stacks {
	s := NewStacks(initial)
	for _, op := range ops {
		switch op {
		case "pa":
			s.Pa()
		case "pb":
			s.Pb()
		case "sa":
			s.Sa()
		case "sb":
			s.Sb()
		case "ss":
			s.Ss()
		case "ra":
			s.Ra()
		case "rb":
			s.Rb()
		case "rr":
			s.Rr()
		case "rra":
			s.Rra()
		case "rrb":
			s.Rrb()
		case "rrr":
			s.Rrr()
		}
	}
	return s
}

func TestParseArgsValid(t *testing.T) {
	got, err := ParseArgs([]string{"  2   1  ", "3", "-5"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []int{2, 1, 3, -5}
	if len(got) != len(want) {
		t.Fatalf("got len %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got[%d] = %d, want %d", i, got[i], want[i])
		}
	}
}

func TestParseArgsErrors(t *testing.T) {
	cases := [][]string{
		{"0 one 2 3"},
		{"1 2 2 3"},
		{"-10", "-10"},
		{"abc"},
	}
	for _, tc := range cases {
		if _, err := ParseArgs(tc); err == nil {
			t.Errorf("expected error for input %v, got nil", tc)
		}
	}
}

func TestOperationsBasic(t *testing.T) {
	s := NewStacks([]int{2, 1, 3})
	s.Sa()
	if s.A[0] != 1 || s.A[1] != 2 {
		t.Fatalf("Sa failed: got %v", s.A)
	}

	s.Pb()
	if len(s.B) != 1 || s.B[0] != 1 {
		t.Fatalf("Pb failed: B=%v", s.B)
	}

	s.Ra()
	if s.A[0] != 3 || s.A[1] != 2 {
		t.Fatalf("Ra failed: A=%v", s.A)
	}

	s.Rra()
	if s.A[0] != 2 || s.A[1] != 3 {
		t.Fatalf("Rra failed: A=%v", s.A)
	}

	s.Pa()
	if len(s.B) != 0 || s.A[0] != 1 {
		t.Fatalf("Pa failed: A=%v B=%v", s.A, s.B)
	}
}

func TestSolveSorting(t *testing.T) {
	tests := [][]int{
		{},
		{42},
		{1, 2, 3, 4, 5},
		{2, 1},
		{3, 2, 1},
		{2, 1, 3, 6, 5, 8},
		{-50, 100, 0, -1, 42, 7, -100, 25},
		{4, 3, 2, 1},
		{5, 4, 3, 2, 1},
	}

	for _, input := range tests {
		ops := Solve(input)
		res := applyOps(input, ops)
		if !res.IsSorted() {
			t.Errorf("Solve(%v) did not sort properly; final A=%v B=%v ops=%v",
				input, res.A, res.B, ops)
		}
	}
}
func TestIsSortedWithNonEmptyB(t *testing.T) {
	s := NewStacks([]int{1, 2, 3})
	s.B = []int{0}
	if s.IsSorted() {
		t.Fatalf("expected IsSorted() to be false when B is not empty")
	}
}
