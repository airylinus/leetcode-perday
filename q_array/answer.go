package q_array

// MaxMate 乘积最大的
func MaxMate(inputArray []int32) (max int32) {
	i := 1
	max = inputArray[0] * inputArray[1]
	for i < len(inputArray)-1 {
		nMax := inputArray[i] * inputArray[i+1]
		if nMax > max {
			max = nMax
		}
		i += 1
	}
	return
}

// MaxSubArray leetcode  @fixme
func MaxSubArray(arr []int32) (max int32) {
	if len(arr) == 1 {
		return arr[0]
	}
	//pre := MaxSubArray(arr[:len(arr)-1])

	return
}
