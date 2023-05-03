package internal_test

import (
	"reflect"
	"testing"

	"github.com/Lium1126/hexdump/internal"
)

func TestContextProcess(t *testing.T) {
	type fields struct {
		text string
	}

	tests := []struct {
		name   string
		fields fields
		want   internal.Context
	}{
		{
			name: "abc",
			fields: fields{
				text: "abc",
			},
			want: *internal.NewHashSettedContext("abc", "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"),
		},
		{
			name: "test",
			fields: fields{
				text: "test",
			},
			want: *internal.NewHashSettedContext("test", "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"),
		},
		{
			name: "empty",
			fields: fields{
				text: "",
			},
			want: *internal.NewHashSettedContext("", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"),
		},
		{
			name: "escape",
			fields: fields{
				text: "abc\"de\"fgh",
			},
			want: *internal.NewHashSettedContext("abc\"de\"fgh", "480ed800368a2ec122098d275da7d7f049f7ea868ea3427992ee9275c285a194"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := internal.NewContext(tt.fields.text)
			internal.Process(ctx)
			if !reflect.DeepEqual(*ctx, tt.want) {
				t.Errorf("process()=%v want %v", ctx, tt.want)
			}
		})
	}
}
