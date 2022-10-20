package vlc

import (
	"reflect"
	"testing"
)

func Test_prepareText(t *testing.T) {

	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "basic case",
			str:  "My name Forever_Eight",
			want: "!my name !forever_!eight",
		},
		{
			name: "lots of ! case",
			str:  "HI!",
			want: "!h!i!",
		}, {
			name: "no lower case",
			str:  "no lower case",
			want: "no lower case"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepareText(tt.str); got != tt.want {
				t.Errorf("prepareText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encodeBin(t *testing.T) {

	tests := []struct {
		name string
		str  string
		want string
	}{
		{name: "basic case",
			str:  "hi",
			want: "001101001"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeBin(tt.str); got != tt.want {
				t.Errorf("encodeBin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitByChunks(t *testing.T) {
	type args struct {
		bStr      string
		chunkSize int
	}
	tests := []struct {
		name string
		args args
		want binaryChunks
	}{
		{name: "basic case",
			args: args{
				bStr:      "1010101011111111",
				chunkSize: 8,
			},
			want: binaryChunks{
				binaryChunk("10101010"),
				binaryChunk("11111111")},
		}, {name: "left runes",
			args: args{
				bStr:      "001101001",
				chunkSize: 8,
			},
			want: binaryChunks{
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
		bcs  binaryChunks
		want hexChunks
	}{
		{
			name: "basic case",
			bcs:  binaryChunks{"1"},
			want: hexChunks{"01"}},
		{
			name: "more interesting case",
			bcs: binaryChunks{
				binaryChunk("00110100"),
				binaryChunk("10000000"),
			},
			want: hexChunks{"34", "80"},
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

func TestEncode(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "basic case",
			str:  "hi",
			want: "34 80",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.str); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
