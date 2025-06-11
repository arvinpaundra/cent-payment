package donate

import (
	"github.com/arvinpaundra/cent/payment/application/rest"
	"github.com/gin-gonic/gin"
)

func PublicRoute(g *gin.RouterGroup, cont rest.Controller) {
	donate := g.Group("/donate/:slug")

	donate.POST("", cont.CreateDonation)
}
