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
				BinaryChunk("10101010"),
				BinaryChunk("11111111")},
		}, {name: "left runes",
			args: args{
				bStr:      "001101001",
				chunkSize: 8,
			},
			want: BinaryChunks{
				BinaryChunk("00110100"),
				BinaryChunk("10000000"),
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

func TestBinaryChunks_Join(t *testing.T) {
	tests := []struct {
		name string
		bcs  BinaryChunks
		want string
	}{

		{name: "basic case",
			bcs: BinaryChunks{
				BinaryChunk("10101010"),
				BinaryChunk("11111111")},
			want: "1010101011111111",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bcs.Join(); got != tt.want {
				t.Errorf("Join() = %v, want %v", got, tt.want)
			}
		})
	}
}
