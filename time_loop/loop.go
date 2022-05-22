package time_loop

import "fmt"

func AssembleTime(inputNums []int32) map[string]struct{} {
	result := map[string]struct{}{}
	for k1, v1 := range inputNums {
		if v1 > 2 {
			continue
		}
		for k2, v2 := range inputNums {
			// skip the first number
			if k2 == k1 {
				continue
			}
			// skip > 23
			if v1*10+v2 > 23 {
				continue
			}
			for k3, v3 := range inputNums {
				if k3 == k1 || k3 == k2 {
					continue
				}
				if v3 > 5 {
					continue
				}
				for k4, v4 := range inputNums {
					if k4 == k3 || k4 == k2 || k4 == k1 {
						continue
					}
					if v3*10+v4 > 59 {
						continue
					}
					result[fmt.Sprintf("%d%d:%d%d", v1, v2, v3, v4)] = struct{}{}
					fmt.Println(v1, v2, v3, v4)
				}
			}
		}
	}
	fmt.Println(result)
	return result
}
