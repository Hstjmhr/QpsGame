package main

import "fmt"

func main() {
	nums := []int{1, 4, 2, 8, 33, 4, 6, 2}
	num := xuanze(nums)
	fmt.Println(num)

}
func xuanze(nums []int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		minindex := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[minindex] {
				minindex = j
			}
		}
		nums[i], nums[minindex] = nums[minindex], nums[i]
	}
	return nums
}
