package route

import (
	"mezink-assignment/internal/handler"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Handler handler.IHandler
	Engine  *gin.Engine
}

func Endpoints(r *Route) {
	r.Engine.GET("/record/:id", r.Handler.GetRecord)
	r.Engine.POST("/record/search", r.Handler.SearchRecords)
	r.Engine.POST("/record", r.Handler.CreateRecord)
	r.Engine.PUT("/record/:id", r.Handler.UpsertRecord)
	r.Engine.DELETE("/record/:id", r.Handler.DeleteRecord)
}
