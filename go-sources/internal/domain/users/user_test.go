package users

import (
	"reflect"
	"testing"
)

func TestNewUser(t *testing.T) {
	type args struct {
		chatId int64
	}
	tests := []struct {
		name string
		args args
		want *User
	}{
		{
			name: "test success",
			args: struct{ chatId int64 }{chatId: int64(12)},
			want: &User{id: int64(12)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUser(tt.args.chatId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_ID(t *testing.T) {
	type fields struct {
		id int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name:   "get success userId",
			fields: fields{id: int64(12)},
			want:   int64(12),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				id: tt.fields.id,
			}
			if got := u.ID(); got != tt.want {
				t.Errorf("ID() = %v, want %v", got, tt.want)
			}
		})
	}
}
