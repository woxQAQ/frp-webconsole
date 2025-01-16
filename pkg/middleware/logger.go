package middleware

import (
	"net/http"
	"time"

	"github.com/woxQAQ/frp-webconsole/pkg/log"
	"go.uber.org/zap"
)

// 添加一个响应写入器的包装器
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func Logger() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			h.ServeHTTP(w, r)
			wrapped := newResponseWriter(w)
			duration := time.Since(startTime)
			log.Logger().Info("request",
				zap.String("method", r.Method),
				zap.String("url", r.URL.String()),
				zap.Duration("duration", duration),
				zap.Int("status", wrapped.statusCode),
			)
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
