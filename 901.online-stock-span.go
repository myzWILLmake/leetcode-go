package main

/*
 * @lc app=leetcode id=901 lang=golang
 *
 * [901] Online Stock Span
 */

type StockSpanner struct {
	lastMax []int
	price   []int
}

func Constructor() StockSpanner {
	s := StockSpanner{}
	s.lastMax = make([]int, 0)
	s.price = make([]int, 0)
	return s
}

func (s *StockSpanner) Next(price int) int {
	ans := 0
	idx := len(s.price) - 1
	s.lastMax = append(s.lastMax, -1)
	s.price = append(s.price, price)

	for true {
		if idx == -1 {
			ans = len(s.price)
			break
		}

		if s.price[idx] > price {
			ans = len(s.price) - 1 - idx
			s.lastMax[len(s.lastMax)-1] = idx
			break
		} else {
			idx = s.lastMax[idx]
		}
	}

	return ans
}

// func main() {
// 	obj := Constructor()
// 	fmt.Println(obj.Next(100))
// 	fmt.Println(obj.Next(80))
// 	fmt.Println(obj.Next(60))
// 	fmt.Println(obj.Next(70))
// 	fmt.Println(obj.Next(60))
// 	fmt.Println(obj.Next(75))
// 	fmt.Println(obj.Next(85))
// }

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
