package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"wells-go/application/dtos"
	"wells-go/application/usecases"
	"wells-go/response"
)

type UserController struct {
	usecase *usecases.UserUsecase
}

func NewUserController(usecase *usecases.UserUsecase) *UserController {
	return &UserController{usecase: usecase}
}

func (ctl *UserController) Register(c *gin.Context) {
	startedAt := time.Now()

	var req dtos.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c.Writer, http.StatusBadRequest, "invalid request body", err.Error(), startedAt)
		return
	}

	user, err := ctl.usecase.Register(req)
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusBadRequest, "failed to register user", err.Error(), startedAt)
		return
	}
	response.SuccessResponse(c.Writer, "user registered successfully", user, startedAt)
}

func (ctl *UserController) Login(c *gin.Context) {
	startedAt := time.Now()

	var req dtos.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c.Writer, http.StatusBadRequest, "invalid request body", err.Error(), startedAt)
		return
	}

	token, err := ctl.usecase.Login(req)
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusUnauthorized, "login failed", err.Error(), startedAt)
		return
	}
	response.SuccessResponse(c.Writer, "login success", gin.H{"token": token}, startedAt)
}

func (ctl *UserController) GetUsers(c *gin.Context) {
	startedAt := time.Now()

	users, err := ctl.usecase.GetUsers()
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusInternalServerError, "failed to fetch users", err.Error(), startedAt)
		return
	}
	response.SuccessResponse(c.Writer, "list of users", users, startedAt)
}

func (ctl *UserController) GetUserByID(c *gin.Context) {
	startedAt := time.Now()

	id := c.Param("id")
	user, err := ctl.usecase.GetUserByID(id)
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusNotFound, "user not found", err.Error(), startedAt)
		return
	}
	response.SuccessResponse(c.Writer, "user detail", user, startedAt)
}

func (ctl *UserController) UpdateUser(c *gin.Context) {
	startedAt := time.Now()

	id := c.Param("id")
	var req dtos.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c.Writer, http.StatusBadRequest, "invalid request body", err.Error(), startedAt)
		return
	}

	user, err := ctl.usecase.UpdateUser(id, req)
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusInternalServerError, "failed to update user", err.Error(), startedAt)
		return
	}
	response.SuccessResponse(c.Writer, "user updated successfully", user, startedAt)
}

func (ctl *UserController) DeleteUser(c *gin.Context) {
	startedAt := time.Now()

	id := c.Param("id")
	if err := ctl.usecase.DeleteUser(id); err != nil {
		response.ErrorResponse(c.Writer, http.StatusInternalServerError, "failed to delete user", err.Error(), startedAt)
		return
	}
	response.SuccessResponse(c.Writer, "user deleted successfully", nil, startedAt)
}

func (ctl *UserController) Logout(c *gin.Context) {
	startedAt := time.Now()
	response.SuccessResponse(c.Writer, "logout success", nil, startedAt)
}
