package main

import "fmt"

/*
 * @lc app=leetcode id=641 lang=golang
 *
 * [641] Design Circular Deque
 */

// @lc code=start
type Node struct {
	val  int
	prev *Node
	next *Node
}

type MyCircularDeque struct {
	size   int
	length int
	head   *Node
	tail   *Node
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	hNode := &Node{}
	tNode := &Node{}
	hNode.next = tNode
	tNode.prev = hNode
	deque := MyCircularDeque{k, 0, hNode, tNode}
	return deque
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.length == this.size {
		return false
	}

	node := &Node{value, nil, nil}
	this.head.next.prev = node
	node.next = this.head.next
	node.prev = this.head
	this.head.next = node
	this.length++
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.length == this.size {
		return false
	}

	node := &Node{value, nil, nil}
	this.tail.prev.next = node
	node.prev = this.tail.prev
	node.next = this.tail
	this.tail.prev = node
	this.length++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.length == 0 {
		return false
	}

	node := this.head.next
	this.head.next = node.next
	node.next.prev = this.head
	node.prev = nil
	node.next = nil
	this.length--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.length == 0 {
		return false
	}

	node := this.tail.prev
	this.tail.prev = node.prev
	node.prev.next = this.tail
	node.prev = nil
	node.next = nil
	this.length--
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.length == 0 {
		return -1
	}
	return this.head.next.val
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.length == 0 {
		return -1
	}
	return this.tail.prev.val
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.length == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.length == this.size
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
// @lc code=end

func main() {
	c := Constructor(3)
	c.InsertLast(1)
	c.InsertLast(2)
	c.InsertFront(3)
	c.InsertFront(4)

	fmt.Println(c.GetRear())
	fmt.Println(c.IsFull())
	c.DeleteLast()
	c.InsertFront(4)
	fmt.Println(c.GetFront())
}
