/**
 * @Time: 2022/3/6 00:37
 * @Author: yt.yin
 */

package mchtapi

import (
	"github.com/gin-gonic/gin"
	"github.com/goworkeryyt/aramis/server/merchant/model"
	"github.com/goworkeryyt/aramis/server/merchant/model/request"
	"github.com/goworkeryyt/aramis/server/user/model"
	"github.com/goworkeryyt/go-core/global"
	"github.com/goworkeryyt/go-core/jwt"
	"github.com/goworkeryyt/go-toolbox/page"
	"github.com/goworkeryyt/go-toolbox/request"
	"github.com/goworkeryyt/go-toolbox/result"
	"github.com/goworkeryyt/go-toolbox/validator"
	"go.uber.org/zap"
)

type MerchantApi struct{}

// CreateMerchant 创建商户
func (merchantAPI *MerchantApi) CreateMerchant(c *gin.Context) {
	// json绑定
	var merchant mchtmod.Merchant
	var merchantParentNo string
	err := c.ShouldBindJSON(&merchant)
	if err != nil {
		result.FailMsg("绑定的Json格式有误", c)
		return
	}
	// 数据效验
	if err = validator.Verify(merchant, mchtmod.MerchantCreateVerify); err != nil {
		result.FailMsg(err.Error(), c)
		return
	}
	// 从ctx中获取claims
	claims, err := jwt.GetClaims(c)
	if err != nil {
		global.LOG.Error("解析Token失败！", zap.Any("err:", err))
		result.FailMsg("解析Token失败:"+err.Error(), c)
		return
	}
	if claims.UserType == usermod.ADMI {
		// 超管只能创建一级商户
		merchant.Level = 1
	} else if claims.UserType == usermod.MERT {
		// 一级商户可以创建二级商户
		merchant.Level = 2
		merchantParentNo = claims.MerchantNo
	} else {
		result.FailMsg("您没有创建商户的权限", c)
		return
	}
	// 创建商户
	if err = merchantService.CreateMerchant(merchant, merchantParentNo); err != nil {
		global.LOG.Error("创建失败:", zap.Any("err", err))
		result.FailMsg("创建失败:"+err.Error(), c)
	} else {
		result.OkMsg("创建成功", c)
	}
}

// UpdateMerchant 更新商户
func (merchantAPI *MerchantApi) UpdateMerchant(c *gin.Context) {
	// json绑定
	var merchant mchtreq.UpdateMerchantReq
	err := c.ShouldBindJSON(&merchant)
	if err != nil {
		result.FailMsg("绑定的Json格式有误", c)
		return
	}
	// 数据效验
	if err = validator.Verify(merchant, mchtmod.MerchantUpdateVerify); err != nil {
		result.FailMsg(err.Error(), c)
		return
	}
	// 更新商户
	if err = merchantService.UpdateMerchant(merchant); err != nil {
		global.LOG.Error("更新失败:", zap.Any("err", err))
		result.FailMsg("更新失败:"+err.Error(), c)
	} else {
		result.OkMsg("更新成功", c)
	}
}

// DeleteMerchant 删除商户
func (merchantAPI *MerchantApi) DeleteMerchant(c *gin.Context) {
	// json绑定
	var rqId request.ID
	err := c.ShouldBindJSON(&rqId)
	if err != nil {
		result.FailMsg("绑定的Json格式有误", c)
		return
	}
	if err = merchantService.DeleteMerchant(rqId.ID); err != nil {
		global.LOG.Error("删除失败:", zap.Any("err", err))
		result.FailMsg("删除失败"+err.Error(), c)
	} else {
		result.OkMsg("删除成功", c)
	}
}

// GetMerchantPage 分页查询商户
func (merchantAPI *MerchantApi) GetMerchantPage(c *gin.Context) {
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
		merchant, err := merchantService.FindMerchantByMerchantNo(claims.MerchantNo)
		if err != nil {
			result.FailMsg(claims.MerchantNo+"商户异常", c)
			return
		}
		// 商户只能只看当前商户的用户
		pageInfo.AndParams["parent_id = ?"] = merchant.ID
	}
	err, pageBean := merchantService.GetMerchantPage(pageInfo)
	if err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		result.FailMsg("获取失败"+err.Error(), c)
	} else {
		result.OkDataMsg(pageBean, "获取成功", c)
	}
}

// GetAllMerchants 查询所有商户
func (merchantAPI *MerchantApi) GetAllMerchants(c *gin.Context) {
	var parentId string

	// 从ctx中获取claims
	claims, errToken := jwt.GetClaims(c)
	if errToken != nil || claims == nil {
		result.FailMsg("解析token失败", c)
		return
	}
	if claims.UserType == usermod.MERT {
		merchant, err := merchantService.FindMerchantByMerchantNo(claims.MerchantNo)
		if err != nil {
			result.FailMsg(claims.MerchantNo+"商户异常", c)
			return
		}
		parentId = merchant.ID
	} else {
		result.FailMsg("您的身份无权查询商户", c)
		return
	}

	if menuInfos, err := merchantService.GetAllMerchants(parentId); err != nil {
		global.LOG.Error("查询失败:", zap.Any("err", err))
		result.FailMsg("查询失败:"+err.Error(), c)
	} else {
		result.OkData(menuInfos, c)
	}
}

