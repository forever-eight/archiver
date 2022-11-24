package vlc

import (
	"reflect"
	"testing"
)

func Test_encodingTable_Node(t *testing.T) {
	tests := []struct {
		name string
		et   encodingTable
		want *Node
	}{
		{
			name: "base one",
			et: encodingTable{
				'a': "11",
				'b': "1001",
				'z': "0101",
			},
			want: &Node{
				Zero: &Node{
					One: &Node{
						Zero: &Node{
							One: &Node{
								Value: "z",
							},
						},
					},
				},
				One: &Node{
					Zero: &Node{
						Zero: &Node{
							One: &Node{
								Value: "b",
							},
						},
					},
					One: &Node{
						Value: "a",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.et.DecodingTree(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node() = %v, want %v", got, tt.want)
			}
		})
	}
}
