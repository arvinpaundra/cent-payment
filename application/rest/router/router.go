package router

import (
	"github.com/arvinpaundra/cent/payment/application/rest/handler"
	"github.com/arvinpaundra/cent/payment/application/rest/middleware"
	"github.com/arvinpaundra/cent/payment/application/rest/router/donate"
	"github.com/arvinpaundra/cent/payment/application/rest/router/webhook"
	"github.com/arvinpaundra/cent/payment/core/grpc"
	"github.com/arvinpaundra/cent/payment/core/validator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type router struct {
	g   *gin.Engine
	db  *gorm.DB
	vld *validator.Validator
	hdl handler.Handler
}

func Register(g *gin.Engine, grpcClient *grpc.Client, db *gorm.DB, vld *validator.Validator) {
	h := handler.NewHandler(db, grpcClient, vld)

	g.Use(middleware.Cors())
	g.Use(gin.Recovery())
	g.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/metrics"},
	}))

	r := router{g, db, vld, h}

	r.public()
	r.private()
}

func (r *router) public() {
	v1 := r.g.Group("/api/v1")

	donate.PublicRoute(v1, r.hdl)
	webhook.PublicRoute(v1, r.hdl)
}

func (r *router) private() {
	// v1 := r.g.Group("/api/v1")
}
