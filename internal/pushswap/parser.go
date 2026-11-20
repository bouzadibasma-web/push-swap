package pushswap

import (
	"errors"
	"strconv"
	"strings"
)

// ParseArgs convertit les arguments en une slice d'entiers uniques.
// Retourne une erreur s'il y a un élément non entier ou un doublon.
func ParseArgs(args []string) ([]int, error) {
	var result []int
	seen := make(map[int]bool)

	for _, arg := range args {
		tokens := strings.Fields(arg)
		for _, tok := range tokens {
			val, err := strconv.Atoi(tok)
			if err != nil {
				return nil, errors.New("invalid integer")
			}
			if seen[val] {
				return nil, errors.New("duplicate integer")
			}
			seen[val] = true
			result = append(result, val)
		}
	}

	return result, nil
}
