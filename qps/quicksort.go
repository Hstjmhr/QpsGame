package main

import (
	"fmt"
)

// QuickSort 实现快速排序
func QuickSort(arr []int, low, high int) {
	if low < high {
		// 获取分区后的基准索引
		pi := partition(arr, low, high)
		// 递归排序基准左侧和右侧的子数组
		QuickSort(arr, low, pi-1)
		QuickSort(arr, pi+1, high)
	}
}

// partition 函数用于分区操作
func partition(arr []int, low, high int) int {
	avgnum := low + (high-low)/2
	pivot := arr[avgnum] // 选择最后一个元素作为基准
	arr[avgnum], arr[high] = arr[high], arr[avgnum]
	i := low // 小于基准的元素的索引

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i] // 交换
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i] // 交换基准到正确位置
	return i
}
func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("原始数组:", arr)

	QuickSort(arr, 0, len(arr)-1) // 调用快速排序
	fmt.Println("排序后数组:", arr)
}
