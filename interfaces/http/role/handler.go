package role

import (
	"net/http"
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

	var req dtos.CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c.Writer, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	res, err := h.usecase.CreateRole(req)
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusInternalServerError, "Failed to create role", err.Error())
		return
	}
	response.SuccessResponse(c.Writer, "Role created successfully", res)
}

func (h *RoleHandler) GetAllRoles(c *gin.Context) {

	res, err := h.usecase.GetAllRoles()
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusInternalServerError, "Failed to fetch roles", err.Error())
		return
	}
	response.SuccessResponse(c.Writer, "Roles retrieved successfully", res)
}

func (h *RoleHandler) GetRoleByID(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusBadRequest, "Invalid ID format", err.Error())
		return
	}

	res, err := h.usecase.GetRoleByID(id)
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusNotFound, "Role not found", err.Error())
		return
	}
	response.SuccessResponse(c.Writer, "Role retrieved successfully", res)
}

func (h *RoleHandler) UpdateRole(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusBadRequest, "Invalid ID format", err.Error())
		return
	}

	var req dtos.UpdateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c.Writer, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	res, err := h.usecase.UpdateRole(id, req)
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusInternalServerError, "Failed to update role", err.Error())
		return
	}
	response.SuccessResponse(c.Writer, "Role updated successfully", res)
}

func (h *RoleHandler) DeleteRole(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusBadRequest, "Invalid ID format", err.Error())
		return
	}

	if err := h.usecase.DeleteRole(id); err != nil {
		response.ErrorResponse(c.Writer, http.StatusInternalServerError, "Failed to delete role", err.Error())
		return
	}
	response.SuccessResponse(c.Writer, "Role deleted successfully", nil)
}
