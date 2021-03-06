/**
 * @Time: 2022/3/6 23:35
 * @Author: yt.yin
 */

package menu

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"github.com/goworkeryyt/aramis/server/menu/api"
	"github.com/goworkeryyt/aramis/server/menu/model"
	"github.com/goworkeryyt/go-core/global"
	"github.com/goworkeryyt/go-middle/record/operate"
	"go.uber.org/zap"
)

// 引入的时候自动创建表
func autoCreateTables() {
	if global.DB != nil {
		// 数据库自动迁移
		err := global.DB.AutoMigrate(
			menumod.Menu{},
			menumod.Button{},
		)
		if err != nil && global.LOG != nil {
			global.LOG.Error("初始化表时异常：", zap.Any("err", err))
		}
	}
}

// RouterRegister 注册菜单模块路由
func RouterRegister(rGroup *gin.RouterGroup) {
	// 初始化表
	autoCreateTables()
	menuApi := menuapi.ApiGroupApp
	menuGroup := rGroup.Group("menu").Use(operate.OperateRecordHandler(carbon.DaysPerNormalYear))
	{
		// 创建菜单
		menuGroup.POST("createMenu", menuApi.CreateMenu)
		// 删除菜单(单个)
		menuGroup.POST("deleteMenu", menuApi.DeleteMenu)
		// 更新菜单
		menuGroup.POST("updateMenu", menuApi.UpdateMenu)
	}
	menuGroupWithoutRecord := rGroup.Group("menu")
	{
		// 查询菜单(详情)
		menuGroupWithoutRecord.GET("findMenu", menuApi.FindMenu)
	}
}
