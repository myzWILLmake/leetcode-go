package main

// f[i] = f[i-1] + f[i-2] + g[i-1]
// g[i] = g[i-2] * 2 + g[i-1]

func numTilings(N int) int {
	if N == 1 {
		return 1
	} else if N == 2 {
		return 2
	}

	f := make([]int, N+1)
	g := make([]int, N+1)

	f[1] = 1
	f[2] = 2
	g[1] = 0
	g[2] = 2

	for i := 3; i <= N; i++ {
		f[i] = (f[i-1] + f[i-2] + g[i-1]) % 1000000007
		g[i] = (f[i-2]*2 + g[i-1]) % 1000000007
	}

	return f[N]
}

// func main() {
// 	fmt.Println(numTilings(300))
// }
