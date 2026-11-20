package pushswap

import "sort"

// Solve calcule et retourne la suite d'instructions pour trier la pile A.
func Solve(initial []int) []string {
	s := NewStacks(initial)
	if s.IsSorted() {
		return []string{}
	}

	s.A = normalize(s.A)

	var ops []string
	record := func(op string, fn func()) {
		fn()
		ops = append(ops, op)
	}

	switch len(s.A) {
	case 2:
		if s.A[0] > s.A[1] {
			record("sa", s.Sa)
		}
	case 3:
		ops = sortThree(s)
	case 4, 5:
		ops = sortSmall(s)
	default:
		ops = radixSort(s)
	}

	return ops
}

// normalize remplace chaque valeur par son rang trié (0 à n-1).
func normalize(arr []int) []int {
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	sort.Ints(sorted)

	rank := make(map[int]int, len(arr))
	for i, v := range sorted {
		rank[v] = i
	}

	res := make([]int, len(arr))
	for i, v := range arr {
		res[i] = rank[v]
	}
	return res
}

func sortThree(s *Stacks) []string {
	var ops []string
	rec := func(op string, fn func()) {
		fn()
		ops = append(ops, op)
	}

	for !aIsSorted(s) {
		a, b, c := s.A[0], s.A[1], s.A[2]
		if a > b && b < c && a < c {
			rec("sa", s.Sa)
		} else if a > b && b > c {
			rec("sa", s.Sa)
			rec("rra", s.Rra)
		} else if a > b && b < c && a > c {
			rec("ra", s.Ra)
		} else if a < b && b > c && a < c {
			rec("sa", s.Sa)
			rec("ra", s.Ra)
		} else {
			rec("rra", s.Rra)
		}
	}
	return ops
}

func sortSmall(s *Stacks) []string {
	var ops []string
	rec := func(op string, fn func()) {
		fn()
		ops = append(ops, op)
	}

	targetPushes := len(s.A) - 3
	for pushed := 0; pushed < targetPushes; pushed++ {
		minIdx := 0
		for i := 1; i < len(s.A); i++ {
			if s.A[i] < s.A[minIdx] {
				minIdx = i
			}
		}

		if minIdx <= len(s.A)/2 {
			for i := 0; i < minIdx; i++ {
				rec("ra", s.Ra)
			}
		} else {
			for i := 0; i < len(s.A)-minIdx; i++ {
				rec("rra", s.Rra)
			}
		}

		if s.IsSorted() && len(s.B) == 0 {
			return ops
		}
		rec("pb", s.Pb)
	}

	ops = append(ops, sortThree(s)...)

	for len(s.B) > 0 {
		rec("pa", s.Pa)
	}
	return ops
}

func radixSort(s *Stacks) []string {
	var ops []string
	rec := func(op string, fn func()) {
		fn()
		ops = append(ops, op)
	}

	n := len(s.A)
	maxVal := n - 1
	maxBits := 0
	for (maxVal >> maxBits) > 0 {
		maxBits++
	}

	for bit := 0; bit < maxBits; bit++ {
		if s.IsSorted() {
			break
		}
		for i := 0; i < n; i++ {
			if (s.A[0]>>bit)&1 == 0 {
				rec("pb", s.Pb)
			} else {
				rec("ra", s.Ra)
			}
		}
		for len(s.B) > 0 {
			rec("pa", s.Pa)
		}
	}

	return ops
}
func aIsSorted(s *Stacks) bool {
	for i := 0; i < len(s.A)-1; i++ {
		if s.A[i] > s.A[i+1] {
			return false
		}
	}
	return true
}
