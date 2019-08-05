package main

/*
 * @lc app=leetcode id=416 lang=golang
 *
 * [416] Partition Equal Subset Sum
 */
func canPartition(nums []int) bool {
	n := len(nums)
	if n == 1 {
		if nums[0] == 0 {
			return true
		} else {
			return false
		}
	}

	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}

	val := sum / 2
	sumMap := make(map[int]bool)
	for i := 0; i < n; i++ {
		if nums[i] == val {
			return true
		}

		keys := make([]int, len(sumMap))
		for k := range sumMap {
			keys = append(keys, k)
		}

		for _, k := range keys {
			v := k + nums[i]
			if v < val {
				if !sumMap[v] {
					sumMap[v] = true
				}
			} else if v == val {
				return true
			}
		}
		if nums[i] <= val && !sumMap[nums[i]] {
			sumMap[nums[i]] = true
		}
	}

	return false
}

// func main() {
// 	nums := []int{2, 2, 3, 5}
// 	fmt.Println(canPartition(nums))
// }
