package main

import (
	"fmt"
	"math"
)

func problem19() {
	daysInMonth := []func(int)int{
		func(year int)int{ return 31 },
		func(year int)int{ 
			if year % 4 == 0 {
				if year % 100 == 0 && year % 400 != 0 {
					return 28
				}
				return 29
			}
			return 28 
		},
		func(year int)int{ return 31 },
		func(year int)int{ return 30 },
		func(year int)int{ return 31 },
		func(year int)int{ return 30 },
		func(year int)int{ return 31 },
		func(year int)int{ return 31 },
		func(year int)int{ return 30 },
		func(year int)int{ return 31 },
		func(year int)int{ return 30 },
		func(year int)int{ return 31 },		
	}
	
	// Jan 1, 1900 was a monday
	dayOfWeek := 1
	
	cnt := 0
	
	// Go to 2001
	for y := 1900; y < 2001; y++ {
		for m, f := range daysInMonth {
			if dayOfWeek == 0 && y > 1900 {
				cnt++
			}
			days := f(y)
			fmt.Println(m, days, dayOfWeek)
			dayOfWeek += days
			dayOfWeek = dayOfWeek % 7
		}
	}
	
	fmt.Println("")
	fmt.Println("Total: ", cnt)
}

func problem20() {
	quad := func(a, b int) func(int)int {
		return func(n int) int {
			return n*n + a*n + b
		}
	}
	
	highv := 0
	higha := -1000
	highb := -1000
	
	for a := -1000; a < 1000; a++ {
		for b := -1000; b < 1000; b++ {
			f := quad(a, b)
			for i := 0; true; i++ {
				next := f(i)
				if !isPrime(int64(next)) {
					if i > highv {
						highv = i
						higha = a
						highb = b
						fmt.Println(highv, higha, highb)
					}
					break
				}
			}
		}
	}	
	
	fmt.Println(highv, higha, highb)
}

func problem32() {
	MUL := -2
	EQU := -1
	
	items := []int{MUL,EQU,1,2,3,4,5,6,7,8,9}
	n := len(items)

	swap := func(i, j int) {
		t := items[i]
		items[i] = items[j]
		items[j] = t
	}

	permute := func() bool {
		i := n-1
		for items[i-1] >= items[i] {
			i--
			if i < 1 {
				return false
			}
		}
		j := n
		for items[j-1] <= items[i-1] {
			j--
		}
		swap(i-1, j-1)
		i++
		j = n
		
		for i < j {
			swap(i-1, j-1)
			i++
			j--
		}
		return true
	}
	
	valid := func() (bool, int) {
		m := 0
		for ; items[m] != MUL; m++ {}
		e := 0
		for ; items[e] != EQU; e++ {}
		
		if e < m || m == 0 || e == n-1 || e-m < 2 {
			return false, 0
		}
		
		left := 0
		for i := 0; i < m; i++ {
			left += items[i] * int(math.Pow(10, float64(m-i-1)))
		}
		
		right := 0
		for i := m+1; i < e; i++ {
			right += items[i] * int(math.Pow(10, float64(e-i-1)))
		}
		
		prod := 0
		for i := e+1; i < n; i++ {
			prod += items[i] * int(math.Pow(10, float64(n-i-1)))
		}
		
		if left * right == prod {		
			fmt.Println(left, "*", right, "=", prod)
			return true, prod
		}
		
		return false, 0
	}
	
	results := make(map[int]bool)	
	for {
		if v, prod := valid(); v {
			results[prod] = true
		}
	
		if !permute() {
			break
		}
	}
	fmt.Println("Results:")
	sum := 0
	for k, _ := range results {
		fmt.Println(k)
		sum += k
	}
	fmt.Println("Sum:")
	fmt.Println(sum)
}

func problem38() {
	scale := func(i int64) int64 {
		for j := int64(10); true; j *= int64(10) { 
			if i < j {			
				return j
			}
		}
		return 0
	}

	max := int64(987654321)
	for n := int64(2); n < max; n++ {
		conc := int64(0)
		for i := int64(1); true; i++ {
			product := n * i
			conc = conc * scale(product) + product
			if conc > max {
				break
			}
			if i > 1 && pandigital(conc, 9) {
				fmt.Println(conc)
			}
		}
	}
}

func problem39() {
	maxv := 0
	maxp := 0
	for p := 1; p <= 1000; p++ {
		cnt := 0
		for a := 1; a <= p; a++ {
			for b := 1; a+b <= p; b++ {
				for c := 1; a+b+c <= p; c++ {
					if a+b+c == p && a*a + b*b == c*c {
						cnt++
						fmt.Println(p, ":", a, b, c)
					}
				}
			}
		}
		if cnt > maxv {
			maxv = cnt
			maxp = p
		}
	}
	fmt.Println("Solution:")
	fmt.Println("p: ", maxp, "v: ", maxv)
}

func problem41() {
	scale := func(i int64) int {
		cnt := 1
		for j := int64(10); true; j *= int64(10) { 
			if i < j {			
				return cnt
			}
			cnt++
		}
		return 0
	}
	//max := int64(987654321)
	max := int64(987654321)
	for i := max; i > 1; i-- {
		if isPrime(i) {
			fmt.Println("Prime: ", i)
			if pandigital(i, scale(i)) {
				fmt.Println(i)
				break
			}
		}
	}
}

func main() {
	problem44()
}
