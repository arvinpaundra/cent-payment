package webhook

import (
	"github.com/arvinpaundra/cent/payment/application/rest/handler"
	"github.com/gin-gonic/gin"
)

func PublicRoute(g *gin.RouterGroup, h handler.Handler) {
	webhook := g.Group("/webhook")

	midtrans := webhook.Group("/midtrans")
	midtrans.POST("/payments", h.MidtransUpdatePayment)
}
