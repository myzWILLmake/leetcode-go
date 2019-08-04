package main

/*
 * @lc app=leetcode id=295 lang=golang
 *
 * [295] Find Median from Data Stream
 */
type Node struct {
	val  int
	next *Node
}

type MedianFinder struct {
	data *Node
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	m := MedianFinder{}
	return m
}

func (this *MedianFinder) AddNum(num int) {
	newNode := &Node{num, nil}
	if this.data == nil {
		this.data = newNode
		return
	}

	if newNode.val <= this.data.val {
		newNode.next = this.data
		this.data = newNode
		return
	}

	it := this.data
	for it.next != nil && it.next.val < newNode.val {
		it = it.next
	}

	newNode.next = it.next
	it.next = newNode
}

func (this *MedianFinder) FindMedian() float64 {
	if this.data == nil {
		return 0
	}

	it := this.data
	it2 := this.data

	for it != nil {
		if it2.next != nil {
			if it2.next.next != nil {
				it = it.next
				it2 = it2.next.next
			} else {
				return float64(it.val+it.next.val) / 2
			}
		} else {
			return float64(it.val)
		}
	}
	return float64(it.val)
}

// func main() {
// 	obj := Constructor()
// 	obj.AddNum(1)
// 	obj.AddNum(2)
// 	fmt.Println(obj.FindMedian())
// 	obj.AddNum(3)
// 	fmt.Println(obj.FindMedian())
// }

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
