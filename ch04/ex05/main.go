// ex05 写一个函数在原地消除 []string 中相邻重复的字符串
package main

import (
	"fmt"
)

func main() {
	s := []string{"a", "a", "b", "c", "c", "c", "a", "d", "e"}
	s = removeDup(s)
	fmt.Println(s)

	s2 := removeDup2(s)
	fmt.Println(s2)
}

// 第一种实现方法
func removeDup(strings []string) []string {
	out := strings[:1]
	for _, s := range strings[1:] {
		if out[len(out)-1] == s { // out 中添加的都是不相邻重复的元素，所以用最后一个与遍历的元素比较，如果相同，则跳到下一次迭代
			continue
		} else {
			out = append(out, s) // 如果不相同，添加到 s 中
		}
	}
	return out
}

// 第二种实现方法
func removeDup2(s []string) []string {
	for i := 0; i < len(s)-1; {
		if s[i] == s[i+1] {
			copy(s[i:], s[i+1:]) // 如果相邻两个元素相同，则用 copy 函数将前一个删除
			s = s[:len(s)-1]     // 每一次 copy，最后一个元素都会重复一次，所以要进行一次 s[:len(s)-1] 删去此处操作带来的重复元素
		} else {
			i++ // 如相邻元素不同，才移动 i 到下一个元素
		}
	}
	return s
}
