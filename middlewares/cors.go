package middlewares

import "github.com/gin-gonic/gin"

func CORSAllowed() gin.HandlerFunc {
	return func(gCtx *gin.Context) {
		gCtx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		gCtx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		gCtx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		gCtx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		// OPTIONS
		if gCtx.Request.Method == "OPTIONS" {
			gCtx.AbortWithStatus(204)
			return
		}

		gCtx.Next()
	}
}
