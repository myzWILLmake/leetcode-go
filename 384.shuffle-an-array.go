package main

import (
	"math/rand"
	"time"
)

/*
 * @lc app=leetcode id=384 lang=golang
 *
 * [384] Shuffle an Array
 */
type Solution struct {
	data []int
	r    *rand.Rand
}

func Constructor(nums []int) Solution {
	s := Solution{}
	s.data = nums
	source := rand.NewSource(time.Now().UnixNano())
	s.r = rand.New(source)
	return s
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.data
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	n := len(this.data)
	f := make([]int, n)
	copy(f, this.data)
	for i := n - 1; i > 0; i-- {
		random := this.r.Intn(i + 1)
		f[i], f[random] = f[random], f[i]
	}
	return f
}

// func main() {
// 	nums := []int{1, 2, 3, 4, 5}
// 	obj := Constructor(nums)
// 	fmt.Println(obj.Shuffle())
// 	fmt.Println(obj.Shuffle())
// 	fmt.Println(obj.Shuffle())
// 	fmt.Println(obj.Shuffle())
// 	fmt.Println(obj.Shuffle())
// }

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
