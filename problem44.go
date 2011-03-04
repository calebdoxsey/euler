package main

import (
	"fmt"
	//"math"
)

func problem44() {
	pentagonal := func(n int64) int64 {
		return (n * (3*n - 1)) / 2
	}
	
	lookup := map[int64]int64{}
	
	min := int64(99999999)
	for i := int64(2); i < 100000; i++ {
		p1 := pentagonal(i)
		lookup[p1] = i
		
		inner:
		for j := i-1; j > 0; j-- {
			p2 := pentagonal(j)
			
			diff := p1 - p2
			
			if _, e := lookup[diff]; !e {
				continue
			}
			
			//fmt.Println("-", i, j, p1, p2, diff)
			
			sum := p2 + p1
			
			if _, e := lookup[sum]; !e {			
				for k := j+1; true; k++ {
					p3 := pentagonal(k)
					lookup[p3] = k
				
					if p3 > sum {
						continue inner
					} else if p3 == sum {
						break
					}
				}
			}
			
			if diff < min {
				fmt.Println("+", i, j, p1, p2, diff, sum)
				min = diff
			}
			
			//return
		}
	}
}
