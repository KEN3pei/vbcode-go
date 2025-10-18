package main

import (
	"bufio"
	"os"
	"vbcode-project/vbcode"
)

func main() {
	// fmt.Println(vbcode.VbEncodeNumber(300)) // [44 130]

	// numbers := []int{300, 1000}
	// fmt.Println(vbcode.VbEncode(numbers)) // [[44 130] [104 135]]

	// bytestreams := []int{44, 130}
	// fmt.Println(vbcode.VbDecode(bytestreams)) // 300

	// nums := []int{3, 5, 20, 21, 23, 76, 77, 78}
	// fmt.Println(vbcode.Gap(nums)) // [3, 2, 15, 1, 2, 53, 1, 1]

	// 実際のファイルを圧縮するコード
	f, _ := os.Create("output.txt")
	defer f.Close()

	w := bufio.NewWriter(f)
	vbcode.TransferGapAndVBCode("./vbcode/eid_tags.txt", w)
}
