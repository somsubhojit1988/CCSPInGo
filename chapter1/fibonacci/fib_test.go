package fibonacci

import "testing"

var tt = []int{1, 2, 5, 10, 15, 25}

func TestFibNaiveRecursion(t *testing.T) {
	for _, tc := range tt {
		ret, n := fib1(tc)
		t.Logf("fib1(%d)= %d, # recursive calls= %d\n", tc, ret, n)
	}
}

func TestFibMemoizedRecusion(t *testing.T) {
	for _, tc := range tt {
		ret, n := fib2(tc)
		t.Logf("fib2(%d)= %d, # recursive calls= %d\n", tc, ret, n)
	}
}

func TestFibIterative(t *testing.T) {
	for _, tc := range tt {
		ret := fib3(tc)
		t.Logf("fib3(%d)= %d\n", tc, ret)
	}
}

func TestFibGenerator(t *testing.T) {

	for _, tc := range tt {
		c := make(chan int,
			func(x, y int) int {
				if x > y {
					return x
				}
				return y
			}(tc, 1))
		fib4(cap(c), c)
		i := 1
		for x := range c {
			t.Logf("fibonacci(%d) = %d\n", i, x)
			i++
		}
		t.Logf("\n")
	}
}
