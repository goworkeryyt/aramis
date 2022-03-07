/**
 * @Time: 2022/3/7 20:20
 * @Author: yt.yin
 */

package userapi

import (
	"github.com/gin-gonic/gin"
	"github.com/goworkeryyt/aramis/server/user/model"
	"github.com/goworkeryyt/aramis/server/user/model/request"
	"github.com/goworkeryyt/go-core/global"
	"github.com/goworkeryyt/go-core/jwt"
	"github.com/goworkeryyt/go-toolbox/page"
	"github.com/goworkeryyt/go-toolbox/request"
	"github.com/goworkeryyt/go-toolbox/result"
	"github.com/goworkeryyt/go-toolbox/validator"
	"go.uber.org/zap"
)

type UserApi struct{}

// CreateUser 新增用户
func (u *UserApi) CreateUser(c *gin.Context) {
	// json绑定
	var user usereq.CreateUserRequest
	err := c.ShouldBindJSON(&user)
	if err != nil {
		result.FailMsg("绑定的Json格式有误", c)
		return
	}
	// 数据效验
	if err = validator.Verify(user, usermod.CreateUserVerify); err != nil {
		result.FailMsg(err.Error(), c)
		return
	}

	// 从ctx中获取claims
	claims, errToken := jwt.GetClaims(c)
	if errToken != nil || claims == nil {
		result.FailMsg("解析token失败", c)
		return
	}

	// 新创建的用户类型效验
	if claims.UserType == usermod.ADMI {
		// 管理员可以创建任何类型用户
	} else if claims.UserType == usermod.COMM {
		// 普通用户只能创建商户用户
		if user.UserType != usermod.MERT {
			result.FailMsg("普通用户只能创建商户类型的用户", c)
			return
		}
	} else if claims.UserType == usermod.MERT {
		user.MerchantNo = claims.MerchantNo
	} else {
		result.FailMsg("您没有创建用户的权限", c)
		return
	}
	// 绑定创建者类型及ID
	user.CreatorType = claims.UserType
	user.CreatorId = claims.UserId

	// 创建用户
	if err = userService.CreateUser(user); err != nil {
		global.LOG.Error("创建失败:", zap.Any("err", err))
		result.FailMsg("创建失败:"+err.Error(), c)
	} else {
		result.OkMsg("创建成功", c)
	}
}

// UpdateUser 更新用户
func (u *UserApi) UpdateUser(c *gin.Context) {
	// json绑定
	var user usereq.UpdateUserRequest
	err := c.ShouldBindJSON(&user)
	if err != nil {
		result.FailMsg("绑定的Json格式有误", c)
		return
	}
	// 数据效验
	if err = validator.Verify(user, usermod.UpdateUserVerify); err != nil {
		result.FailMsg(err.Error(), c)
		return
	}

	// 更新用户
	if err = userService.UpdateUser(user); err != nil {
		global.LOG.Error("更新失败:", zap.Any("err", err))
		result.FailMsg("更新失败:"+err.Error(), c)
	} else {
		result.OkMsg("更新成功", c)
	}
}

// UpdateUserPassword 重置用户密码
func (u *UserApi) UpdateUserPassword(c *gin.Context) {
	// json绑定
	var rqId request.ID
	err := c.ShouldBindJSON(&rqId)
	if err != nil {
		result.FailMsg("绑定的Json格式有误", c)
		return
	}
	if err = userService.UpdateUserPassword(rqId.ID); err != nil {
		global.LOG.Error("重置密码失败:", zap.Any("err", err))
		result.FailMsg("重置密码失败:"+err.Error(), c)
	} else {
		result.OkMsg("重置密码成功", c)
	}
}

// UpdateUserStatus 修改用户状态
func (u *UserApi) UpdateUserStatus(c *gin.Context) {
	// json绑定
	var user usereq.UpdateUserStatusReq
	err := c.ShouldBindJSON(&user)
	if err != nil {
		result.FailMsg("绑定的Json格式有误", c)
		return
	}
	// 数据效验
	if err = validator.Verify(user, usermod.UpdateUserStatusVerify); err != nil {
		result.FailMsg(err.Error(), c)
		return
	}

	if user.Status != usermod.ENABLE && user.Status != usermod.DISABLE {
		result.FailMsg("非法的用户状态", c)
		return
	}

	if err = userService.UpdateUserStatus(user.UserId, user.Status); err != nil {
		global.LOG.Error("更新失败:", zap.Any("err", err))
		result.FailMsg("更新失败:"+err.Error(), c)
	} else {
		result.OkMsg("更新成功", c)
	}
}

// DeleteUser 删除用户
func (u *UserApi) DeleteUser(c *gin.Context) {
	// json绑定
	var rqId request.ID
	err := c.ShouldBindJSON(&rqId)
	if err != nil {
		result.FailMsg("绑定的Json格式有误", c)
		return
	}
	if err = userService.DeleteUser(rqId.ID); err != nil {
		global.LOG.Error("删除失败:", zap.Any("err", err))
		result.FailMsg("删除失败:"+err.Error(), c)
	} else {
		result.OkMsg("删除成功", c)
	}
}

// GetUserPage 分页查询用户
func (u *UserApi) GetUserPage(c *gin.Context) {
	pageInfo := page.PageParam(c)
	if pageInfo == nil {
		result.FailMsg("获取失败,解析请求参数异常", c)
		return
	}

	// 从ctx中获取claims
	claims, errToken := jwt.GetClaims(c)
	if errToken != nil || claims == nil {
		result.FailMsg("解析token失败", c)
		return
	}

	if claims.UserType == usermod.ADMI {
		// 管理员可以查看任何用户
	} else if claims.UserType == usermod.MERT {
		// 商户只能只看当前商户的用户
		pageInfo.AndParams["merchant_no = ?"] = claims.MerchantNo
		pageInfo.AndParams["creator_type = ?"] = usermod.MERT
	} else if claims.UserType == usermod.COMM {
		// 普通用户只能查看自己创建的账号
		pageInfo.AndParams["creator_type = ?"] = usermod.COMM
		pageInfo.AndParams["creator_id = ?"] = claims.UserId
	}
	// 不能看见自己的账号
	pageInfo.AndParams["id != ?"] = claims.UserId

	err, pageBean := userService.GetUserPage(pageInfo)
	if err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		result.FailMsg("获取失败:"+err.Error(), c)
	} else {
		result.OkDataMsg(pageBean, "获取成功", c)
	}
}

// UpdateUserRoles 更新用户角色
func (u *UserApi) UpdateUserRoles(c *gin.Context) {
	// json绑定
	var upRole usereq.UpdateUserRolesReq
	err := c.ShouldBindJSON(&upRole)
	if err != nil {
		result.FailMsg("绑定的Json格式有误", c)
		return
	}
	// 数据效验
	if err = validator.Verify(upRole, usermod.UpdateUserRolesVerify); err != nil {
		result.FailMsg(err.Error(), c)
		return
	}
	// 更新用户角色
	if err = userService.UpdateUserRoles(upRole.UserId, upRole.RoleIds); err != nil {
		global.LOG.Error("更新用户角色失败:", zap.Any("err", err))
		result.FailMsg("更新用户角色失败:"+err.Error(), c)
	} else {
		result.OkMsg("更新用户角色成功", c)
	}
}



