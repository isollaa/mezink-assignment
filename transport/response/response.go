package response

import (
	"mezink-assignment/shared/base"

	"github.com/gin-gonic/gin"
)

func JSON(c *gin.Context, base base.Base) {
	obj := map[string]any{
		"code": base.Code,
		"msg":  base.Msg,
	}

	for _, v := range base.Data {
		obj[v.Field] = v.Data
	}

	c.JSON(base.HTTPCode, obj)
}
