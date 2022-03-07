/**
 * @Time: 2022/3/7 17:46
 * @Author: yt.yin
 */

package roleapi

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/goworkeryyt/aramis/server/role/model"
	"github.com/goworkeryyt/aramis/server/user/model"
	"github.com/goworkeryyt/go-core/global"
	"github.com/goworkeryyt/go-core/jwt"
	"github.com/goworkeryyt/go-toolbox/page"
	"github.com/goworkeryyt/go-toolbox/request"
	"github.com/goworkeryyt/go-toolbox/result"
	"github.com/goworkeryyt/go-toolbox/validator"
	"go.uber.org/zap"

)

type RoleApi struct{}

// CreateRole 新建角色
func (u *RoleApi) CreateRole(c *gin.Context) {
	// json绑定
	var role rolemod.Role
	err := c.ShouldBindJSON(&role)
	if err != nil {
		result.FailMsg("绑定的Json格式有误", c)
		return
	}

	// 数据效验
	if err = validator.Verify(role, rolemod.RoleVerify); err != nil {
		result.FailMsg(err.Error(), c)
		return
	}

	// 管理员创建全局角色，商户创建商户角色
	// 从ctx中获取claims
	claims, err := jwt.GetClaims(c)
	if err != nil {
		global.LOG.Error("解析Token失败:", zap.Any("err:", err))
		result.FailMsg("解析Token失败:"+err.Error(), c)
		return
	}

	// 身份效验
	if claims.UserType == usermod.ADMI {
		// 全局角色
		role.MerchantNo = ""
	} else if claims.UserType ==usermod.MERT {
		// 商户角色
		if strings.TrimSpace(claims.MerchantNo) != "" {
			role.MerchantNo = strings.TrimSpace(claims.MerchantNo)
		}
	} else {
		// 设备供应商用户和运营单位用户无权创建角色
		result.FailMsg("您没有创建角色的权限", c)
		return
	}

	// 创建角色
	if err = roleService.CreateRole(role); err != nil {
		global.LOG.Error("创建失败:", zap.Any("err", err))
		result.FailMsg("创建失败:"+err.Error(), c)
	} else {
		result.OkMsg("创建成功", c)
	}
}

// DeleteRole 删除角色
func (u *RoleApi) DeleteRole(c *gin.Context) {
	// json绑定
	var rqId request.ID
	err := c.ShouldBindJSON(&rqId)
	if err != nil {
		result.FailMsg("绑定的Json格式有误", c)
		return
	}
	// 删除角色
	if err = roleService.DeleteRole(rqId.ID); err != nil {
		global.LOG.Error("删除失败:", zap.Any("err", err))
		result.FailMsg("删除失败:"+err.Error(), c)
	} else {
		result.OkMsg("删除成功", c)
	}
}

// UpdateRole 更新角色
func (u *RoleApi) UpdateRole(c *gin.Context) {
	// json绑定
	var role rolemod.Role
	err := c.ShouldBindJSON(&role)
	if err != nil {
		result.FailMsg("绑定的Json格式有误", c)
		return
	}

	// 更新角色
	if err = roleService.UpdateRole(role); err != nil {
		global.LOG.Error("更新失败:", zap.Any("err", err))
		result.FailMsg("更新失败:"+err.Error(), c)
	} else {
		result.OkMsg("更新成功", c)
	}
}

// FindRole 通过ID查询角色
func (u *RoleApi) FindRole(c *gin.Context) {
	// 获取ID
	var r request.ID
	_ = c.ShouldBindQuery(&r)

	// 查询角色
	roleInfo, err := roleService.FindRoleByID(r.ID)
	if err != nil {
		global.LOG.Error("查询失败:", zap.Any("err:", err))
		result.FailMsg("查询失败:"+err.Error(), c)
	} else {
		result.OkData(roleInfo, c)
	}
}

// FindAllRoles 查询所有角色
func (u *RoleApi) FindAllRoles(c *gin.Context) {
	isCreate := c.DefaultQuery("isCreate", "0")

	var merchantNo string
	// 从ctx中获取claims
	claims, err := jwt.GetClaims(c)
	if err != nil {
		global.LOG.Error("解析Token失败:", zap.Any("err:", err))
		result.FailMsg("解析Token失败:"+err.Error(), c)
		return
	}

	// 查看范围指定
	if claims.UserType == usermod.ADMI {
		// 管理员可以查看所有角色
		// 从url参数筛选商户
		merchantNo = c.Query("merchantNo")
	} else if claims.UserType == usermod.MERT {
		// 商户用户可以查询自己创建的角色
		// 获取失败或获取到的商户号为空，则从URL中读取
		if claims.MerchantNo == "" {
			global.LOG.Error("该用户Token异常:", zap.Any("err:", "商户编号缺失"))
			result.FailMsg("该用户Token异常:商户编号缺失", c)
			return
		} else {
			merchantNo = claims.MerchantNo
		}
	}
	// 查询对应得角色
	roles, err := roleService.FindAllRoles(merchantNo,isCreate)
	if err != nil {
		global.LOG.Error("查询失败:", zap.Any("err:", err))
		result.FailMsg("查询失败:"+err.Error(), c)
	} else {
		result.OkData(roles, c)
	}
}

// GetRolePage 角色分页查询
func (u *RoleApi) GetRolePage(c *gin.Context) {
	pageInfo := page.PageParam(c)
	if pageInfo == nil {
		result.FailMsg("获取失败,解析请求参数异常", c)
		return
	}

	// 从ctx中获取claims
	claims, err := jwt.GetClaims(c)
	if err != nil {
		global.LOG.Error("解析Token失败！", zap.Any("err:", err))
		result.FailMsg("解析Token失败:"+err.Error(), c)
		return
	}

	// 商户用户只能查看自己商户创建的角色
	if claims.UserType == usermod.MERT {
		pageInfo.AndParams["merchant_no = ?"] = claims.MerchantNo
	}

	err, pageBean := roleService.GetRolePage(pageInfo)
	if err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		result.FailMsg("获取失败:"+err.Error(), c)
	} else {
		result.OkDataMsg(pageBean, "获取成功", c)
	}
}


