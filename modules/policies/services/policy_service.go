package policy_service

import (
	policy "policy/modules/policies/domain"
	"policy/utils/date_utils"
	"policy/utils/errors"
	"policy/utils/utils"
)

var (
	PolicyService policyServiceInterface = &policyService{}
)

type policyService struct {
}

type policyServiceInterface interface {
	CreatePolicy(location policy.Policy) (*policy.Policy, *errors.RestErr)
	GetAllPolicies() ([]policy.Policy, *errors.RestErr)
	GetPolicyById(string) (*policy.Policy, *errors.RestErr)
}

func (s *policyService) GetAllPolicies() ([]policy.Policy, *errors.RestErr) {
	result := &policy.Policy{}
	return result.GetAllPolicies()
}

func (s *policyService) GetPolicyById(id string) (*policy.Policy, *errors.RestErr) {
	result := &policy.Policy{Id: id}
	err := result.GetPolicyById()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *policyService) CreatePolicy(policy policy.Policy) (*policy.Policy, *errors.RestErr) {
	id, err := utils.GetRandomString()
	if err != nil {
		return nil, err
	}
	policy.Id = id
	date := date_utils.GetNowStringForDB()

	policy.CreatedDateTime = &date
	if err := policy.Save(); err != nil {
		return nil, err
	}
	return &policy, nil
}
