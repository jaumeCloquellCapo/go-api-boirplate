package service

import (
	"ApiRest/app/repository"
	"reflect"
	"testing"
)

func TestNewUserService(t *testing.T) {
	type args struct {
		userRepository repository.UserRepositoryInterface
	}
	tests := []struct {
		name string
		args args
		want UserServiceInterface
	}{
		{
			name: "success",
			args: args{
				userRepository: nil,
			},
			want: &userService{
				userRepo: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserService(tt.args.userRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserInit() = %v, want %v", got, tt.want)
			}
		})
	}
}
