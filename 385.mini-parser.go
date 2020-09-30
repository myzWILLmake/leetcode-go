package main

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode id=385 lang=golang
 *
 * [385] Mini Parser
 */
type NestedInteger struct {
	val  int
	list []*NestedInteger
}

func (n *NestedInteger) SetInteger(value int) {
	n.val = value
}
func (n *NestedInteger) Add(elem NestedInteger) {
	n.list = append(n.list, &elem)
}

// @lc code=start
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
func deserialize(s string) *NestedInteger {
	n := len(s)
	listStack := make([]int, 0)
	listPos := make(map[int]int)
	for i := 0; i < n; i++ {
		if s[i] == '[' {
			listStack = append(listStack, i)
		} else if s[i] == ']' {
			l := listStack[len(listStack)-1]
			listPos[l] = i
			listStack = listStack[:len(listStack)-1]
		}
	}
	// return doDeserialize(s, listPos, 0)
	res := &NestedInteger{}
	if s[0] == '[' {
		num := 0
		hasNum := false
		isNegtive := false
		for i := 1; i < n; i++ {
			if s[i] == '[' {
				subs := s[i : listPos[i]+1]
				subNest := deserialize(subs)
				res.Add(*subNest)
				i = listPos[i]
			} else if s[i] == ']' || s[i] == ',' {
				if hasNum {
					subNest := &NestedInteger{}
					subNest.SetInteger(num)
					res.Add(*subNest)
					num = 0
					hasNum = false
					isNegtive = false
				}
			} else if s[i] == '-' {
				isNegtive = true
			} else {
				hasNum = true
				digit, _ := strconv.Atoi(s[i : i+1])
				if isNegtive {
					num = num*10 - digit
				} else {
					num = num*10 + digit
				}
			}
		}
	} else {
		val, _ := strconv.Atoi(s)
		res.SetInteger(val)
	}

	return res
}

// @lc code=end

func main() {
	s := "[-1]"
	fmt.Println(deserialize(s))
}
