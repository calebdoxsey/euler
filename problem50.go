package main

import (
	"fmt"
)

/*
The prime 41, can be written as the sum of six consecutive primes:

41 = 2 + 3 + 5 + 7 + 11 + 13
This is the longest sum of consecutive primes that adds to a prime below one-hundred.

The longest sum of consecutive primes below one-thousand that adds to a prime, contains 21 terms, and is equal to 953.

Which prime, below one-million, can be written as the sum of the most consecutive primes?
*/

func problem50() {
	below := int64(1000000)
	// Calculate Primes
	fmt.Println("Calcuting primes")
	lookup := make([]int64, 1)
	lookup[0] = 2
	length := 1
	for i := int64(3); i < below; i++ {
		prime := true
		for k := 0; true; k++ {
			v := lookup[k]
			if i % lookup[k] == 0 {
				prime = false
				break
			}
			if v*v > i {
				break
			}
		}
		if prime {			
			if length >= len(lookup) {
				old := lookup
				lookup = make([]int64, length + 100)
				copy(lookup, old)
			}
			lookup[length] = i
			length++
		}
	}
	
	sum := func(from, to int) int64 {
		v := int64(0)
		for i := from; i < to; i++ {
			v += lookup[i]
		}
		return v
	}	
	
	longest := 1
	prime := int64(0)
	for i := length-1; i > 0; i-- {		
		outer:
		for j := 0; j <= i-longest; j++ {
			for k := j+longest; k <= i; k++ {
				s := sum(j,k)
				if s == lookup[i] {
					fmt.Println(lookup[i], k-j)
					longest = k-j
					prime = lookup[i]					
				} else if s > lookup[i] {
					continue outer
				}
			}
		}
	}
	
	fmt.Println("Answer: ", prime, " #: ", longest)
}


