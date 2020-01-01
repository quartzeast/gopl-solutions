package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6}
	reverse(&a)
	fmt.Println(a)
}

// reverse reverse a array of ints in place.
func reverse(a *[6]int) {
	for i, j := 0, len(a)-1; i < len(a)/2; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
