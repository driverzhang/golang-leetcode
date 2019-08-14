/*
 * golang 实现
 * 1. 冒泡排序、插入排序、选择排序、 O(n²)
 * 2. 归并排序、快速排序、         O(nlogn)
 * 3. 计数排序、基数排序、桶排序   O(n)
 */
package quickSort

import (
	"math/rand"
	"time"
)

// 获取随机数据源 随机种子传入随机个数的数据
func RandArray(num int) []int {
	arr := make([]int, num)
	rand.Seed(time.Now().Unix()) //以当前时间为随机数种子
	for i := 0; i < num; i++ {
		arr[i] = rand.Intn(100)
		
	}
	
	return arr
}

// 快排1 高明些
func QuickSort(items []int) (newItems []int) {
	if len(items) < 1 {
		return
	}
	
	key := items[rand.Intn(len(items))] // 随机选择基准 或者 itmes[0]
	head, tail := 0, len(items)-1
	
	for i := 1; i <= tail; {
		if items[i] > key { // 比 基准 大的值放后面
			items[i], items[tail] = items[tail], items[i] // 只要满足 该数比 基准key 大 则 i不变 把该数往后换位置 将tail--
			tail--
		} else { // 反之 小的放前面
			items[i], items[head] = items[head], items[i] // 若 该值比 基准key 小 则 tai不变 把该数往前换位置 将i++ head++
			head++
			i++
		}
	}
	
	items[head] = key
	QuickSort(items[:head])
	QuickSort(items[head+1:])
	newItems = items
	return
}

// 快排2 易懂
func QuickSort2(arr []int) (newItems []int) {
	if len(arr) < 1 {
		return
	}
	
	median := arr[0] // arr[rand.Intn(len(arr))] // 随机取 基准值
	
	low_part := make([]int, 0, len(arr))
	high_part := make([]int, 0, len(arr))
	middle_part := make([]int, 0, len(arr))
	
	for _, item := range arr {
		switch {
		case item < median:
			low_part = append(low_part, item)
		case item == median:
			middle_part = append(middle_part, item)
		case item > median:
			high_part = append(high_part, item)
		}
	}
	
	low_part = QuickSort2(low_part)
	high_part = QuickSort2(high_part)
	
	low_part = append(low_part, middle_part...)
	low_part = append(low_part, high_part...)
	newItems = low_part
	return
}

// 选择排序，a表示数组，n表示数组大小
func SelectionSort(a []int, n int) {
	if n <= 1 {
		return
	}
	for i := 0; i < n; i++ {
		// 查找最小值
		minIndex := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[minIndex] { // 按照从小到大排序
				minIndex = j
			}
		}
		// 交换
		a[i], a[minIndex] = a[minIndex], a[i]
		
	}
	
}
