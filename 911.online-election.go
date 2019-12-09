package main

import (
	"container/heap"
	"fmt"
)

/*
 * @lc app=leetcode id=911 lang=golang
 *
 * [911] Online Election
 */

// @lc code=start
type Candidate struct {
	no    int
	times int
	idx   int
	last  int
}

type CandidateHeap []*Candidate

func (h CandidateHeap) Len() int { return len(h) }
func (h CandidateHeap) Less(i, j int) bool {
	if h[i].times == h[j].times {
		return h[i].last > h[j].last
	}
	return h[i].times > h[j].times
}
func (h CandidateHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].idx = i
	h[j].idx = j
}

func (h *CandidateHeap) Push(x interface{}) {
	c := x.(*Candidate)
	c.idx = h.Len()
	*h = append(*h, c)
}

func (h *CandidateHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	x.idx = -1
	*h = old[0 : n-1]
	return x
}

type TopVotedCandidate struct {
	candidateMap  map[int]*Candidate
	candidateHeap CandidateHeap
	tops          []int
	times         []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	t := TopVotedCandidate{}
	t.candidateMap = map[int]*Candidate{}
	t.candidateHeap = CandidateHeap{}
	heap.Init(&t.candidateHeap)
	t.tops = make([]int, len(times))
	t.times = times
	for i := 0; i < len(persons); i++ {
		c, ok := t.candidateMap[persons[i]]
		if !ok {
			c = &Candidate{persons[i], 1, -1, times[i]}
			t.candidateMap[persons[i]] = c
			heap.Push(&t.candidateHeap, c)
		} else {
			c.times++
			c.last = times[i]
			heap.Fix(&t.candidateHeap, c.idx)
		}
		top := t.candidateHeap[0]
		t.tops[i] = top.no
	}
	return t
}

func (this *TopVotedCandidate) Q(t int) int {
	pos := 0
	for ; pos < len(this.times); pos++ {
		if t < this.times[pos] {
			break
		}
	}
	if pos == 0 {
		return 0
	}
	top := this.tops[pos-1]
	return top
}

// [[0,1,1,  0, 0, 1, 0],
//  [0,5,10,15,20,25,30]]
/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */
// @lc code=end

func main() {
	persons := []int{0, 1, 1, 0, 0, 1, 0}
	times := []int{0, 5, 10, 15, 20, 25, 30}
	t := Constructor(persons, times)
	fmt.Println(t.Q(3))
	fmt.Println(t.Q(12))
	fmt.Println(t.Q(25))
	fmt.Println(t.Q(15))
	fmt.Println(t.Q(24))
	fmt.Println(t.Q(8))
}
