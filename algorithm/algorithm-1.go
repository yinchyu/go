package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	// greeting = "Hello, Go"
	one = 1 << iota
	two
	three
	four
	fife
)

// startgorountine 启动goroutine 的顺序测试
func startgorountine() int {
	// 启动之后就和这个函数没有关系了， 可以在函数的末尾设置  wg.wait
	// 或者在main 函数的开始设置结束
	for i := 0; i < 10; i++ {
		go func() {
			for {
				fmt.Println("alive")
				time.Sleep(time.Second * 1)
			}
		}()
	}
	return 10
}
func avg(s []int) float64 {
	temp := 0
	for _, value := range s {
		temp += value
	}
	return float64(temp) / float64(len(s))
}
func findindex(s []int, target int) int {
	for index, value := range s {
		if value == target {
			return index
		}
	}
	return 0
}

// IotaTest 测试Iota 的初始化值
func IotaTest() {
	fmt.Println(one, two, three, four, fife)
	// 1 2 4 8 16
}

// PalinDrome 回文替换的最小代价
func PalinDrome(str string) int {
	add := []int{0, 100, 200, 360, 220}
	dec := []int{0, 120, 350, 200, 320}
	n := len(str)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if str[i] == str[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp1 := dp[i+1][j] + min(add[str[i]-'0'], dec[str[i]-'0'])
				dp2 := dp[i][j-1] + min(add[str[j]-'0'], dec[str[j]-'0'])
				dp[i][j] = min(dp1, dp2)
			}
		}
	}
	return dp[0][n-1]
}

// difftimer 比较两个时间的不同
func difftimer() {
	timer, _ := time.Parse("15:04", "10:00")
	timer2, _ := time.Parse("15:04", "11:00")
	if timer.Sub(timer2) < 0 {
		fmt.Println("hello")
	}
}

//countMaxActivity  不同时间区间进行合并
func countMaxActivity(timeSchedule [][]string) int {
	// write code here
	//将timeschedule 转换成时间数据
	sort.Slice(timeSchedule, func(i, j int) bool {
		timer, _ := time.Parse("15:04", timeSchedule[i][0])
		timer2, _ := time.Parse("15:04", timeSchedule[j][0])
		return timer.Sub(timer2) < 0
	})
	// 排完序之后的数据
	starttime, _ := time.Parse("15:04", timeSchedule[0][1])
	counter := 1
	for i := 1; i < len(timeSchedule); i++ {
		timer1, _ := time.Parse("15:04", timeSchedule[i][0])
		timer2, _ := time.Parse("15:04", timeSchedule[i][1])
		if starttime.Sub(timer1) > 0 {
			if starttime.Sub(timer2) > 0 {
				starttime = timer2
			}
		} else {
			counter++
			starttime = timer2
		}
	}
	return counter
}

//findMin  找到最短的路径
func findMin(mapArray [][]int) int {
	// write code here
	m, n := len(mapArray), len(mapArray[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	// 初始化第一行和第一列的元素
	temp := 0
	for i := range dp[0] {
		temp += mapArray[0][i]
		dp[0][i] = temp
	}
	temp = 0
	for i := range dp {
		temp += mapArray[i][0]
		dp[i][0] = temp
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + mapArray[i][j]
		}
	}
	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func largestRectangleArea(heights []int) int {
	heights = append(append([]int{0}, heights...), 0)
	stack := make([]int, 0)
	res := 0
	for i := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			temp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = max(res, (i-stack[len(stack)-1]-1)*heights[temp])
		}
		stack = append(stack, i)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	dp := make([]int, len(temperatures))
	for i := range temperatures {
		if len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			temp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			dp[temp] = i - temp
		}
		stack = append(stack, i)
	}
	return dp
}

//ReplaceCompare 将所有的ab 替换成bba 的次数
func ReplaceCompare(str string) int {
	counter := 0
	bnum := 0
	mod := 1000000007
	if len(str) < 2 {
		return counter
	}
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == 'b' {
			bnum++
		}
		if str[i] == 'a' {
			// 通过简单的计数和反转操作
			counter += bnum % mod
			bnum = bnum * 2
		}
	}
	return counter % mod
}

// DeferFunc 观察defer的输出顺序
func DeferFunc() {
	var out []*int
	for i := 0; i < 3; i++ {
		out = append(out, &i)
		fmt.Println(i)
	}

	fmt.Println(*out[0], *out[1], *out[2])
	for i := 0; i < 5; i++ {
		// defer 的执行顺序是按照栈的执行顺序进行操作的
		defer func(i int) { fmt.Println(i) }(i)
	}
	for i := 0; i < 5; i++ {
		// defer 的执行顺序是按照栈的执行顺序进行操作的,所有的输出都是4
		defer func() { fmt.Println(i) }()
	}
}

func heapify(nums []int, start int) {
	leftindex := start*2 + 1
	rightindex := start*2 + 2
	nextindex := start
	if leftindex < len(nums) {
		if rightindex < len(nums) && nums[rightindex] > nums[leftindex] {
			nextindex = rightindex
		} else {
			nextindex = leftindex
		}
	}
	if nums[start] < nums[nextindex] {
		nums[start], nums[nextindex] = nums[nextindex], nums[start]
		heapify(nums, nextindex)
	}
}

func heapsort(nums []int) []int {
	n := len(nums)
	for i := n / 2; i >= 0; i-- {
		heapify(nums, i)
	}
	// 总体排序操作
	for j := n - 1; j > 0; j-- {
		nums[0], nums[j] = nums[j], nums[0]
		heapify(nums[:j], 0)
	}
	return nums
}

//maxProfit 股票的最大利润 使用单调栈进行求解
func maxProfit(prices []int) int {
	stack := make([]int, 0)
	for _, val := range prices {
		for len(stack) > 0 && val < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, val)
	}
	if len(stack) > 1 {
		return stack[len(stack)-1] - stack[0]
	} else {
		return 0
	}
}

// MinChange  找硬币这个问题直接使用一维动态规划进行求解， 和背包问题有些类似
func MinChange(amount int, coins []int) int {
	dp := make([][]int, len(coins)+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1
	}

	for j, coin := range coins {
		for i := 1; i <= amount; i++ {
			if i < coin {
				dp[j+1][i] = dp[j][i]
			} else {
				dp[j+1][i] = dp[j][i] + dp[j+1][i-coin]
			}

		}
	}
	// fmt.Println(dp)
	return dp[len(coins)][amount]
}

// heapsorttest 百度一面 手写堆排序
func heapsorttest() {
	nums := []int{4, 5, 6, 2, 1}
	fmt.Println("sort, result", heapsort(nums))
}

func ReplaceCompareTest() {
	for {
		var a string
		h, err := fmt.Scanln(&a)
		if err != nil {
			fmt.Println(err, h)
		}
		fmt.Println(ReplaceCompare(a))
	}
}

func countMaxActivityTest() {
	timeschedule := [][]string{{"10:00", "12:00"}, {"03:00", "11:30"}, {"11:30", "14:00"}}
	fmt.Println(countMaxActivity(timeschedule))
}

func findMinTest() {
	fmt.Println(findMin([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}
func PalinDromeTest() {
	fmt.Println(PalinDrome("123123"))
}

// avg  计算一个字符串的平均值
func average(a string) {
	// a="5,6,8,26,50,48,52,55,10,1,2,1,20,5:3"
	newstr := strings.Split(a, ",")
	list := make([]int, len(newstr))
	windows := 0
	for i := range newstr {
		val, err := strconv.Atoi(newstr[i])
		if err != nil {
			temp := strings.Split(newstr[i], ":")
			windows, _ = strconv.Atoi(temp[1])
			val, _ = strconv.Atoi(temp[0])
		}
		list[i] = val
	}
	startvalue := 0
	res := 0.0
	for i := 0; i < windows; i++ {
		startvalue += list[i]
	}
	for i := windows; i < len(list); i++ {
		diff := list[i] - list[i-windows]
		temp := float64(diff) / float64(startvalue)
		if temp > res {
			res = temp
		}
		startvalue += diff
	}
	fmt.Printf("%.2f%%", res*100)
}

// 晴天抽水， 雨天下雨， 问水库会不会满
func rain(str string) {
	flagmap := make(map[string]bool)
	newstr := strings.Split(str[1:len(str)-1], ",")
	index := 0
	res := make([]string, len(newstr))
	for i := range newstr {
		// 表示有雨
		if newstr[i] != "0" {
			// 没有办法阻止洪水
			if flagmap[newstr[i]] {
				fmt.Println("[]")
				return
			} else {
				// 标记水库有水
				flagmap[newstr[i]] = true
				res[i] = "-1"
			}
		} else {

			if index < i+1 {
				index = i + 1
			}
			// 向后遍历
			for index < len(newstr) {
				if newstr[index] != "0" {
					// 这个问题更新的时候有很多的问题， 直接复制过来导致没有修改内部的索引状态， 导致检查了好长时间
					// 如果不是特别长的变量最好还是自己敲，避免这种低级的问题发生。
					flagmap[newstr[index]] = false
					// 将对应的水库进行标记
					res[i] = newstr[index]
					index++
					break
				}
				index++
			}
		}
		if index < i {
			index = i
		}
	}

	fmt.Println("[" + strings.Join(res, ",") + "]")
}
