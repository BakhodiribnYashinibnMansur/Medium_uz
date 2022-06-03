package cors

import "github.com/gin-gonic/gin"

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept,  Cache-Control, X-Requested-With, multipart-form-data, multipart/form-data ,X-Requested-With, X-Request-ID, X-HTTP-Method-Override, Upload-Length, Upload-Offset, Tus-Resumable, Upload-Concat, User-Agent, Referrer, Origin,  Location,  Tus-Version,  Tus-Max-Size, Tus-Extension, Upload-Metadata, Upload-Defer-Length, ")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST,  GET, PUT , DELETE ,PATCH, HEAD, OPTIONS")
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}
		ctx.Next()
	}
}
