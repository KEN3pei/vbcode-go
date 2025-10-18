package vbcode

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func VbEncodeNumber(n int) []int {
	bytes := []int{}
	for {
		mod := n & 127 //下位7bit取り出し
		bytes = append(bytes, mod)
		if n < 128 {
			break
		}
		n = n >> 7 //7bit右シフト 1000000000
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

func VbDecode(bytestreams []int) int {
	n := 0
	len := len(bytestreams)
	// 後ろの値から復元する必要がある
	for i := len - 1; 0 <= i; i-- {
		if bytestreams[i] < 128 {
			// 2. elseでnに入れた元の最後の8bitの値に+下位8bitを付け足して完了
			n = 128*n + bytestreams[i]
		} else {
			// 一番最初はこちらが実行される（例:130->2）
			// 1. 2の2進数10に2の7乗(128)をかけると10 0000 0000に戻せる
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

func TransferGapAndVBCode(path string, w *bufio.Writer) {
	f, _ := os.Open(path)
	defer f.Close()

	// tokenはデフォルトの行単位のまま
	sc := bufio.NewScanner(f)

	for sc.Scan() {
		line := sc.Text() // スキャンした内容を文字列で取得
		result := strings.Split(line, "\t")
		// gapで処理できる形式に変換
		strNums := strings.Split(result[1], ",")
		ints := parseInt(strNums)
		sInts := fmt.Sprintf("%v", VbEncode(Gap(ints)))
		// bufがいっぱいになるごとに書き込み
		w.WriteString(result[0] + "\t" + sInts + "\n")
	}
	// bufの残り分も書き込み
	w.Flush()
}

func parseInt(strs []string) []int {
	ints := []int{}
	for _, s := range strs {
		if s == "" {
			continue // 空文字をスキップ
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("変換エラー:", err)
			continue
		}
		ints = append(ints, n)
	}
	return ints
}
