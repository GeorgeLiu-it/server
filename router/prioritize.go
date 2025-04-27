package router

import (
	"server/api"

	"github.com/gin-gonic/gin"
)

type PrioritizeRouter struct {
}

func (p *PrioritizeRouter) InitPrioritizeRouter(PublicRouter *gin.RouterGroup) {
	prioritizePublicRouter := PublicRouter.Group("prioritize")

	prioritizeApi := api.ApiGroupApp.PrioritizeApi
	{
		prioritizePublicRouter.GET("list", prioritizeApi.PrioritizeList)
	}
}
