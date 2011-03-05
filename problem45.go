package main

import (
	"fmt"
)

func problem45() {
	//triangle := func(n int64) int64 {
	//	return (n * (n + 1)) / 2
	//}
	pentagonal := func(n int64) int64 {
		return (n * (3*n - 1)) / 2
	}
	hexagonal := func(n int64) int64 {
		return (n * (2*n - 1))
	}
	
	//plookup := map[int64]int64{}
	//tlookup := map[int64]int64{}
	
	outer:
	for i := int64(1); true; i++ {
		h := hexagonal(i)
		
		for j := int64(1); true; j++ {
			p := pentagonal(j)
			
			if p > h {
				continue outer
			} else if p == h {
				break
			}
		}
		
		fmt.Println(h)
	}
}
