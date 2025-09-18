package dtos

type RouteAccessRequestDTO struct {
	RoutePath      string `json:"route_path" binding:"required"`
	HTTPMethod     string `json:"http_method" binding:"required"`
	RoleName       string `json:"role_name" binding:"required"`
	PermissionName string `json:"permission_name" binding:"required"`
}
