package service

import (
	"ApiRest/app/repository"
	"reflect"
	"testing"
)


func TestAuthRepositoryInit(t *testing.T) {
	type args struct {
		authRepository repository.AuthRepositoryInterface
		userRepository repository.UserRepositoryInterface
	}
	tests := []struct {
		name string
		args args
		want AuthServiceInterface
	}{
		{
			name: "success",
			args: args{
				authRepository: nil,
				userRepository: nil,
			},
			want: &authService{
				authRepository: nil,
				userRepository: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthService(tt.args.authRepository, tt.args.userRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserInit() = %v, want %v", got, tt.want)
			}
		})
	}
}