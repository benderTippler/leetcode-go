package code1

func TwoSum(nums []int, target int) []int {
	//2、 hash优化版
	numsIndex := make(map[int]int)
	for i2, v2 := range nums {
		if j, ok := numsIndex[target-v2]; ok {
			return []int{j, i2}
		}
		numsIndex[v2] = i2
	}
	return nil

	//1、自己写的代码
	for i, v := range nums {
		for ii, vv := range nums {
			if v+vv == target && i != ii {
				return []int{i, ii}
			}
		}
	}
	return nil
}
