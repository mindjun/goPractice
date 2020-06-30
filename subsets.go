package main

import "fmt"

// 计算子集
func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	track := make([]int, 0)
	backtrack(nums, 0, track, &result)
	return result
}

func backtrack(nums []int, pos int, track []int, result *[][]int) {
	// 进行拷贝操作
	//temp := make([]int, len(track))
	//copy(temp, track)
	// 	*result = append(*result, temp)

	*result = append(*result, track)
	for i := pos; i < len(nums); i++ {
		track = append(track, nums[i])
		backtrack(nums, i+1, track, result)
		// pop
		track = track[0 : len(track)-1]
	}
}

func main() {
	a := []int{0, 1, 2, 3}
	fmt.Println(subsets(a))
}
