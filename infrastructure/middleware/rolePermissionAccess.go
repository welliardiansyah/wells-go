package middleware

import (
	"net/http"
	"strings"
	"time"
	"wells-go/response"
	"wells-go/util/security"

	"github.com/gin-gonic/gin"
)

func RoleAndPermissionMiddleware(allowedRoles []string, allowedActions []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		startedAt := time.Now()

		roleI, exists := c.Get("roles")
		if !exists {
			response.ErrorResponse(c.Writer, http.StatusForbidden, "Roles not found", nil, startedAt)
			c.Abort()
			return
		}

		userRoles, ok := roleI.([]string)
		if !ok {
			response.ErrorResponse(c.Writer, http.StatusForbidden, "Invalid roles type", nil, startedAt)
			c.Abort()
			return
		}

		hasRole := false
		for _, r := range userRoles {
			if stringInSliceCI(r, allowedRoles) {
				hasRole = true
				break
			}
		}
		if !hasRole {
			response.ErrorResponse(c.Writer, http.StatusForbidden, "Role not allowed", nil, startedAt)
			c.Abort()
			return
		}

		permsI, exists := c.Get("permissions")
		if !exists {
			response.ErrorResponse(c.Writer, http.StatusForbidden, "Permissions not found", nil, startedAt)
			c.Abort()
			return
		}

		perms, ok := permsI.([]security.Permission)
		if !ok {
			response.ErrorResponse(c.Writer, http.StatusForbidden, "Invalid permissions type", nil, startedAt)
			c.Abort()
			return
		}

		hasPermission := false
		for _, perm := range perms {
			for _, action := range allowedActions {
				switch strings.ToLower(action) {
				case "create":
					if perm.CanCreate {
						hasPermission = true
					}
				case "read":
					if perm.CanRead {
						hasPermission = true
					}
				case "update":
					if perm.CanUpdate {
						hasPermission = true
					}
				case "delete":
					if perm.CanDelete {
						hasPermission = true
					}
				case "export":
					if perm.CanExport {
						hasPermission = true
					}
				case "import":
					if perm.CanImport {
						hasPermission = true
					}
				case "view":
					if perm.CanView {
						hasPermission = true
					}
				}
				if hasPermission {
					break
				}
			}
			if hasPermission {
				break
			}
		}

		if !hasPermission {
			response.ErrorResponse(c.Writer, http.StatusForbidden, "Access denied for this action", nil, startedAt)
			c.Abort()
			return
		}

		c.Next()
	}
}

func stringInSliceCI(s string, list []string) bool {
	s = strings.ToLower(s)
	for _, a := range list {
		if s == strings.ToLower(a) {
			return true
		}
	}
	return false
}
