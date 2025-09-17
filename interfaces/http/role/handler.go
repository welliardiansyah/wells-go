package role

import (
	"net/http"
	"time"
	"wells-go/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"wells-go/application/dtos"
	"wells-go/application/usecases"
)

type RoleHandler struct {
	usecase *usecases.RoleUsecase
}

func NewRoleController(usecase *usecases.RoleUsecase) *RoleHandler {
	return &RoleHandler{usecase: usecase}
}

func (h *RoleHandler) CreateRole(c *gin.Context) {
	startedAt := time.Now()

	var req dtos.CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c.Writer, http.StatusBadRequest, "Invalid request body", err.Error(), startedAt)
		return
	}

	res, err := h.usecase.CreateRole(req)
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusInternalServerError, "Failed to create role", err.Error(), startedAt)
		return
	}
	response.SuccessResponse(c.Writer, "Role created successfully", res, startedAt)
}

func (h *RoleHandler) GetAllRoles(c *gin.Context) {
	startedAt := time.Now()

	res, err := h.usecase.GetAllRoles()
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusInternalServerError, "Failed to fetch roles", err.Error(), startedAt)
		return
	}
	response.SuccessResponse(c.Writer, "Roles retrieved successfully", res, startedAt)
}

func (h *RoleHandler) GetRoleByID(c *gin.Context) {
	startedAt := time.Now()

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusBadRequest, "Invalid ID format", err.Error(), startedAt)
		return
	}

	res, err := h.usecase.GetRoleByID(id)
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusNotFound, "Role not found", err.Error(), startedAt)
		return
	}
	response.SuccessResponse(c.Writer, "Role retrieved successfully", res, startedAt)
}

func (h *RoleHandler) UpdateRole(c *gin.Context) {
	startedAt := time.Now()

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusBadRequest, "Invalid ID format", err.Error(), startedAt)
		return
	}

	var req dtos.UpdateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c.Writer, http.StatusBadRequest, "Invalid request body", err.Error(), startedAt)
		return
	}

	res, err := h.usecase.UpdateRole(id, req)
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusInternalServerError, "Failed to update role", err.Error(), startedAt)
		return
	}
	response.SuccessResponse(c.Writer, "Role updated successfully", res, startedAt)
}

func (h *RoleHandler) DeleteRole(c *gin.Context) {
	startedAt := time.Now()

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusBadRequest, "Invalid ID format", err.Error(), startedAt)
		return
	}

	if err := h.usecase.DeleteRole(id); err != nil {
		response.ErrorResponse(c.Writer, http.StatusInternalServerError, "Failed to delete role", err.Error(), startedAt)
		return
	}
	response.SuccessResponse(c.Writer, "Role deleted successfully", nil, startedAt)
}
