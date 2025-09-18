package pathRoute

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"wells-go/application/usecases"
	"wells-go/response"
)

type PathRouteHandler struct {
	usecase usecases.PathRouteUsecase
}

func NewPathRouteHandler(usecase usecases.PathRouteUsecase) *PathRouteHandler {
	return &PathRouteHandler{
		usecase: usecase,
	}
}

func (h *PathRouteHandler) GetAllRoutes(c *gin.Context) {
	name := c.Query("name")
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

	routes, total, err := h.usecase.GetAllRoutes(name, limit, offset)
	if err != nil {
		response.ErrorResponse(c.Writer, http.StatusInternalServerError, "Failed to fetch routes", err.Error())
		return
	}

	response.SuccessResponse(c.Writer, "Routes fetched successfully", gin.H{
		"data":       routes,
		"page":       page,
		"limit":      limit,
		"total":      total,
		"totalPages": (total + int64(limit) - 1) / int64(limit),
	})
}
