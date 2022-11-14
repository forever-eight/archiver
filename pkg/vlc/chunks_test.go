package vlc

import (
	"reflect"
	"testing"
)

func Test_splitByChunks(t *testing.T) {
	type args struct {
		bStr      string
		chunkSize int
	}
	tests := []struct {
		name string
		args args
		want BinaryChunks
	}{
		{name: "basic case",
			args: args{
				bStr:      "1010101011111111",
				chunkSize: 8,
			},
			want: BinaryChunks{
				binaryChunk("10101010"),
				binaryChunk("11111111")},
		}, {name: "left runes",
			args: args{
				bStr:      "001101001",
				chunkSize: 8,
			},
			want: BinaryChunks{
				binaryChunk("00110100"),
				binaryChunk("10000000"),
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitByChunks(tt.args.bStr, tt.args.chunkSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitByChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binaryChunks_ToHex(t *testing.T) {
	tests := []struct {
		name string
		bcs  BinaryChunks
		want HexChunks
	}{
		{
			name: "basic case",
			bcs:  BinaryChunks{"1"},
			want: HexChunks{"01"}},
		{
			name: "more interesting case",
			bcs: BinaryChunks{
				binaryChunk("00110100"),
				binaryChunk("10000000"),
			},
			want: HexChunks{"34", "80"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bcs.ToHex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHexChunks(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		str  string
		want HexChunks
	}{
		{
			name: "basic case",
			str:  "34 80",
			want: HexChunks{"34", "80"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHexChunks(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHexChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hexChunk_ToBinary(t *testing.T) {
	tests := []struct {
		name string
		c    hexChunk
		want binaryChunk
	}{
		{
			name: "basic case",
			c:    hexChunk("34"),
			want: binaryChunk("00110100"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ToBinary(); got != tt.want {
				t.Errorf("ToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
