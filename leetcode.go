package main

import "fmt"

/*
 * 两数之和 暴力循环法
 * @nums 要查找的数组
 * @target 和
 */
func twoSum(nums []int, target int) []int {
	res := make([]int,0,len(nums))
	for i := 0;i < len(nums) - 1; i++ {
		for j:=i+1;j < len(nums);j++ { //这里j = i+1 保证重复的元素不会再被用到
			if target == nums[i] + nums[j] {
				res = append(res,i)
				res = append(res,j)
				return res
			}
		}
	}
	return []int{-1,-1}
}

/*
 * 两数之和 哈希表（集合）
 * @nums 要查找的数组
 * @target 和
 */
func twoSumHash(nums []int, target int) []int {
	//创建一个集合保存未使用的元素
	indexMap := make(map[int]int)
	for k,v := range nums {
		temp := target - v
		if value,ok := indexMap[temp]; ok { //检测是否在集合中存在 存在返回值和true
			return []int{value,k}
		}
		indexMap[v] = k
	}
	return []int{-1,-1}
}

/*
 * 约瑟夫环  递归实现
 * n个人排成一圈。从某个人开始，按顺时针方向依次编号。从编号为1的人开始顺时针报数，报到m的人退出圈子。这样不断循环下去，
 * 圈子里的人将不断减少。由于人的个数是有限的，因此最终会剩下一个人。试问最后剩下的人最开始的编号。
 * @n 初始长度
 * @m 报数到m就出列
 */
func lastRemaining(n int, m int) int {
	if n < 1 || m < 1 {
		return -1
	}
	result := 0 
	for i := 1; i <= n; i++ {
		result = (result + m) % i 
	}
	return result + 1
}

/* 杨辉三角
 *
 */
func triangles(n int)  {
	// number := 1
	arr := make([][]int,0)
	arr = append(arr, []int{1}) // 第一行是1
	fmt.Println(arr)
	fmt.Println([]int{1})
	for i := 1; i < n; i++ {  // i 是 行
		first := []int{1} // 行头部添加元素 1
		for j := 1; j <= i - 1; j++ {  // j 是列
			first = append(first, arr[i-1][j-1] + arr[i-1][j])
		}
		first = append(first, 1) // 给行末尾添加1
		fmt.Println(first)
		arr = append(arr, first)
	}
	fmt.Println(arr)
}
