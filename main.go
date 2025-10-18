package main

import (
	"fmt"

	"./vbcode"
)

func main() {
	// [44 130]
	// fmt.Println(vbEncodeNumber(300))
	// [[44 130] [104 135]]
	numbers := []int{300, 1000}
	fmt.Println(vbcode.VbEncode(numbers))
	// 300
	bytestreams := []int{44, 130}
	fmt.Println(vbcode.VbDecode(bytestreams)) // 300

	nums := []int{3, 5, 20, 21, 23, 76, 77, 78}
	fmt.Println(vbcode.Gap(nums)) // [3, 2, 15, 1, 2, 53, 1, 1]
}
