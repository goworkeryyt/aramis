/**
 * @Time: 2022/3/7 19:17
 * @Author: yt.yin
 */

package datasrv

import (
	"errors"
	"github.com/goworkeryyt/aramis/server/data/model"
	"github.com/goworkeryyt/aramis/server/role/model"
	"github.com/goworkeryyt/go-core/global"
	"github.com/goworkeryyt/go-toolbox/array"
	"strings"
)

type RoleApiService struct{}

// RoleApiBind
/**
 *  @Description: 角色API绑定
 *  @receiver roleApiService
 *  @param roleId
 *  @param apiIds
 *  @return err
 */
func (roleApiService *RoleApiService) RoleApiBind(roleId string, apiIds []string) (err error) {
	// 过滤空字符串
	apiIds = array.RemoveEmptyStrInArray(apiIds)

	// 效验角色是否合法
	var role rolemod.Role
	err = global.DB.Where("id = ?", roleId).First(&role).Error
	if err != nil || strings.TrimSpace(role.ID) == "" {
		return errors.New("绑定的角色不合法")
	}

	// 效验绑定的API是否合法
	var apiInfos []datamod.ApiInfo
	if len(apiIds) > 0 {
		err = global.DB.Where("id in ?", apiIds).Find(&apiInfos).Error
		if err != nil {
			return
		}
		if len(apiInfos) != len(apiIds) {
			return errors.New("绑定的API存在非法数据")
		}
		global.DB.Where("id in ?", apiIds).Distinct("method", "path").Find(&apiInfos)
	}

	// 批量删除API
	_, err = global.CSBEF.DeletePermissionsForUser(role.ID)
	if err != nil {
		return errors.New("批量删除API权限（策略）失败")
	}

	// 存在API才需要批量添加
	if len(apiInfos) > 0 {
		// 构建批量添加需要的策略
		var rules [][]string
		for _, info := range apiInfos {
			// 角色,Url,请求方法
			rule := []string{role.ID, info.Path, info.Method}
			rules = append(rules, rule)
		}
		_, err = global.CSBEF.AddPolicies(rules)
		if err != nil {
			return errors.New("批量添加API权限（策略）失败")
		}
	}
	return
}

// GetRoleApi
/**
 *  @Description: 角色API获取
 *  @receiver roleApiService
 *  @param roleId
 *  @return apiInfos
 *  @return err
 */
func (roleApiService *RoleApiService) GetRoleApi(roleId string) (apiInfos []datamod.ApiInfo, err error) {
	apiInfos = make([]datamod.ApiInfo, 0)

	var role rolemod.Role
	err = global.DB.Where("id = ?", roleId).First(&role).Error
	if err != nil {
		return nil, errors.New("该角色不存在")
	}

	// 获取该角色具有的权限
	res := global.CSBEF.GetPermissionsForUser(role.ID)
	pathArray := make([]string, len(res))
	for i, r := range res {
		pathArray[i] = r[1]
	}

	// 查询API对象
	err = global.DB.Where("path in ?", pathArray).Find(&apiInfos).Error
	if err != nil {
		return nil, errors.New("查询api记录失败")
	}
	return
}

