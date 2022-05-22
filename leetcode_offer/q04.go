package leetcode_offer

import (
	"fmt"
)

/*
在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

示例:

现有矩阵 matrix 如下：

[
	[1,   4,  7, 11, 15],
	[2,   5,  8, 12, 19],
	[3,   6,  9, 16, 22],
	[10, 13, 14, 17, 24],
	[18, 21, 23, 26, 30]
]
给定 target = 5，返回 true。

给定 target = 20，返回 false。



来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// GenData 获取 matrix 矩阵
func GenData() [][]int {
	var d [][]int
	d = append(d, []int{1, 4, 7, 11, 15})
	d = append(d, []int{2, 5, 8, 12, 19})
	d = append(d, []int{3, 6, 9, 16, 22})
	d = append(d, []int{10, 13, 14, 17, 24})
	d = append(d, []int{18, 21, 23, 26, 30})
	d = append(d, []int{19, 22, 24, 27, 37})
	//d[3] = []int{}
	//d[4] = []int{10, 13, 14, 17, 24}
	//d[5] = []int{18, 21, 23, 26, 30}
	return d
}

/*
 * 思路错误
 * 思路：对角方向，右下是最大值。
 * 假设 matrix 是
*/

// IsInMatrix if needle in matrix
func IsInMatrix(matrix [][]int, needle, i, j int) (in bool, x, y int) {
	i = len(matrix)
	j = len(matrix[0])
	fmt.Println(i, j)
	for i > 0 || j > 0 {
		// found
		if matrix[i][j] == needle {
			return true, i, j
		}
		// bigger than max
		if matrix[i][j] < needle {
			return false, 0, 0
		}
		//
		//if
	}
	return false, 0, 0
	//for i < len(matrix) && j < len(matrix[0]) {
	//}
	return
}

/**
 *
 */


