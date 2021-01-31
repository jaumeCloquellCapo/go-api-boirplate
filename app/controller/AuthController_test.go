package controller

import (
	"ApiRest/app/service"
	"reflect"
	"testing"
)

func TestAuthControllerInit(t *testing.T) {
	type args struct {
		authService service.AuthServiceInterface
		userService service.UserServiceInterface
	}
	tests := []struct {
		name string
		args args
		want AuthControllerInterface
	}{
		{
			name: "success",
			args: args{
				authService: nil,
				userService: nil,
			},
			want: &authController{
				authService: nil,
				userService: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthController(tt.args.authService, tt.args.userService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Auth controller = %v, want %v", got, tt.want)
			}
		})
	}
}