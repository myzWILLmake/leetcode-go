package main

import "fmt"

/*
 * @lc app=leetcode id=399 lang=golang
 *
 * [399] Evaluate Division
 */

// @lc code=start
func doquery(cur, end string, flags map[string]bool, edges map[string]map[string]float64) float64 {

	e, ok := edges[cur]
	if !ok {
		return -1
	}
	if cur == end {
		return 1
	}
	for next, value := range e {
		if flags[next] {
			continue
		}
		if next == end {
			return value
		} else {
			flags[next] = true
			res := doquery(next, end, flags, edges)
			if res != -1 {
				return res * value
			}
			flags[next] = false
		}

	}
	return -1
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	edges := make(map[string]map[string]float64)
	for i := 0; i < len(equations); i++ {
		e := equations[i]
		x := e[0]
		y := e[1]
		if _, ok := edges[x]; !ok {
			edges[x] = make(map[string]float64)
		}
		if _, ok := edges[y]; !ok {
			edges[y] = make(map[string]float64)
		}
		edges[x][y] = values[i]
		edges[y][x] = 1 / values[i]
	}

	res := []float64{}
	for i := 0; i < len(queries); i++ {
		flags := make(map[string]bool)
		start := queries[i][0]
		end := queries[i][1]
		flags[start] = true
		res = append(res, doquery(start, end, flags, edges))
	}

	return res
}

// @lc code=end

func main() {
	equations := [][]string{{"a", "b"}, {"b", "c"}}
	values := []float64{2.0, 3.0}
	queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
	fmt.Println(calcEquation(equations, values, queries))
}
