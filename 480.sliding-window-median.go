package main

import (
	"sort"
)

type Node struct {
	val    int
	heapId int
}

type Heap struct {
	data        []*Node
	length      int
	nodeCompare func(x, y *Node) bool
}

func (h *Heap) compare(idx, idy int) bool {
	return h.nodeCompare(h.data[idx], h.data[idy])
}

func (h *Heap) pop(idx int) {
	for idx > 0 {
		fIdx := (idx - 1) / 2
		if !h.compare(fIdx, idx) {
			h.swap(fIdx, idx)
			idx = fIdx
		} else {
			break
		}
	}
}

func (h *Heap) sink(idx int) {
	for idx < h.length {
		lIdx := idx*2 + 1
		rIdx := idx*2 + 2
		if rIdx < h.length && h.compare(rIdx, lIdx) {
			if h.compare(rIdx, idx) {
				h.swap(rIdx, idx)
				idx = rIdx
			} else {
				break
			}
		} else if lIdx < h.length {
			if h.compare(lIdx, idx) {
				h.swap(lIdx, idx)
				idx = lIdx
			} else {
				break
			}
		} else {
			break
		}
	}
}

func (h *Heap) swap(idx, idy int) {
	h.data[idx], h.data[idy] = h.data[idy], h.data[idx]
	h.data[idx].heapId = idx
	h.data[idy].heapId = idy
}

func (h *Heap) add(node *Node) {
	h.length++
	idx := h.length - 1
	node.heapId = idx
	h.data[idx] = node
	h.pop(idx)
}

func (h *Heap) remove(node *Node) {
	idx := node.heapId
	h.swap(idx, h.length-1)
	node.heapId = -1

	h.data[h.length-1] = nil
	h.length--

	if idx != 0 && idx < h.length && h.compare(idx, (idx-1)/2) {
		h.pop(idx)
	} else {
		h.sink(idx)
	}
}

func (h *Heap) top() *Node {
	return h.data[0]
}

/*
 * @lc app=leetcode id=480 lang=golang
 *
 * [480] Sliding Window Median
 */
func medianSlidingWindow(nums []int, k int) []float64 {
	n := len(nums)
	minHeap := Heap{nil, 0, func(l, r *Node) bool {
		return l.val <= r.val
	}}
	maxHeap := Heap{nil, 0, func(l, r *Node) bool {
		return l.val >= r.val
	}}
	minHeap.data = make([]*Node, k-k/2)
	maxHeap.data = make([]*Node, k/2)

	nodes := make([]*Node, n)
	for i := 0; i < n; i++ {
		n := Node{nums[i], -1}
		nodes[i] = &n
	}

	initNodes := make([]*Node, k)
	copy(initNodes, nodes[0:k])
	sort.SliceStable(initNodes, func(i, j int) bool {
		return initNodes[i].val < initNodes[j].val
	})

	for i := 0; i < k/2; i++ {
		maxHeap.add(initNodes[i])
	}
	for i := k / 2; i < k; i++ {
		minHeap.add(initNodes[i])
	}

	ans := make([]float64, n-k+1)
	if k%2 == 0 {
		ans[0] = float64(minHeap.top().val+maxHeap.top().val) / 2
	} else {
		ans[0] = float64(minHeap.top().val)
	}

	for i := k; i < n; i++ {
		removeNode := nodes[i-k]
		addNode := nodes[i]
		if removeNode != minHeap.data[removeNode.heapId] {
			maxHeap.remove(removeNode)
			if addNode.val > minHeap.top().val {
				tmpNode := minHeap.top()
				minHeap.remove(tmpNode)
				maxHeap.add(tmpNode)
				minHeap.add(addNode)
			} else {
				maxHeap.add(addNode)
			}
		} else {
			minHeap.remove(removeNode)
			if maxHeap.length > 0 && addNode.val < maxHeap.top().val {
				tmpNode := maxHeap.top()
				maxHeap.remove(tmpNode)
				minHeap.add(tmpNode)
				maxHeap.add(addNode)
			} else {
				minHeap.add(addNode)
			}
		}

		if k%2 == 0 {
			ans[i-k+1] = float64(minHeap.top().val+maxHeap.top().val) / 2
		} else {
			ans[i-k+1] = float64(minHeap.top().val)
		}
	}
	return ans
}

// func main() {
// 	nums := []int{5, 5, 8, 1, 4, 7, 1, 3, 8, 4}
// 	k := 8
// 	fmt.Println(medianSlidingWindow(nums, k))
// }
