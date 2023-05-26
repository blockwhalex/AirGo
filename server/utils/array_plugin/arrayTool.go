package array_plugin

// 数组去重
func ArrayDeduplication(slice []int) []int {
	tempMap := make(map[int]struct{}, len(slice))
	j := 0
	for _, v := range slice {
		_, ok := tempMap[v]
		if ok {
			continue
		}
		tempMap[v] = struct{}{}
		slice[j] = v
		j++
	}
	return slice[:j]
}
