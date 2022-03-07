/**
 * @Time: 2022/3/7 19:34
 * @Author: yt.yin
 */

package datapi

import (
	"github.com/gin-gonic/gin"
	"github.com/goworkeryyt/aramis/server/data/model"
	"github.com/goworkeryyt/aramis/server/data/model/request"
	"github.com/goworkeryyt/go-core/global"
	"github.com/goworkeryyt/go-toolbox/request"
	"github.com/goworkeryyt/go-toolbox/result"
	"github.com/goworkeryyt/go-toolbox/validator"
	"go.uber.org/zap"
)

type RoleDataApi struct{}

// RoleApiBind 角色API绑定
func (roleApiApi *RoleDataApi) RoleApiBind(c *gin.Context) {
	// json绑定
	var roleApiReq datareq.RoleApiBindReq
	err := c.ShouldBindJSON(&roleApiReq)
	if err != nil {
		result.FailMsg("绑定的Json格式有误", c)
		return
	}
	// 数据效验
	if err = validator.Verify(roleApiReq, datamod.RoleApiBindVerify); err != nil {
		result.FailMsg(err.Error(), c)
		return
	}

	// 角色API绑定
	if err = roleApiService.RoleApiBind(roleApiReq.RoleId, roleApiReq.ApiInfoIds); err != nil {
		global.LOG.Error("绑定失败:", zap.Any("err", err))
		result.FailMsg("绑定失败:"+err.Error(), c)
	} else {
		result.OkMsg("绑定成功", c)
	}
}

// GetRoleApi 获取角色API
func (roleApiApi *RoleDataApi) GetRoleApi(c *gin.Context) {
	// json绑定
	var roleId request.ID
	_ = c.ShouldBindQuery(&roleId)

	// 查询角色
	if apis, err := roleApiService.GetRoleApi(roleId.ID); err != nil {
		global.LOG.Error("查询失败:", zap.Any("err", err))
		result.FailMsg("查询失败:"+err.Error(), c)
	} else {
		result.OkData(apis, c)
	}
}

