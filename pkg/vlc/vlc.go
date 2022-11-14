package vlc

import (
	"fmt"
	"strings"
	"unicode"
)

func Encode(str string) string {
	// prepare text M -> !m
	prepared := prepareText(str)

	// encode to binary 10101101
	bStr := encodeBin(prepared)

	// bits to bytes (8) 10101010 10110101 (split by chunks)
	chunks := splitByChunks(bStr, chunkSize)

	// bytes to hex (2C)

	// return hexChunksStr
	return chunks.ToHex().ToString()
}

// prepareText prepares text to be fit to encode
// example: M -> !m
func prepareText(str string) string {
	var buf strings.Builder
	for _, s := range str {
		if unicode.IsUpper(s) {
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(s))
		} else {
			buf.WriteRune(s)
		}
	}

	return buf.String()
}

// encodeBin encodes to binary
func encodeBin(str string) string {
	var buf strings.Builder
	for _, s := range str {
		buf.WriteString(bin(s))
	}
	return buf.String()
}

func bin(r rune) string {
	table := getEncodingTable()
	res, ok := table[r]
	if !ok {
		panic("unknown character:" + string(r))
	}

	return res
}

func getEncodingTable() encodingTable {
	return encodingTable{
		' ': "11",
		't': "1001",
		'n': "10000",
		's': "0101",
		'r': "01000",
		'd': "00101",
		'!': "001000",
		'c': "000101",
		'm': "000011",
		'g': "0000100",
		'b': "0000010",
		'v': "00000001",
		'k': "0000000001",
		'q': "000000000001",
		'e': "101",
		'o': "10001",
		'a': "011",
		'i': "01001",
		'h': "0011",
		'l': "001001",
		'u': "00011",
		'f': "000100",
		'p': "0000101",
		'w': "0000011",
		'y': "0000001",
		'j': "000000001",
		'x': "00000000001",
		'z': "000000000000",
	}

}

func Decode(encodedText string) string {
	// hex chunks -> binary chunk
	hexChunks := NewHexChunks(encodedText)
	// bChunks -> binary string
	binaryStrings := hexChunks.ToBinary()
	fmt.Println(binaryStrings)
	// build decoding tree

	// dTree(bString) -> text

	// return text
	return ""
}
