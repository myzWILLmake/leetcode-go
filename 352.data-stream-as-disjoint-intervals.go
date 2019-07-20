package main

/*
 * @lc app=leetcode id=352 lang=golang
 *
 * [352] Data Stream as Disjoint Intervals
 */
type Node struct {
	start int
	end   int
	next  *Node
	prev  *Node
}

type SummaryRanges struct {
	head *Node
	tail *Node
}

/** Initialize your data structure here. */
func Constructor() SummaryRanges {
	s := SummaryRanges{}
	initNode := Node{-2, -2, nil, nil}
	s.head = &initNode
	s.tail = &initNode
	return s
}

func (this *SummaryRanges) doCombine(left *Node, right *Node) bool {
	if left == nil || right == nil {
		return false
	}

	if left.end+1 == right.start {
		left.end = right.end
		left.next = right.next
		if right.next != nil {
			right.next.prev = left
		}
		if right == this.tail {
			this.tail = left
		}
		return true
	}
	return false
}

func (this *SummaryRanges) AddNum(val int) {
	node := this.head

	for node != nil {
		if node.start > val {
			newNode := Node{val, val, nil, nil}
			newNode.prev = node.prev
			newNode.next = node
			node.prev.next = &newNode
			node.prev = &newNode

			this.doCombine(&newNode, node)
			this.doCombine(newNode.prev, &newNode)
			break
		} else if node.start <= val && node.end >= val {
			break
		} else if node.end < val {
			node = node.next
		}
	}

	if node == nil {
		newNode := Node{val, val, nil, nil}
		newNode.prev = this.tail
		this.tail.next = &newNode
		this.tail = &newNode

		this.doCombine(newNode.prev, &newNode)
	}
}

func (this *SummaryRanges) GetIntervals() [][]int {
	node := this.head.next

	res := [][]int{}
	for node != nil {
		item := []int{node.start, node.end}
		res = append(res, item)
		node = node.next
	}
	return res
}

// func main() {
// 	obj := Constructor()
// 	obj.AddNum(1)
// 	obj.AddNum(3)
// 	obj.AddNum(7)
// 	obj.AddNum(2)
// 	obj.AddNum(6)
// 	fmt.Println(obj.GetIntervals())
// }
