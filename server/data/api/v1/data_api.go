/**
 * @Time: 2022/3/7 19:27
 * @Author: yt.yin
 */

package datapi

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/goworkeryyt/aramis/server/data/model"
	"github.com/goworkeryyt/go-core/global"
	"github.com/goworkeryyt/go-toolbox/page"
	"github.com/goworkeryyt/go-toolbox/request"
	"github.com/goworkeryyt/go-toolbox/result"
	"github.com/goworkeryyt/go-toolbox/validator"
	"go.uber.org/zap"
)

type DataApi struct{}

// CreateApi 创建API
func (ApiInfoApi *DataApi) CreateApi(c *gin.Context) {
	// json绑定
	var apiInfo datamod.ApiInfo
	err := c.ShouldBindJSON(&apiInfo)
	if err != nil {
		result.FailMsg("绑定的Json格式有误", c)
		return
	}
	// 数据效验
	if err = validator.Verify(apiInfo, datamod.ApiVerify); err != nil {
		result.FailMsg(err.Error(), c)
		return
	}
	// 判定Method
	apiInfo.Method = strings.ToUpper(apiInfo.Method)
	if apiInfo.Method != http.MethodGet && apiInfo.Method != http.MethodPost {
		result.FailMsg("错误的请求方法", c)
		return
	}

	// 创建API
	if err = apiInfoService.CreateApiInfo(apiInfo); err != nil {
		global.LOG.Error("创建失败:", zap.Any("err", err))
		result.FailMsg("创建失败:"+err.Error(), c)
	} else {
		result.OkMsg("创建成功", c)
	}
}

// DeleteApi 删除API(单个)
func (ApiInfoApi *DataApi) DeleteApi(c *gin.Context) {
	// json绑定
	var rqId request.ID
	err := c.ShouldBindJSON(&rqId)
	if err != nil {
		result.FailMsg("绑定的Json格式有误", c)
		return
	}

	if err = apiInfoService.DeleteApiInfo(rqId.ID); err != nil {
		global.LOG.Error("删除失败:", zap.Any("err", err))
		result.FailMsg("删除失败:"+err.Error(), c)
	} else {
		result.OkMsg("删除成功", c)
	}
}

// FindApis 查询API(所有)
func (ApiInfoApi *DataApi) FindApis(c *gin.Context) {
	if apiInfos, err := apiInfoService.GetAllApiInfo(); err != nil {
		global.LOG.Error("查询失败:", zap.Any("err", err))
		result.FailMsg("查询失败:"+err.Error(), c)
	} else {
		result.OkData(apiInfos, c)
	}
}

// FindApiTree 查询API(树)
func (ApiInfoApi *DataApi) FindApiTree(c *gin.Context) {
	if apiTrees, err := apiInfoService.GetApiInfoTree(); err != nil {
		global.LOG.Error("查询失败:", zap.Any("err", err))
		result.FailMsg("查询失败:"+err.Error(), c)
	} else {
		result.OkData(apiTrees, c)
	}
}

// FindApiPage 查询API(分页)
func (ApiInfoApi *DataApi) FindApiPage(c *gin.Context) {
	pageInfo := page.PageParam(c)
	if pageInfo == nil {
		result.FailMsg("获取失败,解析请求参数异常", c)
		return
	}
	err, pageBean := apiInfoService.GetApiInfoPage(pageInfo)
	if err != nil {
		global.LOG.Error("获取失败:", zap.Any("err", err))
		result.FailMsg("获取失败:"+err.Error(), c)
	} else {
		result.OkDataMsg(pageBean, "获取成功", c)
	}
}



