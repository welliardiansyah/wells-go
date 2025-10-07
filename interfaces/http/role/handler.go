package role

import (
	"net/http"
	"strconv"
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
	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit
	search := c.Query("search")

	roles, total, err := h.usecase.GetAllRolesWithPagination(search, limit, offset)
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusInternalServerError, "Failed to fetch roles", err.Error())
		return
	}

	var res []*dtos.RoleResponse
	for _, r := range roles {
		res = append(res, r)
	}

	paging := dtos.PagingResponseFlat[*dtos.RoleResponse]{
		Page:  page,
		Limit: limit,
		Total: total,
		Data:  res,
	}

	response.SuccessResponsePaging(c.Writer, "Roles fetched successfully", paging)
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
