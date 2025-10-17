package main

import "fmt"

func main() {
	// [44 130]
	fmt.Println(vbEncodeNumber(300))
	// [[44 130] [104 135]]
	numbers := []int{300, 1000}
	fmt.Println(vbEncode(numbers))
	// 300
	bytestreams := []int{44, 130}
	fmt.Println(vbDecode(bytestreams)) // 300
}

func vbEncodeNumber(n int) []int {
	bytes := []int{}
	for {
		mod := n & 127 //下位7bit取り出し
		bytes = append(bytes, mod)
		if n < 128 {
			break
		}
		n = n >> 7 //7bit右シフト
	}
	bytes[len(bytes)-1] += 128
	return bytes
}

func vbEncode(numbers []int) [][]int {
	bytestreams := [][]int{}
	for _, n := range numbers {
		bytes := vbEncodeNumber(n)
		bytestreams = append(bytestreams, bytes)
	}
	return bytestreams
}

// 後ろの値から復元する必要がある
func vbDecode(bytestreams []int) int {
	n := 0
	len := len(bytestreams)
	for i := len - 1; 0 <= i; i-- {
		if bytestreams[i] < 128 {
			n = 128*n + bytestreams[i]
		} else {
			// 一番最初はこちらが実行される（例:130->2）
			n = 128*n + (bytestreams[i] - 128)
		}
	}
	return n
}
