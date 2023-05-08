package admin_service

import (
	agent "policy/modules/agent/domain"
	"policy/utils/date_utils"
	"policy/utils/errors"
	"policy/utils/utils"
)

var (
	UserSerivce userServiceInterface = &userSerivce{}
)

type userSerivce struct {
}

type userServiceInterface interface {
	CreateUser(user agent.UserData) (*agent.UserData, *errors.RestErr)
	GetAllUsers() ([]agent.UserData, *errors.RestErr)
	GetUserById(string) (*agent.UserData, *errors.RestErr)
}

func (s *userSerivce) GetAllUsers() ([]agent.UserData, *errors.RestErr) {
	result := &agent.UserData{}
	return result.GetAllUsers()
}

func (s *userSerivce) GetUserById(id string) (*agent.UserData, *errors.RestErr) {
	result := &agent.UserData{Id: id}
	err := result.GetUserById()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *userSerivce) CreateUser(agent agent.UserData) (*agent.UserData, *errors.RestErr) {
	id, err := utils.GetRandomString()
	if err != nil {
		return nil, err
	}
	agent.Id = id
	date := date_utils.GetNowStringForDB()

	agent.CreatedDateTime = &date
	if err := agent.Save(); err != nil {
		return nil, err
	}
	return &agent, nil
}
