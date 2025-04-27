package api

import (
	"server/config"
	"server/global"
	"server/model/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ConfigApi struct {
}

// GetWebsite 获取网站配置
func (configApi *ConfigApi) GetWebsite(c *gin.Context) {
	response.OkWithData(global.Config.Website, c)
}

// UpdateWebsite 更新网站配置
func (configApi *ConfigApi) UpdateWebsite(c *gin.Context) {
	var req config.Website
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = configService.UpdateWebsite(req)
	if err != nil {
		global.Log.Error("Failed to update website:", zap.Error(err))
		response.FailWithMessage("Failed to update website", c)
		return
	}
	response.OkWithMessage("Successfully updated website", c)
}

// GetSystem 获取系统配置
func (configApi *ConfigApi) GetSystem(c *gin.Context) {
	response.OkWithData(global.Config.System, c)
}

// UpdateSystem 更新系统配置
func (configApi *ConfigApi) UpdateSystem(c *gin.Context) {
	var req config.System
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = configService.UpdateSystem(req)
	if err != nil {
		global.Log.Error("Failed to update system:", zap.Error(err))
		response.FailWithMessage("Failed to update system", c)
		return
	}
	response.OkWithMessage("Successfully updated system", c)
}

// GetEmail 获取邮箱配置
func (configApi *ConfigApi) GetEmail(c *gin.Context) {
	response.OkWithData(global.Config.Email, c)
}

// UpdateEmail 更新邮箱配置
func (configApi *ConfigApi) UpdateEmail(c *gin.Context) {
	var req config.Email
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = configService.UpdateEmail(req)
	if err != nil {
		global.Log.Error("Failed to update email:", zap.Error(err))
		response.FailWithMessage("Failed to update email", c)
		return
	}
	response.OkWithMessage("Successfully updated email", c)
}

// GetJwt 获取Jwt配置
func (configApi *ConfigApi) GetJwt(c *gin.Context) {
	response.OkWithData(global.Config.Jwt, c)
}

// UpdateJwt 更新Jwt配置
func (configApi *ConfigApi) UpdateJwt(c *gin.Context) {
	var req config.Jwt
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = configService.UpdateJwt(req)
	if err != nil {
		global.Log.Error("Failed to update jwt:", zap.Error(err))
		response.FailWithMessage("Failed to update jwt", c)
		return
	}
	response.OkWithMessage("Successfully updated jwt", c)
}
