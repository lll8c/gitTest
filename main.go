package main

import "fmt"

var (
	n   int
	st  [10]bool
	arr [10]int
)

func dfs(k int) {
	if k == n {
		for i := 0; i < n; i++ {
			fmt.Println(arr)
		}
	}

	for i := 1; i <= n; i++ {
		if st[i] {
			continue
		}
		st[i] = true
		arr[k] = i
		dfs(k + 1)
		st[i] = false
	}
}

func main() {
	n = 5

	dfs(0)
}
