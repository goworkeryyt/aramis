/**
 * @Time: 2022/3/5 15:22
 * @Author: yt.yin
 */

package role

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"github.com/goworkeryyt/aramis/server/role/api/v1"
	"github.com/goworkeryyt/aramis/server/role/model"
	"github.com/goworkeryyt/go-core/global"
	"github.com/goworkeryyt/go-middle/record/operate"
	"go.uber.org/zap"
)

// 引入的时候自动创建表
func autoCreateTables() {
	if global.DB != nil {
		// 数据库自动迁移
		err := global.DB.AutoMigrate(
			rolemod.Role{},
			rolemod.RoleMenu{},
			rolemod.RoleButton{},
		)
		if err != nil && global.LOG != nil {
			global.LOG.Error("初始化表时异常：", zap.Any("err", err))
		}
	}
}

// RouterRegister 注册角色模块路由
func RouterRegister(rGroup *gin.RouterGroup) {
	autoCreateTables()
	api := roleapi.ApiGroupApp
	// 角色相关接口
	roleGroup := rGroup.Group("role").Use(operate.OperateRecordHandler(carbon.DaysPerNormalYear))
	{
		// 新建角色
		roleGroup.POST("createRole", api.CreateRole)
		// 删除角色
		roleGroup.POST("deleteRole", api.DeleteRole)
		// 更新角色
		roleGroup.POST("updateRole", api.UpdateRole)
		// 角色菜单绑定
		roleGroup.POST("roleMenuBind", api.RoleMenuBind)
	}

	// 无操作记录的接口
	roleGroupWithoutRecord := rGroup.Group("role")
	{
		// 查询所有角色
		roleGroupWithoutRecord.GET("findAllRoles", api.FindAllRoles)
		// 查询单个角色
		roleGroupWithoutRecord.GET("findRole", api.FindRole)
		// 分页查询角色
		roleGroupWithoutRecord.GET("findRolePage", api.GetRolePage)
		// 获取角色菜单
		roleGroupWithoutRecord.GET("getRoleMenu", api.GetRoleMenu)
		// 获取自身菜单
		roleGroupWithoutRecord.GET("getSelfMenuTree", api.GetSelfMenuTree)
		// 获取自身按钮权限
		roleGroupWithoutRecord.GET("getSelfButtons", api.GetSelfButtons)
	}

}