package dtos

type GetByNameRequest struct {
	RoleName string `json:"role_name" binding:"required"`
}

type RouteAccessResponse struct {
	ID             string `json:"id"`
	RoutePath      string `json:"route_path"`
	HTTPMethod     string `json:"http_method"`
	RoleName       string `json:"role_name"`
	PermissionName string `json:"permission_name"`
}
