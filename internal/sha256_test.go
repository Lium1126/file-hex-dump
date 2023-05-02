package internal

import (
	"reflect"
	"testing"
)

func Test_getSHA256Binary(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "abc",
			args: args{"abc"},
			want: []byte{186, 120, 22, 191, 143, 1, 207, 234, 65, 65, 64, 222, 93, 174, 34, 35, 176, 3, 97, 163, 150, 23, 122, 156, 180, 16, 255, 97, 242, 0, 21, 173},
		},
		{
			name: "test",
			args: args{"test"},
			want: []byte{159, 134, 208, 129, 136, 76, 125, 101, 154, 47, 234, 160, 197, 90, 208, 21, 163, 191, 79, 27, 43, 11, 130, 44, 209, 93, 108, 21, 176, 240, 10, 8},
		},
		{
			name: "empty",
			args: args{""},
			want: []byte{227, 176, 196, 66, 152, 252, 28, 20, 154, 251, 244, 200, 153, 111, 185, 36, 39, 174, 65, 228, 100, 155, 147, 76, 164, 149, 153, 27, 120, 82, 184, 85},
		},
		{
			name: "escape",
			args: args{"abc\"de\"fgh"},
			want: []byte{72, 14, 216, 0, 54, 138, 46, 193, 34, 9, 141, 39, 93, 167, 215, 240, 73, 247, 234, 134, 142, 163, 66, 121, 146, 238, 146, 117, 194, 133, 161, 148},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSHA256Binary(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSHA256Binary() = %v, want %v", got, tt.want)
			}
		})
	}
}
