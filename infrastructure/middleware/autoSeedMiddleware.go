package middleware

import (
	"strings"
	"time"
	"wells-go/domain/entities"
	"wells-go/domain/repositories"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func getFriendlyNameFromPath(path, method string) string {
	parts := strings.Split(path, "/")
	name := ""
	for _, p := range parts {
		if p != "" && p != "api" && p != "v1" {
			name = strings.Title(p)
			break
		}
	}

	if strings.Contains(path, ":id") {
		if strings.ToUpper(method) == "GET" {
			name += " Get ID"
		} else if strings.ToUpper(method) == "PUT" {
			name += " Update"
		} else if strings.ToUpper(method) == "DELETE" {
			name += " Delete"
		} else {
			name += " " + strings.Title(strings.ToLower(method))
		}
	} else if strings.HasSuffix(path, "/") && strings.ToUpper(method) == "GET" {
		name += " Get All"
	} else {
		last := parts[len(parts)-1]
		if last != "" && last != name {
			name += " " + strings.Title(last)
		}
		name += " " + strings.ToUpper(method)
	}

	return name
}

func AutoSeedRouteMiddleware(repo repositories.PathRouteRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.FullPath()
		method := c.Request.Method

		if path != "" && path != "/" {
			name := getFriendlyNameFromPath(path, method)

			route := &entities.PathRouteEntities{
				Path:      path,
				Method:    method,
				Name:      name,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			if err := repo.SeedRoute(route); err != nil {
				log.Error().Err(err).Str("path", path).Str("method", method).Str("name", name).Msg("Failed to auto-seed route on request")
			}
		}

		c.Next()
	}
}

func SeedAllRoutesFromRouter(router *gin.Engine, repo repositories.PathRouteRepository) {
	routes := router.Routes()
	for _, r := range routes {
		name := getFriendlyNameFromPath(r.Path, r.Method)

		route := &entities.PathRouteEntities{
			Path:      r.Path,
			Method:    r.Method,
			Name:      name,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if err := repo.SeedRoute(route); err != nil {
			log.Error().Err(err).Str("path", r.Path).Str("method", r.Method).Str("name", name).Msg("Failed to seed route from router")
		} else {
			log.Info().Str("path", r.Path).Str("method", r.Method).Str("name", name).Msg("Route seeded from router")
		}
	}
}
