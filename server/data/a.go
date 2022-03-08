/**
 * @Time: 2022/3/7 19:36
 * @Author: yt.yin
 */

package data

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"github.com/goworkeryyt/aramis/server/data/api/v1"
	"github.com/goworkeryyt/aramis/server/data/model"
	"github.com/goworkeryyt/go-core/global"
	"github.com/goworkeryyt/go-middle/record/operate"
	"go.uber.org/zap"
)

// 引入的时候自动创建表（Api信息表）
func autoCreateTables() {
	if global.DB != nil {
		// 数据库自动迁移
		err := global.DB.AutoMigrate(
			datamod.ApiInfo{},
		)
		if err != nil && global.LOG != nil {
			global.LOG.Error("初始化表时异常：", zap.Any("err", err))
		}
	}
}

// RouterRegister 注册API模块路由
func RouterRegister(rGroup *gin.RouterGroup) {
	// 初始化表
	autoCreateTables()
	api := datapi.ApiGroupApp
	// 创建路由
	apiGroup := rGroup.Group("data_api").Use(operate.OperateRecordHandler(carbon.DaysPerNormalYear))
	{
		// 创建API
		apiGroup.POST("createApi", api.CreateApi)
		// 删除API(单个)
		apiGroup.POST("deleteApi", api.DeleteApi)
		// 角色API绑定
		apiGroup.POST("roleApiBind", api.RoleApiBind)
	}
	apiGroupWithoutRecord := rGroup.Group("data_api")
	{
		// 查询API(所有)
		apiGroupWithoutRecord.GET("findApis", api.FindApis)
		// 查询API(分页)
		apiGroupWithoutRecord.GET("findApiPage", api.FindApiPage)
		// 查询API(树)
		apiGroupWithoutRecord.GET("findApiTree", api.FindApiTree)
		// 获取角色API
		apiGroupWithoutRecord.GET("findRoleApi", api.GetRoleApi)
	}
}