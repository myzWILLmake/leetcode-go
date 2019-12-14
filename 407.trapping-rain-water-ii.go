package main

import (
	"container/heap"
	"fmt"
)

/*
 * @lc app=leetcode id=407 lang=golang
 *
 * [407] Trapping Rain Water II
 */

// @lc code=start

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

type Point struct {
	x int
	y int
	h int
}

type pHeap []*Point

func (h pHeap) Len() int           { return len(h) }
func (h pHeap) Less(i, j int) bool { return h[i].h < h[j].h }
func (h pHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *pHeap) Push(x interface{}) {
	p := x.(*Point)
	*h = append(*h, p)
}

func (h *pHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func trapRainWater(heightMap [][]int) int {
	if len(heightMap) == 0 {
		return 0
	}

	n := len(heightMap)
	m := len(heightMap[0])
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	h := &pHeap{}
	heap.Init(h)

	for i := 0; i < n; i++ {
		p1 := &Point{i, 0, heightMap[i][0]}
		p2 := &Point{i, m - 1, heightMap[i][m-1]}
		heap.Push(h, p1)
		heap.Push(h, p2)
		visited[i][0] = true
		visited[i][m-1] = true
	}

	for i := 0; i < m; i++ {
		p1 := &Point{0, i, heightMap[0][i]}
		p2 := &Point{n - 1, i, heightMap[n-1][i]}
		heap.Push(h, p1)
		heap.Push(h, p2)
		visited[0][i] = true
		visited[n-1][i] = true
	}

	ans := 0
	for h.Len() > 0 {
		p := heap.Pop(h).(*Point)
		adjacent := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for i := 0; i < 4; i++ {
			x := adjacent[i][0] + p.x
			y := adjacent[i][1] + p.y
			if x >= 0 && x < n && y >= 0 && y < m {
				if !visited[x][y] {
					visited[x][y] = true
					ans += max(0, p.h-heightMap[x][y])
					p := &Point{x, y, max(p.h, heightMap[x][y])}
					heap.Push(h, p)
				}
			}
		}
	}
	return ans
}

// @lc code=end

func main() {
	heightMap := [][]int{
		{1, 4, 3, 1, 3, 2},
		{3, 2, 1, 3, 2, 4},
		{2, 3, 3, 2, 3, 1},
	}
	fmt.Println(trapRainWater(heightMap))

}
