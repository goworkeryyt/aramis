/**
 * @Time: 2022/3/7 18:26
 * @Author: yt.yin
 */

package roleapi

import (
	"github.com/gin-gonic/gin"
	"github.com/goworkeryyt/aramis/server/user/model"
	"github.com/goworkeryyt/go-core/global"
	"github.com/goworkeryyt/go-core/jwt"
	"github.com/goworkeryyt/go-toolbox/result"
	"go.uber.org/zap"
	"strings"
)

type RoleButtonApi struct{}

// GetSelfButtons 获取角色按钮权限
func (roleButtonApi *RoleButtonApi) GetSelfButtons(c *gin.Context) {
	url := c.Query("url")
	// 从ctx中获取claims
	claims, err := jwt.GetClaims(c)
	if err != nil {
		global.LOG.Error("解析Token失败:", zap.Any("err", err))
		result.FailMsg("解析Token失败:"+err.Error(), c)
	}
	// 管理员直接获得完整按钮权限
	if claims.UserType == usermod.ADMI {
		if buttons, err := roleButtonService.GetAllButtons(url); err != nil {
			global.LOG.Error("查询失败:", zap.Any("err", err))
			result.FailMsg("查询失败:"+err.Error(), c)
			return
		} else {
			result.OkData(buttons, c)
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
	if buttons, err := roleButtonService.GetRoleButtons(roleIds, url); err != nil {
		global.LOG.Error("查询失败:", zap.Any("err", err))
		result.FailMsg("查询失败:"+err.Error(), c)
		return
	} else {
		result.OkData(buttons, c)
		return
	}
}

