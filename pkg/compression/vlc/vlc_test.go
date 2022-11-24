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

func TestEncode(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		str  string
		want []byte
	}{
		{
			name: "basic case",
			str:  "hi",
			want: []byte{52, 128},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoder := New()
			if got := encoder.Encode(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {

	tests := []struct {
		name        string
		encodedText []byte
		want        string
	}{
		{
			name:        "basic case",
			encodedText: []byte{52, 128},
			want:        "hi",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decoder := New()
			if got := decoder.Decode(tt.encodedText); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
