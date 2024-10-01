package words

import (
	"reflect"
	"testing"
)

func TestNewWord(t *testing.T) {
	type args struct {
		id          int
		origin      string
		translation string
	}
	tests := []struct {
		name string
		args args
		want *Word
	}{
		{
			name: "new word successfully",
			args: args{
				id:          0,
				origin:      "hi",
				translation: "привет",
			},
			want: &Word{
				id:          0,
				origin:      "hi",
				translation: "привет",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWord(tt.args.id, tt.args.origin, tt.args.translation); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWord_ID(t *testing.T) {
	type fields struct {
		id          int
		origin      string
		translation string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "get id successfully",
			fields: fields{
				id:          2,
				origin:      "hi",
				translation: "привет",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Word{
				id:          tt.fields.id,
				origin:      tt.fields.origin,
				translation: tt.fields.translation,
			}
			if got := w.ID(); got != tt.want {
				t.Errorf("ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWord_Origin(t *testing.T) {
	type fields struct {
		id          int
		origin      string
		translation string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "get origin successfully",
			fields: fields{
				id:          3,
				origin:      "hi",
				translation: "привет",
			},
			want: "hi",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Word{
				id:          tt.fields.id,
				origin:      tt.fields.origin,
				translation: tt.fields.translation,
			}
			if got := w.Origin(); got != tt.want {
				t.Errorf("Origin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWord_String(t *testing.T) {
	type fields struct {
		id          int
		origin      string
		translation string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "get string successfully",
			fields: fields{
				id:          3,
				origin:      "hi",
				translation: "привет",
			},
			want: "hi - привет",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Word{
				id:          tt.fields.id,
				origin:      tt.fields.origin,
				translation: tt.fields.translation,
			}
			if got := w.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWord_Translation(t *testing.T) {
	type fields struct {
		id          int
		origin      string
		translation string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "get translation successfully",
			fields: fields{
				id:          4,
				origin:      "hi",
				translation: "привет",
			},
			want: "привет",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Word{
				id:          tt.fields.id,
				origin:      tt.fields.origin,
				translation: tt.fields.translation,
			}
			if got := w.Translation(); got != tt.want {
				t.Errorf("Translation() = %v, want %v", got, tt.want)
			}
		})
	}
}
