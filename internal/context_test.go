package internal

import (
	"reflect"
	"testing"
)

func TestContext_process(t *testing.T) {
	type fields struct {
		text string
		hash string
	}
	tests := []struct {
		name   string
		fields fields
		want   *Context
	}{
		{
			name: "abc",
			fields: fields{
				text: "abc",
				hash: "",
			},
			want: &Context{
				text: "abc",
				hash: "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad",
			},
		},
		{
			name: "test",
			fields: fields{
				text: "test",
				hash: "",
			},
			want: &Context{
				text: "test",
				hash: "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
			},
		},
		{
			name: "empty",
			fields: fields{
				text: "",
				hash: "",
			},
			want: &Context{
				text: "",
				hash: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
			},
		},
		{
			name: "escape",
			fields: fields{
				text: "abc\"de\"fgh",
				hash: "",
			},
			want: &Context{
				text: "abc\"de\"fgh",
				hash: "480ed800368a2ec122098d275da7d7f049f7ea868ea3427992ee9275c285a194",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &Context{
				text: tt.fields.text,
				hash: tt.fields.hash,
			}
			ctx.process()
			if !reflect.DeepEqual(ctx, tt.want) {
				t.Errorf("process()=%v want %v", ctx, tt.want)
			}
		})
	}
}
