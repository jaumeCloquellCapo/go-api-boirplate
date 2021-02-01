package mock

import "ApiRest/app/model"

type MockAuthService struct {}

func (m *MockAuthService) Login(userLogin model.Credentials) (token model.TokenDetails, err error) {
	return model.TokenDetails{}, nil
}

func (m *MockAuthService) Logout(accessUUID string) error {
	return nil
}
func (m *MockAuthService) SignUp(UserSignUp model.CreateUser) (user *model.User, tokenDetails model.TokenDetails, err error) {
	return &model.User{}, model.TokenDetails{}, nil
}
func (m *MockAuthService) GetAuth(AccessUUID string) (int64, error) {
	return 0, nil
}

