package policy_service

import (
	user "policy/modules/users/domain"
	"policy/utils/date_utils"
	"policy/utils/errors"
	"policy/utils/utils"
)

var (
	UserService userServiceInterface = &userService{}
)

type userService struct {
}

type userServiceInterface interface {
	CreateUser(user user.User) (*user.User, *errors.RestErr)
	GetAllUsers() ([]user.User, *errors.RestErr)
	GetUserById(string) (*user.User, *errors.RestErr)
}

func (s *userService) GetAllUsers() ([]user.User, *errors.RestErr) {
	result := &user.User{}
	return result.GetAllUsers()
}

func (s *userService) GetUserById(id string) (*user.User, *errors.RestErr) {
	result := &user.User{Id: id}
	err := result.GetUserById()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *userService) CreateUser(user user.User) (*user.User, *errors.RestErr) {
	id, err := utils.GetRandomString()
	if err != nil {
		return nil, err
	}
	user.Id = id
	date := date_utils.GetNowStringForDB()

	user.CreatedDateTime = &date
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
