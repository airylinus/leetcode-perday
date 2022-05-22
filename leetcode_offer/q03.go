package leetcode_offer

//在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。
//数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。
//请找出数组中任意一个重复的数字。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// SimplyFindRepeatNumber 找重复数字
func SimplyFindRepeatNumber(nums []int) int {
	uDict := map[int]struct{}{}
	for _, n := range nums {
		if _, has := uDict[n]; has {
			return n
		}
		uDict[n] = struct{}{}
	}
	return -1
}

// FindRepeatNumber ...
func FindRepeatNumber(nums []int) int {
	var i = 0
	for i < len(nums) {
		// index equal to value, continue next
		if nums[i] == i {
			i += 1
			continue
		}
		// not equal: nums[k]/v current value should move to nums[v],
		if nums[nums[i]] == nums[i] {
			return nums[i]
		}
		x := nums[nums[i]]
		nums[nums[i]] = nums[i]
		nums[i] = x
	}
	return -1
}

func FindRepeatNumberErrorVersion(nums []int) int {
	for i, _ := range nums {
		// index equal to value, continue next
		if nums[i] == i {
			i += 1
			continue
		}
		// not equal: nums[k]/v current value should move to nums[v],
		if nums[nums[i]] == nums[i] {
			return nums[i]
		}
		x := nums[nums[i]]
		nums[nums[i]] = nums[i]
		nums[i] = x
		// 交换之后 i 这个 index 下的新值假如不等于 i，则会被忽略。导致漏判
	}
	return -1
}
