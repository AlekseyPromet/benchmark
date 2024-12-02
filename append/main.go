package main

import (
	"flag"
	"strconv"
)

func main() {

	var n uint64
	var err error

	if n, err = strconv.ParseUint(flag.Arg(1), 10, 64); err != nil {
		n = 1_000_000
	}

	Run(n)
}

func Run(n uint64) {

	s := make([]uint64, 0)

	for i := uint64(0); i < n; i++ {
		s = append(s, i)
	}

}
