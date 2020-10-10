package main

import "fmt"
func lengthOfLongestSubstring(s string) int {
	// 使用滑动窗口的方法
	var lengthOfLongestSubstring int // flag，遇到重复的时候更改， i
	var currentQueue []int32
	for _,v := range s {
		for i,vv := range currentQueue {
			if vv == v {
				currentQueue = currentQueue[i+1:] // 切掉重复的
			}
		}
		// 不管重复不重复都追加
		currentQueue = append(currentQueue,v)
		if lengthOfLongestSubstring < len(currentQueue) {
			lengthOfLongestSubstring = len(currentQueue)
		}
	}
	return lengthOfLongestSubstring
}

func main()  {
	//str := "pwwkew"
	//str := "aabaab!bb"
	//str := "au"
	//str := "bbbbb"
	str := " "
	fmt.Println(	lengthOfLongestSubstring(str)	)
}
