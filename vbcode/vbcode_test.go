package vbcode_test

import (
	"testing"
	"vbcode-project/vbcode"

	"github.com/stretchr/testify/assert"
)

func Test_VbEncodeAndVbDecode_Encodeした値をDecodeすると元に戻ることのテスト(t *testing.T) {
	tests := []struct {
		name  string
		bytes []int
	}{
		{"ソート済み数値配列", []int{1409171, 14711245, 18265928, 21590872}},
		{"ソート済み数値配列（1つしかない）", []int{541878}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encodedBinary := vbcode.VbEncode(tt.bytes)

			decodedNums := []int{}
			for _, num := range encodedBinary {
				dNum := vbcode.VbDecode(num)
				decodedNums = append(decodedNums, dNum)
			}

			assert.Equal(t, decodedNums, tt.bytes)
		})
	}
}
