package main

import (
	"fmt"
	"os"

	"push-swap/internal/pushswap"
)

func main() {
	if len(os.Args) == 1 {
		return
	}

	nums, err := pushswap.ParseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	ops := pushswap.Solve(nums)
	for _, op := range ops {
		fmt.Println(op)
	}
}
