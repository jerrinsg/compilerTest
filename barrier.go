package main

const (
	N = 65536
)

var (
	a [N]int
)

//go:noinline
func compilerBarrier()

var fenceFunc func()

//go:noinline
func compute() {
	for i := 0; i < N; i++ {
		a[i] += 1
		fenceFunc()
		a[i] += 1
	}
}

func main() {
	fenceFunc = compilerBarrier
	sum := 0
	compute()
	for i := 0; i < N; i++ {
		sum += a[i]
	}
	println(sum)
}
