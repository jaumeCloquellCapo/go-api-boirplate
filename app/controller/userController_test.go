package controller

import (
	"ApiRest/app/service"
	"reflect"
	"testing"
)

func TestUserControllerInit(t *testing.T) {
	type args struct {
		service service.UserServiceInterface
	}
	tests := []struct {
		name string
		args args
		want UserControllerInterface
	}{
		{
			name: "success",
			args: args{
				service: nil,
			},
			want: &userController{
				service: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserController(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User controller = %v, want %v", got, tt.want)
			}
		})
	}
}