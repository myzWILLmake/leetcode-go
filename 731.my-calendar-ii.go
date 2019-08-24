package main

/*
 * @lc app=leetcode id=731 lang=golang
 *
 * [731] My Calendar II
 */
type Event struct {
	s int
	e int
}

type MyCalendarTwo struct {
	l1 []*Event
	l2 []*Event
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{[]*Event{}, []*Event{}}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	for _, c := range this.l2 {
		if start >= c.e || end <= c.s {
			continue
		}
		return false
	}

	for _, c := range this.l1 {
		if start >= c.e || end <= c.s {
			continue
		}
		sn := start
		if c.s > sn {
			sn = c.s
		}
		en := end
		if c.e < en {
			en = c.e
		}
		this.l2 = append(this.l2, &Event{sn, en})
	}
	this.l1 = append(this.l1, &Event{start, end})
	return true
}

// func main() {
// 	obj := Constructor()
// 	fmt.Println(obj.Book(28, 46))
// 	fmt.Println(obj.Book(9, 21))
// 	fmt.Println(obj.Book(21, 39))
// 	fmt.Println(obj.Book(37, 48))
// 	fmt.Println(obj.Book(38, 50))
// 	fmt.Println(obj.Book(22, 39))
// 	fmt.Println(obj.Book(45, 50))
// 	fmt.Println(obj.Book(1, 12))
// 	fmt.Println(obj.Book(40, 50))
// 	fmt.Println(obj.Book(31, 44))
// }

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
