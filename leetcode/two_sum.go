package qdxl

func twoSum(nums []int, target int) []int {
	seenMap := make(map[int]int)
	for i, v := range nums {
		if j, ok := seenMap[target-v]; ok {
			return []int{j, i}
		}
		seenMap[v] = i
	}
	return nil
}
