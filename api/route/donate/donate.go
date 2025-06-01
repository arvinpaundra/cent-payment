package donate

import (
	"github.com/arvinpaundra/cent/payment/application/resthttp"
	"github.com/gin-gonic/gin"
)

func PublicRoute(g *gin.RouterGroup, cont resthttp.Controller) {
	donate := g.Group("/donate")

	donate.POST("", cont.CreateDonation)
}
