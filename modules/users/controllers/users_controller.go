package controller_user

import (
	"encoding/json"
	"net/http"
	user "policy/modules/users/domain"
	user_service "policy/modules/users/services"
	"policy/utils/api_response"
	"policy/utils/errors"

	"github.com/gin-gonic/gin"
)

// GetUsersById
// @Summary Get Users by Id
// @Produce  json
// @Tags Users
// @Param user_id path string true "User Id"
// @Success 200 {object} string "ok"
// @Router /users/{user_id} [get]
func GetUserById(c *gin.Context) {
	userId := c.Param("user_id")
	userObject, err := user_service.UserService.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}
	apiResponse, respErr := json.Marshal(userObject)
	if respErr != nil {
		c.JSON(http.StatusInternalServerError, respErr.Error())
		return
	}

	c.JSON(http.StatusOK, api_response.Data(apiResponse))
	return
}

// GetAllUsers
// @Summary Get all users
// @Produce  json
// @Tags Users
// @Success 200 {object} string "ok"
// @Router /users [get]
func GetAllUsers(c *gin.Context) {
	users, err := user_service.UserService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}
	apiResponse, respErr := json.Marshal(users)
	if respErr != nil {
		c.JSON(http.StatusInternalServerError, respErr.Error())
		return
	}

	c.JSON(http.StatusOK, api_response.Data(apiResponse))
	return
}

// AddUser
// @Summary Add User
// @Produce  json
// @Tags Users
// @Param user body user.UserSave true "User Data"
// @Success 200 {object} user.UserSave
// @Router /users [post]
func AddUser(c *gin.Context) {
	// read json body
	var userObject user.User
	if err := c.ShouldBindJSON(&userObject); err != nil {
		restError := errors.NewBadRequest(err.Error())
		c.JSON(restError.Status, restError)
		return
	}

	result, saveErr := user_service.UserService.CreateUser(userObject)
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
