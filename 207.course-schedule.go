package main

import "fmt"

/*
 * @lc app=leetcode id=207 lang=golang
 *
 * [207] Course Schedule
 */

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	nodes := make([]bool, numCourses)
	indgree := make([]int, numCourses)
	nodeMap := make([]map[int]bool, numCourses)

	for i := 0; i < numCourses; i++ {
		nodes[i] = true
		nodeMap[i] = map[int]bool{}
		indgree[i] = 0
	}

	for _, pair := range prerequisites {
		if !nodeMap[pair[1]][pair[0]] {
			nodeMap[pair[1]][pair[0]] = true
			indgree[pair[0]]++
		}
	}

	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if indgree[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		top := queue[0]
		nodes[top] = false
		queue = queue[1:]
		edges := nodeMap[top]
		for n := range edges {
			indgree[n]--
			if indgree[n] == 0 {
				queue = append(queue, n)
			}
		}
	}

	for i := 0; i < numCourses; i++ {
		if nodes[i] {
			return false
		}
	}

	return true
}

// @lc code=end

func main() {
	numCourses := 2
	prerequisites := [][]int{{1, 0}}

	fmt.Print(canFinish(numCourses, prerequisites))
}
