/**
 * @Time: 2022/3/7 18:28
 * @Author: yt.yin
 */

package roleapi

import (
	"github.com/goworkeryyt/aramis/server/user/model"
	"github.com/goworkeryyt/go-core/jwt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/goworkeryyt/aramis/server/role/model/request"
	"github.com/goworkeryyt/go-core/global"
	"github.com/goworkeryyt/go-toolbox/request"
	"github.com/goworkeryyt/go-toolbox/result"
	"go.uber.org/zap"

)

type RoleMenuApi struct{}

// RoleMenuBind 角色菜单绑定
func (roleMenuApi *RoleMenuApi) RoleMenuBind(c *gin.Context) {
	// json绑定
	var rmReq rolereq.RoleMenuBindRequest
	err := c.ShouldBindJSON(&rmReq)
	if err != nil {
		result.FailMsg("绑定的Json格式有误", c)
		return
	}
	// 角色菜单绑定
	err = roleMenuService.RoleMenuBind(rmReq.RoleId, rmReq.MenuIds, rmReq.ButtonIds)
	if err != nil {
		global.LOG.Error("角色菜单绑定失败:", zap.Any("err", err))
		result.FailMsg("角色菜单绑定失败:"+err.Error(), c)
	} else {
		result.OkMsg("角色菜单绑定成功", c)
	}
}

// GetRoleMenu 角色菜单查询
func (roleMenuApi *RoleMenuApi) GetRoleMenu(c *gin.Context) {
	// json绑定
	var rqId request.ID
	_ = c.ShouldBindQuery(&rqId)
	menuIds, err := roleMenuService.GetRoleMenu(rqId.ID)
	menuInfo, err := menuService.GetMenusByIds(menuIds)
	if err != nil {
		global.LOG.Error("查询失败:", zap.Any("err", err))
		result.FailMsg("查询失败:"+err.Error(), c)
	} else {
		result.OkData(menuInfo, c)
	}
}

// GetSelfMenuTree 获取用户自身菜单树
func (roleMenuApi *RoleMenuApi) GetSelfMenuTree(c *gin.Context) {
	// 从ctx中获取claims
	claims, err := jwt.GetClaims(c)
	if err != nil {
		global.LOG.Error("解析Token失败:", zap.Any("err", err))
		result.FailMsg("解析Token失败:"+err.Error(), c)
	}
	// 管理员直接获得完整菜单树
	if claims.UserType == usermod.ADMI {
		var menuIds []string
		if menuInfo, err := menuService.GetMenuTree(menuIds, true); err != nil {
			global.LOG.Error("查询失败:", zap.Any("err", err))
			result.FailMsg("查询失败:"+err.Error(), c)
			return
		} else {
			result.OkData(menuInfo, c)
			return
		}
	}
	// 角色ID为空，则判定为异常
	if claims.AuthorityId == "" {
		global.LOG.Error("Token中角色ID为空:", zap.Any("err", err))
		result.FailMsg("Token中角色ID为空:"+err.Error(), c)
		return
	}

	var roleIds []string
	if strings.Contains(claims.AuthorityId, ",") {
		roleIds = strings.Split(claims.AuthorityId, ",")
	} else {
		roleIds = append(roleIds, claims.AuthorityId)
	}

	menuIds, err := roleMenuService.GetRolesMenu(roleIds)
	if err != nil {
		global.LOG.Error("获取角色菜单权限失败:", zap.Any("err", err))
		result.FailMsg("获取角色菜单权限失败:"+err.Error(), c)
		return
	}

	if menuInfo, err := menuService.GetMenuTree(menuIds, false); err != nil {
		global.LOG.Error("查询失败:", zap.Any("err", err))
		result.FailMsg("查询失败:"+err.Error(), c)
		return
	} else {
		result.OkData(menuInfo, c)
		return
	}
}


