package main

import "fmt"

//冒泡排序

var nums []int

// 时间复杂度
// 最坏情况下O(n方)
// 最好情况下O(n)
// 空间复杂度O(1)
func main() {
	nums = []int{1, 3, 2, 5, 7, 4, 6}
	abortnums := maopao(nums)
	fmt.Println(abortnums)

}

// 冒泡排序
func maopao(nums []int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		isabort := false
		for j := 0; j < n-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				isabort = true
			}
		}
		if !isabort {
			break
		}
	}
	return nums
}
