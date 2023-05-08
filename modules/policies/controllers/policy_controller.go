package controller_policy

import (
	"encoding/json"
	"net/http"
	policy "policy/modules/policies/domain"
	policy_service "policy/modules/policies/services"
	"policy/utils/api_response"
	"policy/utils/errors"

	"github.com/gin-gonic/gin"
)

// GetPoliciesById
// @Summary Get Polcies by Id
// @Produce  json
// @Tags Policies
// @Param policy_id path string true "Policy Id"
// @Success 200 {object} string "ok"
// @Router /policies/{policy_id} [get]
func GetPolicyById(c *gin.Context) {
	policiesId := c.Param("policy_id")
	polciiesObject, err := policy_service.PolicyService.GetPolicyById(policiesId)
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}
	apiResponse, respErr := json.Marshal(polciiesObject)
	if respErr != nil {
		c.JSON(http.StatusInternalServerError, respErr.Error())
		return
	}

	c.JSON(http.StatusOK, api_response.Data(apiResponse))
	return
}

// GetAllPolicies
// @Summary Get all policies
// @Produce  json
// @Tags Policies
// @Success 200 {object} string "ok"
// @Router /policies [get]
func GetAllPolicies(c *gin.Context) {
	policies, err := policy_service.PolicyService.GetAllPolicies()
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}
	apiResponse, respErr := json.Marshal(policies)
	if respErr != nil {
		c.JSON(http.StatusInternalServerError, respErr.Error())
		return
	}

	c.JSON(http.StatusOK, api_response.Data(apiResponse))
	return
}

// AddPolicy
// @Summary Add Policy
// @Produce  json
// @Tags Policies
// @Param policy body policy.PolicySave true "Policies Data"
// @Success 200 {object} policy.PolicySave
// @Router /policies [post]
func AddPolicy(c *gin.Context) {
	// read json body
	var policyObject policy.Policy
	if err := c.ShouldBindJSON(&policyObject); err != nil {
		restError := errors.NewBadRequest(err.Error())
		c.JSON(restError.Status, restError)
		return
	}

	result, saveErr := policy_service.PolicyService.CreatePolicy(policyObject)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	apiResponse, respErr := json.Marshal(result)
	if respErr != nil {
		c.JSON(http.StatusInternalServerError, respErr.Error())
		return
	}

	c.JSON(http.StatusOK, api_response.Data(apiResponse))
}
