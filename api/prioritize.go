package api

import (
	"server/global"
	"server/model/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PrioritizeApi struct {
}

// Get the work items of SR and snow list
func (PrioritizeApi *PrioritizeApi) PrioritizeList(c *gin.Context) {
	list, err := prioritizeService.PrioritizeList()
	if err != nil {
		global.Log.Error("Failed to get the work items from Sieble and SNOW:", zap.Error(err))
		response.FailWithMessage("Failed to get the work items from Sieble and SNOW:", c)
		return
	}
	response.OkWithData(list, c)
}
