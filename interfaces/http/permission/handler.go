package permission

import (
	"net/http"
	"time"
	"wells-go/application/dtos"
	"wells-go/application/usecases"
	"wells-go/response"

	"github.com/gin-gonic/gin"
)

type PermissionHandler struct {
	usecase usecases.PermissionUsecase
}

func NewPermissionController(usecase usecases.PermissionUsecase) *PermissionHandler {
	return &PermissionHandler{usecase}
}

func (h *PermissionHandler) Create(c *gin.Context) {
	startedAt := time.Now()

	var req dtos.CreatePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c.Writer, http.StatusBadRequest, "Invalid request", err.Error(), startedAt)
		return
	}
	permission, err := h.usecase.Create(req)
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusInternalServerError, "Failed to create permission", err.Error(), startedAt)
		return
	}
	response.SuccessResponse(c.Writer, "Permission created successfully", permission, startedAt)
}

func (h *PermissionHandler) Update(c *gin.Context) {
	startedAt := time.Now()

	id := c.Param("id")
	var req dtos.UpdatePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c.Writer, http.StatusBadRequest, "Invalid request", err.Error(), startedAt)
		return
	}
	permission, err := h.usecase.Update(id, req)
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusInternalServerError, "Failed to update permission", err.Error(), startedAt)
		return
	}
	response.SuccessResponse(c.Writer, "Permission updated successfully", permission, startedAt)
}

func (h *PermissionHandler) Delete(c *gin.Context) {
	startedAt := time.Now()

	id := c.Param("id")
	if err := h.usecase.Delete(id); err != nil {
		response.ErrorResponse(c.Writer, http.StatusInternalServerError, "Failed to delete permission", err.Error(), startedAt)
		return
	}
	response.SuccessResponse(c.Writer, "Permission deleted successfully", nil, startedAt)
}

func (h *PermissionHandler) FindByID(c *gin.Context) {
	startedAt := time.Now()

	id := c.Param("id")
	permission, err := h.usecase.FindByID(id)
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusNotFound, "Permission not found", err.Error(), startedAt)
		return
	}
	response.SuccessResponse(c.Writer, "Permission retrieved successfully", permission, startedAt)
}

func (h *PermissionHandler) FindAll(c *gin.Context) {
	startedAt := time.Now()

	permissions, err := h.usecase.FindAll()
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusInternalServerError, "Failed to fetch permissions", err.Error(), startedAt)
		return
	}
	response.SuccessResponse(c.Writer, "Permissions retrieved successfully", permissions, startedAt)
}
