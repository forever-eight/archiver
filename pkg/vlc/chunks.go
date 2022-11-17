package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	chunkSize = 8
)

type encodingTable map[rune]string

type BinaryChunks []BinaryChunk
type BinaryChunk string

func NewBinChunks(data []byte) BinaryChunks {

	res := make(BinaryChunks, 0, len(data))
	for _, d := range data {
		res = append(res, NewBinChunk(d))
	}

	return res
}

func NewBinChunk(code byte) BinaryChunk {
	return BinaryChunk(fmt.Sprintf("%08b", code))
}

func (bcs *BinaryChunks) Bytes() []byte {
	res := make([]byte, 0, len(*bcs))

	for _, chunk := range *bcs {
		res = append(res, chunk.Byte())
	}

	return res
}

func (c *BinaryChunk) Byte() byte {
	num, err := strconv.ParseUint(string(*c), 2, chunkSize)
	if err != nil {
		panic("can't parse binary chunk" + err.Error())
	}
	return byte(num)
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
			res = append(res, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}

	// if we have left runes
	if buf.Len() != 0 {
		lastChunk := buf.String()
		// repeat 0 to last chunk how many times?
		lastChunk += strings.Repeat("0", chunkSize-utf8.RuneCountInString(lastChunk))
		res = append(res, BinaryChunk(lastChunk))
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
