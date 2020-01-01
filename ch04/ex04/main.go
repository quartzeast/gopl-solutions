// ex04 编写通过一次循环完成旋转的 rotate 函数
package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	s = rotate(s, 2)
	fmt.Println(s)

	s2 := []int{0, 1, 2, 3, 4, 5, 6}
	rotate2(s2, 3)
	fmt.Println(s2)
}

// 第一种实现方式
func rotate(s []int, p int) []int {
	for i := 0; i < p; i++ {
		s = append(s, s[i])
	}
	return s[p:]
}

// 第二种实现方式
func rotate2(s []int, p int) {
	num := p % len(s)
	double := append(s, s[:num]...)
	copy(s, double[num:num+len(s)])
}
