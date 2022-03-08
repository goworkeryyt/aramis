/**
 * @Time: 2022/3/4 11:12
 * @Author: yt.yin
 */

package user

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"github.com/goworkeryyt/aramis/server/user/api/v1"
	"github.com/goworkeryyt/aramis/server/user/model"
	"github.com/goworkeryyt/go-core/global"
	"github.com/goworkeryyt/go-core/jwt"
	"github.com/goworkeryyt/go-middle/record/access"
	"github.com/goworkeryyt/go-middle/record/operate"
	"go.uber.org/zap"
)

// 引入的时候自动创建表
func autoCreateTables() {
	if global.DB != nil {
		// 数据库自动迁移
		err := global.DB.AutoMigrate(
			jwt.CustomClaims{},
			usermod.User{},
		)
		if err != nil && global.LOG != nil {
			global.LOG.Error("初始化表时异常：", zap.Any("err", err))
		}
	}
}

// RouterRegister 注册用户模块路由
func RouterRegister(rGroup *gin.RouterGroup) {
	autoCreateTables()
	api := userapi.ApiGroupApp
	// 带访问记录的api组
	userAccessGroup := rGroup.Group("user").Use(access.AccessRecordHandler(carbon.DaysPerNormalYear))
	{
		// 用户登录
		userAccessGroup.POST("login", api.Login)
	}
	// 带操作记录的api组
	userGroup := rGroup.Group("user").Use(operate.OperateRecordHandler(carbon.DaysPerNormalYear))
	{
		// 修改密码
		userGroup.POST("changePassword", api.ChangePassword)
		// 新建用户
		userGroup.POST("createUser", api.CreateUser)
		// 删除用户
		userGroup.POST("deleteUser", api.DeleteUser)
		// 更新用户
		userGroup.POST("updateUser", api.UpdateUser)
		// 重置用户密码
		userGroup.POST("updateUserPassword", api.UpdateUserPassword)
		// 更新用户状态
		userGroup.POST("updateUserStatus", api.UpdateUserStatus)
		// 更新用户角色
		userGroup.POST("updateUserRoles", api.UpdateUserRoles)
	}
	// 不带记录的api组
	userGroupWithoutRecord := rGroup.Group("user")
	{
		// 获取自身信息
		userGroupWithoutRecord.GET("getSelfUserInfo", api.GetUserInfo)
		// 生成验证码
		userGroupWithoutRecord.GET("captcha", api.Captcha)
		// 分页查询用户
		userGroupWithoutRecord.GET("getUserPage", api.GetUserPage)
	}
}
