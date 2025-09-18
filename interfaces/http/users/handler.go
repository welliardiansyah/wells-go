package users

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"wells-go/application/dtos"
	"wells-go/application/usecases"
	"wells-go/infrastructure/redis"
	"wells-go/response"
)

type UserController struct {
	usecase *usecases.UserUsecase
}

func NewUserController(usecase *usecases.UserUsecase) *UserController {
	return &UserController{usecase: usecase}
}

func (ctl *UserController) Register(c *gin.Context) {

	var req dtos.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c.Writer, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	user, err := ctl.usecase.Register(req)
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusBadRequest, "failed to register user", err.Error())
		return
	}
	response.SuccessResponse(c.Writer, "user registered successfully", user)
}

func (ctl *UserController) Login(c *gin.Context) {

	var req dtos.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c.Writer, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	token, err := ctl.usecase.Login(req)
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusUnauthorized, "login failed", err.Error())
		return
	}
	response.SuccessResponse(c.Writer, "login success", gin.H{"token": token})
}

func (ctl *UserController) GetUsers(c *gin.Context) {

	users, err := ctl.usecase.GetUsers()
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusInternalServerError, "failed to fetch users", err.Error())
		return
	}
	response.SuccessResponse(c.Writer, "list of users", users)
}

func (ctl *UserController) GetUserByID(c *gin.Context) {

	id := c.Param("id")
	user, err := ctl.usecase.GetUserByID(id)
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusNotFound, "user not found", err.Error())
		return
	}
	response.SuccessResponse(c.Writer, "user detail", user)
}

func (ctl *UserController) UpdateUser(c *gin.Context) {

	id := c.Param("id")
	var req dtos.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c.Writer, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	user, err := ctl.usecase.UpdateUser(id, req)
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusInternalServerError, "failed to update user", err.Error())
		return
	}
	response.SuccessResponse(c.Writer, "user updated successfully", user)
}

func (ctl *UserController) DeleteUser(c *gin.Context) {

	id := c.Param("id")
	if err := ctl.usecase.DeleteUser(id); err != nil {
		response.ErrorResponse(c.Writer, http.StatusInternalServerError, "failed to delete user", err.Error())
		return
	}
	response.SuccessResponse(c.Writer, "user deleted successfully", nil)
}

func (ctl *UserController) Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		response.ErrorResponse(c.Writer, http.StatusUnauthorized, "Authorization header missing", nil)
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		response.ErrorResponse(c.Writer, http.StatusUnauthorized, "Invalid authorization header format", nil)
		return
	}

	token := parts[1]

	err := redis.Rdb.Del(context.Background(), "jwt:"+token).Err()
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusInternalServerError, "failed to delete token from redis", err.Error())
		return
	}

	response.SuccessResponse(c.Writer, "logout success", nil)
}
