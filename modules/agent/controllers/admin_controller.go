package controller_admin

import (
	"encoding/json"
	"net/http"
	user "policy/modules/agent/domain"
	admin_service "policy/modules/agent/services"
	"policy/utils/api_response"
	"policy/utils/errors"

	"github.com/gin-gonic/gin"
)

// GetUsersById
// @Summary Get Users by Id
// @Produce  json
// @Tags Agent Users
// @Param user_id path string true "User Id"
// @Success 200 {object} string "ok"
// @Router /agent/{user_id} [get]
func GetUsersById(c *gin.Context) {
	userId := c.Param("user_id")
	userObj, err := admin_service.UserSerivce.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}
	apiResponse, respErr := json.Marshal(userObj)
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
// @Tags Agent Users
// @Success 200 {object} string "ok"
// @Router /agent [get]
func GetAllUsersinAdmin(c *gin.Context) {
	user, err := admin_service.UserSerivce.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}
	apiResponse, respErr := json.Marshal(user)
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
// @Tags Agent Users
// @Param users body user.UserDataSave true "User Data"
// @Success 200 {object} user.UserDataSave
// @Router /agent [post]
func AddUserToAdmin(c *gin.Context) {
	// read json body
	var userObj user.UserData
	if err := c.ShouldBindJSON(&userObj); err != nil {
		restError := errors.NewBadRequest(err.Error())
		c.JSON(restError.Status, restError)
		return
	}

	result, saveErr := admin_service.UserSerivce.CreateUser(userObj)
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
