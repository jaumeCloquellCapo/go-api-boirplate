package mock

import "ApiRest/app/model"

type MockAuthService struct {}
// 2
func (m *MockAuthService) Login(userLogin model.Credentials) (token model.TokenDetails, err error) {
	// 3
	return model.TokenDetails{
		AccessToken:  "123",
		RefreshToken: "123",
		AccessUUID:   "123",
		RefreshUUID:  "123",
		AtExpires:    0,
		RtExpires:    0,
	}, nil
}

func (m *MockAuthService) Logout(accessUUID string) error {
	return nil
}
func (m *MockAuthService) SignUp(UserSignUp model.CreateUser) (user *model.User, tokenDetails model.TokenDetails, err error) {
	return &model.User{
			ID:         0,
			Name:       UserSignUp.Name,
			LastName:   UserSignUp.LastName,
			Password:   nil,
			Email:       UserSignUp.Email,
			Country:  UserSignUp.Country,
			Phone:      UserSignUp.Phone,
			PostalCode:  UserSignUp.PostalCode,
		}, model.TokenDetails{
			AccessToken:  "0",
			RefreshToken: "0",
			AccessUUID:   "0",
			RefreshUUID:  "0",
			AtExpires:    0,
			RtExpires:    0,
		}, nil
}
func (m *MockAuthService) GetAuth(AccessUUID string) (int64, error) {
	return 0, nil
}

