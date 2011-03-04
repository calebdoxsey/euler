package main

import (
	"fmt"
	//"math"
)

func problem43() {
	items := []int{0,1,2,3,4,5,6,7,8,9}
	toNumber := func(start int, length int) uint64 {
		sum := uint64(0)
		scale := uint64(1)
		for i := length-1; i >= 0; i-- {
			sum += uint64(items[start+i]) * scale
			scale *= 10
		}
		return sum
	}

	check := []uint64{
		2,
		3,
		5,
		7,
		11,
		13,
		17,
	}
	c := permute(items)
	
	sofar := uint64(0)
	for <-c {
		valid := true
		for k, v := range check {
			n := toNumber(k+1, 3)
			if n % v != 0 {
				valid = false
				break
			}
		}
		if valid {
			sofar += toNumber(0, len(items))
			fmt.Println(items, sofar)
		}
	}
	fmt.Println(sofar)
}
