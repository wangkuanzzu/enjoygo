package algorithm

func ArrayCountValues(args []int) (Status bool, MaxCount int, MaxValue []int) {
	// 验证数组的有效性
	if len(args) == 0 {
		return false, 0, nil
	}

	// 求出每个值对应出现的次数，使用map记录
	newMap := make(map[int]int)
	for _, value := range args {
		if newMap[value] != 0 {
			newMap[value]++
		} else {
			newMap[value] = 1
		}
	}

	// 求出出现最多的次数MaxCount
	MaxCount = 0
	MaxValue = nil
	for key, _ := range newMap {

		if MaxCount < newMap[key] {
			MaxCount = newMap[key]
		}
	}
	// 求出出现最多次数的值MaxValue
	for key, value := range newMap {
		if MaxCount == value {
			MaxValue = append(MaxValue, key)
		}
	}
	return true, MaxCount, MaxValue

}
