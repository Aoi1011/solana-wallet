package arrayshashing

import "sort"

// 1, 1, 1, 2, 2, 3
func topKFrequent(nums []int, k int) []int {
	var res []int
	firstM := make(map[int]int)
	secondM := make(map[int]int)

	// {1, 3}, {2, 2}, {3, 1}
	for _, num := range nums {
		firstM[num] += 1
	}

	//
	var temp []int
	for k, v := range firstM {
		temp = append(temp, v)
		secondM[v] = k
	}
	sort.Ints(temp)

	for i := len(temp) - 1; i >= len(temp)-k; i-- {
		sec := secondM[temp[i]]
		res = append(res, sec)
		secondM[temp[i]] = -1
	}
	return res
}

// key, value = count
