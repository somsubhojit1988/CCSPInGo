package fibonacci

// naive recursion
func fib1(n int) (int, int) {
	recCnt := 0
	var wrkr func(n int) int

	wrkr = func(n int) int {
		recCnt++
		if n < 2 {
			return n
		}
		return wrkr(n-1) + wrkr(n-2)
	}

	ret := wrkr(n)
	return ret, recCnt
}

// recursion with memoization
func fib2(n int) (int, int) {
	recCnt := 0
	mem := map[int]int{0: 0, 1: 1}
	var wrkr func(n int) int

	wrkr = func(n int) int {
		recCnt++
		if v, ok := mem[n]; ok {
			return v
		}
		mem[n] = wrkr(n-1) + wrkr(n-2)
		return mem[n]
	}
	ret := wrkr(n)
	return ret, recCnt
}

// iterative solution
func fib3(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

// fibonacci generator, buffered channel
func fib4(n int, res chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		res <- x
		x, y = y, x+y
	}
	close(res)
}
