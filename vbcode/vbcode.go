package vbcode

func VbEncodeNumber(n int) []int {
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

func VbEncode(numbers []int) [][]int {
	bytestreams := [][]int{}
	for _, n := range numbers {
		bytes := VbEncodeNumber(n)
		bytestreams = append(bytestreams, bytes)
	}
	return bytestreams
}

// 後ろの値から復元する必要がある
func VbDecode(bytestreams []int) int {
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

// numはsort済みであること
// 例: [3, 5, 20, 21, 23, 76, 77, 78] -> [3, 2, 15, 1, 2, 53, 1, 1]
func Gap(nums []int) []int {
	gapnums := []int{}
	for i, n := range nums {
		if i == 0 {
			gapnums = append(gapnums, n)
			continue
		}
		gap := n - nums[i-1]
		gapnums = append(gapnums, gap)
	}
	return gapnums
}
