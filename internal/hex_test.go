package internal_test

import (
	"testing"

	"github.com/Lium1126/hexdump/internal"
)

func TestEncodeHex(t *testing.T) {
	type args struct {
		b []byte
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "abc",
			args: args{
				[]byte{
					186, 120, 22, 191, 143, 1, 207, 234,
					65, 65, 64, 222, 93, 174, 34, 35,
					176, 3, 97, 163, 150, 23, 122, 156,
					180, 16, 255, 97, 242, 0, 21, 173,
				},
			},
			want: "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad",
		},
		{
			name: "test",
			args: args{
				[]byte{
					159, 134, 208, 129, 136, 76, 125, 101,
					154, 47, 234, 160, 197, 90, 208, 21,
					163, 191, 79, 27, 43, 11, 130, 44,
					209, 93, 108, 21, 176, 240, 10, 8,
				},
			},
			want: "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
		},
		{
			name: "empty",
			args: args{
				[]byte{
					227, 176, 196, 66, 152, 252, 28, 20,
					154, 251, 244, 200, 153, 111, 185, 36,
					39, 174, 65, 228, 100, 155, 147, 76,
					164, 149, 153, 27, 120, 82, 184, 85,
				},
			},
			want: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
		{
			name: "escape",
			args: args{
				[]byte{
					72, 14, 216, 0, 54, 138, 46, 193,
					34, 9, 141, 39, 93, 167, 215, 240,
					73, 247, 234, 134, 142, 163, 66, 121,
					146, 238, 146, 117, 194, 133, 161, 148,
				},
			},
			want: "480ed800368a2ec122098d275da7d7f049f7ea868ea3427992ee9275c285a194",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := internal.EncodeHex(tt.args.b); got != tt.want {
				t.Errorf("EncodeHex() = %v, want %v", got, tt.want)
			}
		})
	}
}
