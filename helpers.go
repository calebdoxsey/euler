package main

import (
	//"fmt"
)

var primelookup []int64

func init() {
	primelookup = make([]int64, 1000)
	primelookup[0] = 2
	for i := 1; i < len(primelookup); i++ {
		for j := primelookup[i-1]+1; true; j++ {
			prime := true
			for k := i-1; k >= 0; k-- {
				if j % primelookup[k] == 0 {
					prime = false
					break
				}
			}
			if prime {
				primelookup[i] = j
				break
			}
		}
	}
}

func isPrime(n int64) bool {
	if n < int64(2) {
		return false
	}
	for i := 0; i < len(primelookup); i++ {
		if primelookup[i] >= n {
			return true
		}
		if n % primelookup[i] == 0 {
			return false
		}
	}
	last := primelookup[len(primelookup)-1]
	for i := last+1; i*i <= n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func pandigital(n int64, sz int) bool {
	has := make(map[int64]bool, sz)
	for i := 1; i <= sz; i++ {
		has[int64(i)] = false
	}
	for n > 0 {
		d := n % 10
		if _, e := has[d]; !e {
			return false
		}
		has[d] = true
		n = n / 10
	}
	for _, e := range has {
		if !e {
			return false
		}
	}
	return true
}
