package middleware

import (
	"net/http"

	"github.com/arvinpaundra/cent/payment/core/format"
	"github.com/arvinpaundra/centpb/gen/go/auth/v1"
	"github.com/gin-gonic/gin"
)

func Authenticate(authsvc auth.AuthenticateServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")

		payload := &auth.CheckSessionRequest{
			Token: bearerToken,
		}

		res, err := authsvc.CheckSession(c, payload)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, format.Unauthorized(err.Error()))
			return
		}

		c.Set("session", res)

		c.Next()
	}
}
