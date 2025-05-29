package route

import (
	"github.com/arvinpaundra/cent/payment/api/middleware"
	"github.com/arvinpaundra/cent/payment/application/resthttp"
	"github.com/arvinpaundra/cent/payment/core/validator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Routes struct {
	g    *gin.Engine
	db   *gorm.DB
	vld  *validator.Validator
	cont *resthttp.Controller
}

func NewRoutes(g *gin.Engine, db *gorm.DB, vld *validator.Validator) *Routes {
	controller := resthttp.NewController(db, vld)

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

func (r *Routes) GatherRoutes() {
	r.public()

	r.private()

	r.internal()
}

func (r *Routes) public() {
	// v1 := r.g.Group("/api/v1")
}

func (r *Routes) private() {
	// v1 := r.g.Group("/api/v1")
}

func (r *Routes) internal() {
}
