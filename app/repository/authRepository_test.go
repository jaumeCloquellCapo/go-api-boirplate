package repository

import (
	"ApiRest/internal/storage"
	"reflect"
	"testing"
)


func TestAuthRepositoryInit(t *testing.T) {
	type args struct {
		redis *storage.DbCache
	}
	tests := []struct {
		name string
		args args
		want AuthRepositoryInterface
	}{
		{
			name: "success",
			args: args{
				redis: nil,
			},
			want: &authRepository{
				redis: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthRepository(tt.args.redis); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserInit() = %v, want %v", got, tt.want)
			}
		})
	}
}