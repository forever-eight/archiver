package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	chunkSize         = 8
	hexChunkSeparator = " "
)

type encodingTable map[rune]string

type HexChunks []hexChunk
type hexChunk string

type BinaryChunks []binaryChunk
type binaryChunk string

func (hcs HexChunks) ToString() string {

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
		buf.WriteString(hexChunkSeparator)
		buf.WriteString(string(hc))
	}

	return buf.String()
}

func NewHexChunks(str string) HexChunks {
	parts := strings.Split(str, hexChunkSeparator)

	res := make(HexChunks, 0, len(parts))
	for _, part := range parts {
		res = append(res, hexChunk(part))
	}

	return res
}

func (bcs *BinaryChunks) ToHex() HexChunks {
	res := make(HexChunks, 0, len(*bcs))

	for _, c := range *bcs {
		// chunk -> hex chunk
		hexC := c.ToHex()
		res = append(res, hexC)
	}

	return res
}

func (bc *binaryChunk) ToHex() hexChunk {
	num, err := strconv.ParseUint(string(*bc), 2, chunkSize)
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

func (hcs HexChunks) ToBinary() BinaryChunks {
	res := make(BinaryChunks, 0, len(hcs))

	for _, chunk := range hcs {
		// hex chunk -> chunk
		bc := chunk.ToBinary()
		res = append(res, bc)
	}

	return res
}

func (c *hexChunk) ToBinary() binaryChunk {
	num, err := strconv.ParseUint(string(*c), 16, chunkSize)
	if err != nil {
		panic("parse hex chunk error:" + err.Error())
	}
	// 10 -> 00000010
	// b - binary (двоичное представление)
	return binaryChunk(fmt.Sprintf("%08b", num))
}

// splitByChunks splits binary string by chunks with given size
// example: 1010101011010101 -> 10101010 11010101
func splitByChunks(bStr string, chunkSize int) BinaryChunks {
	strLen := utf8.RuneCountInString(bStr)
	chunksCount := strLen / chunkSize

	if strLen%chunkSize != 0 {
		chunksCount++
	}

	res := make(BinaryChunks, 0, chunksCount)

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

// Join joins binary chunks and return as string
func (bcs *BinaryChunks) Join() string {
	var buf strings.Builder

	for _, chunk := range *bcs {
		buf.WriteString(string(chunk))
	}
	return buf.String()
}
