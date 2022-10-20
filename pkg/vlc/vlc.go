package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

const chunkSize = 8

type encodingTable map[rune]string

type hexChunks []hexChunk
type hexChunk string

func (hcs hexChunks) ToString() string {
	const sep = " "
	switch len(hcs) {
	case 0:
		return ""
	case 1:
		return string(hcs[0])
	}

	var buf strings.Builder

	// to avoid space at the end problem
	buf.WriteString(string(hcs[0]))

	for _, hc := range hcs[1:] {
		buf.WriteString(" ")
		buf.WriteString(string(hc))
	}

	return buf.String()
}

type binaryChunks []binaryChunk
type binaryChunk string

func (bcs binaryChunks) ToHex() hexChunks {
	res := make(hexChunks, 0, len(bcs))

	for _, c := range bcs {
		// chunk -> hex chunk
		hexC := c.ToHex()
		res = append(res, hexC)
	}

	return res
}

func (bc binaryChunk) ToHex() hexChunk {
	num, err := strconv.ParseUint(string(bc), 2, chunkSize)
	if err != nil {
		panic("parse int error:" + err.Error())
	}
	res := strings.ToUpper(fmt.Sprintf("%x", num))

	// to fix len (1 -> 01)
	if len(res) == 1 {
		res = "0" + res
	}

	return hexChunk(res)
}

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

// splitByChunks splits binary string by chunks with given size
// example: 1010101011010101 -> 10101010 11010101
func splitByChunks(bStr string, chunkSize int) binaryChunks {
	strLen := utf8.RuneCountInString(bStr)
	chunksCount := strLen / chunkSize

	if strLen%chunkSize != 0 {
		chunksCount++
	}

	res := make(binaryChunks, 0, chunksCount)

	var buf strings.Builder
	for i, s := range bStr {
		buf.WriteString(string(s))

		if (i+1)%chunkSize == 0 {
			res = append(res, binaryChunk(buf.String()))
			buf.Reset()
		}
	}

	// if we have left runes
	if buf.Len() != 0 {
		lastChunk := buf.String()
		// repeat 0 to last chunk how many times?
		lastChunk += strings.Repeat("0", chunkSize-utf8.RuneCountInString(lastChunk))
		res = append(res, binaryChunk(lastChunk))
	}
	return res
}
