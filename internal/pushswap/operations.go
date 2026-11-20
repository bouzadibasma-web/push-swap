package pushswap

// Sa échange les 2 premiers éléments du sommet de la pile A.
func (s *Stacks) Sa() {
	if len(s.A) < 2 {
		return
	}
	s.A[0], s.A[1] = s.A[1], s.A[0]
}

// Sb échange les 2 premiers éléments du sommet de la pile B.
func (s *Stacks) Sb() {
	if len(s.B) < 2 {
		return
	}
	s.B[0], s.B[1] = s.B[1], s.B[0]
}

// Ss exécute Sa et Sb en même temps.
func (s *Stacks) Ss() {
	s.Sa()
	s.Sb()
}

// Pa prend le premier élément au sommet de B et le met au sommet de A.
func (s *Stacks) Pa() {
	if len(s.B) == 0 {
		return
	}
	top := s.B[0]
	s.B = s.B[1:]
	s.A = append([]int{top}, s.A...)
}

// Pb prend le premier élément au sommet de A et le met au sommet de B.
func (s *Stacks) Pb() {
	if len(s.A) == 0 {
		return
	}
	top := s.A[0]
	s.A = s.A[1:]
	s.B = append([]int{top}, s.B...)
}

// Ra décale tous les éléments de A vers le haut (le 1er devient le dernier).
func (s *Stacks) Ra() {
	if len(s.A) < 2 {
		return
	}
	first := s.A[0]
	s.A = append(s.A[1:], first)
}

// Rb décale tous les éléments de B vers le haut (le 1er devient le dernier).
func (s *Stacks) Rb() {
	if len(s.B) < 2 {
		return
	}
	first := s.B[0]
	s.B = append(s.B[1:], first)
}

// Rr exécute Ra et Rb en même temps.
func (s *Stacks) Rr() {
	s.Ra()
	s.Rb()
}

// Rra décale tous les éléments de A vers le bas (le dernier devient le 1er).
func (s *Stacks) Rra() {
	if len(s.A) < 2 {
		return
	}
	last := s.A[len(s.A)-1]
	s.A = append([]int{last}, s.A[:len(s.A)-1]...)
}

// Rrb décale tous les éléments de B vers le bas (le dernier devient le 1er).
func (s *Stacks) Rrb() {
	if len(s.B) < 2 {
		return
	}
	last := s.B[len(s.B)-1]
	s.B = append([]int{last}, s.B[:len(s.B)-1]...)
}

// Rrr exécute Rra et Rrb en même temps.
func (s *Stacks) Rrr() {
	s.Rra()
	s.Rrb()
}
