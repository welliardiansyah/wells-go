package permission

import (
	"net/http"
	"strconv"
	"wells-go/application/dtos"
	"wells-go/application/usecases"
	"wells-go/response"

	"github.com/gin-gonic/gin"
)

type PermissionHandler struct {
	usecase *usecases.PermissionUsecase
}

func NewPermissionController(usecase *usecases.PermissionUsecase) *PermissionHandler {
	return &PermissionHandler{usecase: usecase}
}

func (h *PermissionHandler) Create(c *gin.Context) {

	var req dtos.CreatePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c.Writer, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}
	permission, err := h.usecase.Create(req)
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusInternalServerError, "Failed to create permission", err.Error())
		return
	}
	response.SuccessResponse(c.Writer, "Permission created successfully", permission)
}

func (h *PermissionHandler) Update(c *gin.Context) {

	id := c.Param("id")
	var req dtos.UpdatePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c.Writer, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}
	permission, err := h.usecase.Update(id, req)
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusInternalServerError, "Failed to update permission", err.Error())
		return
	}
	response.SuccessResponse(c.Writer, "Permission updated successfully", permission)
}

func (h *PermissionHandler) Delete(c *gin.Context) {

	id := c.Param("id")
	if err := h.usecase.Delete(id); err != nil {
		response.ErrorResponse(c.Writer, http.StatusInternalServerError, "Failed to delete permission", err.Error())
		return
	}
	response.SuccessResponse(c.Writer, "Permission deleted successfully", nil)
}

func (h *PermissionHandler) FindByID(c *gin.Context) {

	id := c.Param("id")
	permission, err := h.usecase.FindByID(id)
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusNotFound, "Permission not found", err.Error())
		return
	}
	response.SuccessResponse(c.Writer, "Permission retrieved successfully", permission)
}

func (h *PermissionHandler) FindAll(c *gin.Context) {
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

	permissions, total, err := h.usecase.FindAllWithPagination(search, limit, offset)
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusInternalServerError, "Failed to fetch permissions", err.Error())
		return
	}

	paging := dtos.PagingResponseFlat[dtos.PermissionResponse]{
		Page:  page,
		Limit: limit,
		Total: total,
		Data:  permissions,
	}

	response.SuccessResponsePaging(c.Writer, "Permissions fetched successfully", paging)
}
