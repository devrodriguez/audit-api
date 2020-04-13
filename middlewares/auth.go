package middlewares

import (
	"net/http"

	"github.com/devrodriguez/audit-api/controllers"
	"github.com/devrodriguez/audit-api/models"
	"github.com/gin-gonic/gin"
)

func ValidateAuth() gin.HandlerFunc {
	var resModel models.Response

	return func(gCtx *gin.Context) {
		err := controllers.VerifyToken(gCtx.Request)

		if err != nil {
			resModel.Message = "You need to be authorized"
			resModel.Error = err.Error()

			gCtx.JSON(http.StatusUnauthorized, resModel)
			gCtx.Abort()
			return
		}

		gCtx.Next()
	}
}
