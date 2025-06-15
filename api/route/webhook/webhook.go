package webhook

import (
	"github.com/arvinpaundra/cent/payment/application/rest"
	"github.com/gin-gonic/gin"
)

func PublicRoute(g *gin.RouterGroup, cont rest.Controller) {
	webhook := g.Group("/webhook")

	midtrans := webhook.Group("/midtrans")
	midtrans.POST("/payments", cont.MidtransUpdatePayment)
}
