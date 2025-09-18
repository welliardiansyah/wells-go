package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"wells-go/application/dtos"
	"wells-go/application/mappers"
	"wells-go/application/usecases"
	"wells-go/response"
)

type RouteAccessHandler struct {
	usecase *usecases.RouteAccessUsecase
}

func NewRouteAccessHandler(usecase *usecases.RouteAccessUsecase) *RouteAccessHandler {
	return &RouteAccessHandler{usecase: usecase}
}

func (h *RouteAccessHandler) GetAll(c *gin.Context) {
	data, err := h.usecase.GetAll()
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.SuccessResponse(c.Writer, "Get route access successfully", data)
}

func (h *RouteAccessHandler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusBadRequest, "invalid UUID format", nil)
		return
	}

	data, err := h.usecase.GetByID(id)
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusNotFound, err.Error(), nil)
		return
	}
	response.SuccessResponse(c.Writer, "Get route access by id successfully", data)
}

func (h *RouteAccessHandler) Create(c *gin.Context) {
	var req dtos.RouteAccessRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c.Writer, http.StatusBadRequest, err.Error(), nil)
		return
	}

	entity := mappers.ToRouteAccessEntity(&req)

	if err := h.usecase.Create(entity); err != nil {
		response.ErrorResponse(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.SuccessResponse(c.Writer, "Create route access successfully", entity)
}

func (h *RouteAccessHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusBadRequest, "invalid UUID format", nil)
		return
	}

	var req dtos.RouteAccessRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c.Writer, http.StatusBadRequest, err.Error(), nil)
		return
	}

	entity := mappers.ToRouteAccessEntityWithID(id, &req)

	if err := h.usecase.Update(entity); err != nil {
		response.ErrorResponse(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.SuccessResponse(c.Writer, "Update route access successfully", entity)
}

func (h *RouteAccessHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusBadRequest, "invalid UUID format", nil)
		return
	}

	if err := h.usecase.Delete(id); err != nil {
		response.ErrorResponse(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.SuccessResponse(c.Writer, "Delete route access successfully", nil)
}
