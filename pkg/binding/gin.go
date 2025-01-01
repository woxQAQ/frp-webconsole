package binding

import (
	"github.com/gin-gonic/gin"
)

func BindBody[T any](ctx *gin.Context) (T, error) {
	var v T
	err := ctx.ShouldBind(&v)
	return v, err
}
