package main

import (
	"bufio"
	"fmt"
	"os"

	"push-swap/internal/pushswap"
)

func execInstruction(s *pushswap.Stacks, inst string) bool {
	switch inst {
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
	default:
		return false
	}
	return true
}

func main() {
	if len(os.Args) == 1 {
		return
	}

	nums, err := pushswap.ParseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	s := pushswap.NewStacks(nums)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if !execInstruction(s, line) {
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	if s.IsSorted() {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
