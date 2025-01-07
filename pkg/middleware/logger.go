package middleware

import (
	"net/http"

	"github.com/woxQAQ/frp-webconsole/pkg/log"
	"goa.design/goa/v3/middleware"
)

func Logger() func(http.Handler) http.Handler {
	logger := middleware.NewLogger(log.NewStdLogger())
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Log(logger, r, w, h)
		})
	}
}

// func Recovery(log *zap.Logger) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		defer func() {
// 			if err := recover(); err != nil {
// 				var brokenPipe bool
// 				if ne, ok := err.(*net.OpError); ok {
// 					if se, ok := ne.Err.(*os.SyscallError); ok {
// 						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
// 							brokenPipe = true
// 						}
// 					}
// 				}
// 				httpRequest, _ := httputil.DumpRequest(c.Request, false)
// 				if brokenPipe {
// 					log.Error("gin panic",
// 						zap.Any("error", err),
// 						zap.String("request", string(httpRequest)),
// 					)
// 					c.Error(err.(error))
// 					c.Abort()
// 					return
// 				}
// 				log.Error("gin panic",
// 					zap.Any("error", err),
// 					zap.String("request", string(httpRequest)),
// 					zap.String("stack", string(debug.Stack())),
// 				)
// 				c.AbortWithError(http.StatusInternalServerError, err.(error))
// 			}
// 		}()
// 		c.Next()
// 	}
// }
