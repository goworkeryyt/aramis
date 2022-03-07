/**
 * @Time: 2022/3/7 19:07
 * @Author: yt.yin
 */

package datasrv

import (
	"errors"

	"github.com/goworkeryyt/aramis/server/data/model"
	"github.com/goworkeryyt/aramis/server/data/model/response"
	"github.com/goworkeryyt/go-core/db"
	"github.com/goworkeryyt/go-core/global"
	"github.com/goworkeryyt/go-toolbox/page"
	"gorm.io/gorm"
)

type ApiInfoService struct{}

// CreateApiInfo
/**
 *  @Description: 新建API
 *  @receiver apiInfoService
 *  @param apiInfo
 *  @return err
 */
func (apiInfoService *ApiInfoService) CreateApiInfo(apiInfo datamod.ApiInfo) (err error) {
	var a datamod.ApiInfo
	// 判断是否存在相同的API
	if !errors.Is(global.DB.Where("path = ? AND method = ? AND api_group = ?", apiInfo.Path, apiInfo.Method, apiInfo.ApiGroup).First(&a).Error, gorm.ErrRecordNotFound) {
		return errors.New("该接口已存在")
	}
	err = global.DB.Create(&apiInfo).Error
	return err
}

// DeleteApiInfo
/**
 *  @Description: 删除单个API
 *  @receiver apiInfoService
 *  @param id
 *  @return err
 */
func (apiInfoService *ApiInfoService) DeleteApiInfo(id string) (err error) {
	var apiInfo datamod.ApiInfo
	// 不存在的API无需删除
	err = global.DB.Where("id = ?", id).First(&apiInfo).Error
	if err != nil {
		return nil
	}

	// 被角色绑定的API不能删除
	answer := global.CSBEF.GetFilteredPolicy(0, "", apiInfo.Path, apiInfo.Method)
	if len(answer) > 0 {
		return errors.New("被角色绑定的API不能删除")
	}

	// 删除api记录
	err = global.DB.Where("id = ?", id).Delete(&datamod.ApiInfo{}).Error
	if err != nil {
		return errors.New("删除API失败")
	}
	return err
}

// GetApiInfo
/**
 *  @Description: 获取API详情
 *  @receiver apiInfoService
 *  @param id
 *  @return apiInfo
 *  @return err
 */
func (apiInfoService *ApiInfoService) GetApiInfo(id string) (apiInfo datamod.ApiInfo, err error) {
	err = global.DB.Where("id = ?", id).First(&apiInfo).Error
	return
}

// GetAllApiInfo
/**
 *  @Description: 查询所有API
 *  @receiver apiInfoService
 *  @return apiInfos
 *  @return err
 */
func (apiInfoService *ApiInfoService) GetAllApiInfo() (apiInfos []datamod.ApiInfo, err error) {
	err = global.DB.Order("api_group").Find(&apiInfos).Error
	return
}

// GetApiInfoPage
/**
 *  @Description: API分页查询
 *  @receiver apiInfoService
 *  @param pageInfo
 *  @return err
 *  @return pageBean
 */
func (apiInfoService *ApiInfoService) GetApiInfoPage(pageInfo *page.PageInfo) (err error, pageBean *page.PageBean) {
	pageBean = &page.PageBean{Page: pageInfo.Current, PageSize: pageInfo.RowCount}
	rows := make([]*datamod.ApiInfo, 0)
	err, pageBean = db.FindPage(&datamod.ApiInfo{}, &rows, pageInfo)
	return
}

// GetApiInfoTree
/**
 *  @Description: 获取API树
 *  @receiver apiInfoService
 *  @return apiTrees
 *  @return err
 */
func (apiInfoService *ApiInfoService) GetApiInfoTree() (apiTrees []datares.ApiInfoTreeRes, err error) {
	apiTrees = make([]datares.ApiInfoTreeRes, 0)

	var apiInfos []datamod.ApiInfo
	err = global.DB.Model(&datamod.ApiInfo{}).Distinct().Pluck("api_group", &apiInfos).Error
	if err != nil {
		return nil, errors.New("查询API信息异常")
	}
	apiMap, err := apiInfoService.GetApiInfoMap()
	for _, api := range apiInfos {
		var apiTree datares.ApiInfoTreeRes
		apiTree.Description = api.ApiGroup
		apiTree.ApiInfos = apiMap[api.ApiGroup]
		apiTrees = append(apiTrees, apiTree)
	}
	return
}

// GetApiInfoMap
/**
 *  @Description: 获取API的Map
 *  @receiver apiInfoService
 *  @return apiMap
 *  @return err
 */
func (apiInfoService *ApiInfoService) GetApiInfoMap() (apiMap map[string][]datamod.ApiInfo, err error) {
	var allApiInfos []datamod.ApiInfo
	apiMap = make(map[string][]datamod.ApiInfo)
	err = global.DB.Order("api_group").Find(&allApiInfos).Error
	if err != nil {
		return nil, errors.New("查询API异常")
	}
	for _, v := range allApiInfos {
		apiMap[v.ApiGroup] = append(apiMap[v.ApiGroup], v)
	}
	return
}


