package mock

import "ApiRest/app/model"

type MockUserService struct {

}
func (m *MockUserService)  FindById(id int) (user *model.User, err error){
	return &model.User{}, nil
}
func (m *MockUserService) RemoveById(id int) error {
	return nil
}
func (m *MockUserService) UpdateById(id int, user model.UpdateUser) error {
	return nil
}
func (m *MockUserService) FindAll() ([]model.User, error) {
	return []model.User{}, nil
}
func (m *MockUserService) FindByEmail(email string) (user *model.User, err error) {
	return &model.User{}, nil
}

