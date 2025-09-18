package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"wells-go/domain/repositories"
	"wells-go/response"
	"wells-go/util/security"
)

func RoleAndPermissionMiddlewareDynamic(repo repositories.RouteAccessRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		routePath := c.FullPath()
		method := c.Request.Method

		rawRoles, rolesExist := c.Get("roles")
		rawPermissions, permsExist := c.Get("permissions")
		if !rolesExist || !permsExist {
			response.ErrorResponse(c.Writer, http.StatusUnauthorized, "roles or permissions not found in context", nil)
			c.Abort()
			return
		}

		roles := rawRoles.([]string)
		permissions := rawPermissions.([]security.Permission)

		accessList, err := repo.GetAccessByRoute(method, routePath)
		if err != nil {
			response.ErrorResponse(c.Writer, http.StatusInternalServerError, err.Error(), nil)
			c.Abort()
			return
		}

		allowed := false

		for _, access := range accessList {
			for _, role := range roles {
				if strings.EqualFold(access.RoleName, role) {
					for _, perm := range permissions {
						if perm.Name == role {
							switch strings.ToLower(access.PermissionName) {
							case "create":
								if perm.CanCreate {
									allowed = true
								}
							case "read":
								if perm.CanRead {
									allowed = true
								}
							case "update":
								if perm.CanUpdate {
									allowed = true
								}
							case "delete":
								if perm.CanDelete {
									allowed = true
								}
							case "export":
								if perm.CanExport {
									allowed = true
								}
							case "import":
								if perm.CanImport {
									allowed = true
								}
							case "view":
								if perm.CanView {
									allowed = true
								}
							}
							if allowed {
								break
							}
						}
					}
				}
				if allowed {
					break
				}
			}
			if allowed {
				break
			}
		}

		if !allowed {
			response.ErrorResponse(c.Writer, http.StatusForbidden, "access denied", nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
