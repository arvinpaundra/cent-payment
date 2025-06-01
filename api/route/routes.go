package route

import (
	"github.com/arvinpaundra/cent/payment/api/middleware"
	"github.com/arvinpaundra/cent/payment/api/route/donate"
	"github.com/arvinpaundra/cent/payment/application/rest"
	"github.com/arvinpaundra/cent/payment/core/validator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Routes struct {
	g    *gin.Engine
	db   *gorm.DB
	vld  *validator.Validator
	cont rest.Controller
}

func NewRoutes(g *gin.Engine, db *gorm.DB, vld *validator.Validator) *Routes {
	controller := rest.NewController(db, vld)

	g.Use(middleware.Cors())
	g.Use(gin.Recovery())
	g.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/metrics"},
	}))

	return &Routes{
		g:    g,
		db:   db,
		vld:  vld,
		cont: controller,
	}
}

func (r *Routes) WithPublic() *Routes {
	v1 := r.g.Group("/api/v1")

	donate.PublicRoute(v1, r.cont)

	return r
}

func (r *Routes) WithPrivate() *Routes {
	// v1 := r.g.Group("/api/v1")

	return r
}

func (r *Routes) WithInternal() *Routes {
	return r
}
