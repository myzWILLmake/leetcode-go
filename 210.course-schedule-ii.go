package main

import "fmt"

/*
 * @lc app=leetcode id=210 lang=golang
 *
 * [210] Course Schedule II
 */

// @lc code=start
func findOrder(numCourses int, prerequisites [][]int) []int {
	res := []int{}
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
			res = append(res, i)
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
				res = append(res, n)
				queue = append(queue, n)
			}
		}
	}

	for i := 0; i < numCourses; i++ {
		if nodes[i] {
			return []int{}
		}
	}

	return res
}

// @lc code=end

func main() {
	numCoures := 4
	prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	fmt.Println(findOrder(numCoures, prerequisites))
}
