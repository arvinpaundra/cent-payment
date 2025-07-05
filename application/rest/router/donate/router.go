package donate

import (
	"github.com/arvinpaundra/cent/payment/application/rest/handler"
	"github.com/gin-gonic/gin"
)

func PublicRoute(g *gin.RouterGroup, h handler.Handler) {
	donate := g.Group("/donate/:slug")

	donate.POST("", h.CreateDonation)
}
