package main

func twoSum(nums []int, target int) []int {
	var res [2]int
	if len(nums) < 2 {
		return nil
	}
	for i := 0; i < len(nums)-1; i++ {
		num1 := nums[i]
		for j := i + 1; j < len(nums); j++ {
			num2 := nums[j]
			if num1+num2 == target {
				res[0] = i
				res[1] = j
				return res[:]
			}
		}
	}

	return nil
}

// func main() {
// 	nums := []int{2, 7, 11, 15}
// 	target := 9
// 	res := twoSum(nums, target)
// 	fmt.Println(res)
// }
