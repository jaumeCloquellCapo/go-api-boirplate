package repository

import (
	"ApiRest/internal/storage"
	"reflect"
	"testing"
)

func TestUserRepositoryInit(t *testing.T) {
	type args struct {
		db *storage.DbStore
	}
	tests := []struct {
		name string
		args args
		want UserRepositoryInterface
	}{
		{
			name: "success",
			args: args{
				db: nil,
			},
			want: &userRepository{
				db: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserInit() = %v, want %v", got, tt.want)
			}
		})
	}
}
