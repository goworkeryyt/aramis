/**
 * @Time: 2022/3/7 19:57
 * @Author: yt.yin
 */

package usersrv

import (
	"errors"
	"github.com/goworkeryyt/aramis/server/role/model"
	"github.com/goworkeryyt/go-core/global"
	"github.com/goworkeryyt/go-toolbox/array"
	"strings"
)

type UserRoleService struct{}

// UpdateUserRoles
/**
 *  @Description: 更新用户角色
 *  @receiver userRoleService
 *  @param userId
 *  @param roleIds
 *  @return err
 */
func (userRoleService *UserRoleService) UpdateUserRoles(userId string, roleIds []string) (err error) {
	userId = userRoleService.GetUserId(userId)

	// 过滤空字符串
	roleIds = array.RemoveEmptyStrInArray(roleIds)

	// 在数据库里效验角色
	var roles []rolemod.Role
	err = global.DB.Where("id in ?", roleIds).Find(&roles).Error
	if err != nil {
		return errors.New("获取角色失败")
	}
	// 效验角色数量
	if len(roles) != len(roleIds) {
		return errors.New("传入的角色存在异常")
	}

	// 删除用户原有角色
	_, err = global.CSBEF.DeleteRolesForUser(userId)
	if err != nil {
		return errors.New("删除用户角色失败")
	}

	if len(roleIds) > 0 {
		// 添加用户角色
		for i := range roleIds {
			roleId := roleIds[i]
			_, err = global.CSBEF.AddRoleForUser(userId, roleId)
			if err != nil {
				return errors.New("添加用户角色失败")
			}
		}
	}
	return
}

// GetUserRoles
/**
 *  @Description: 获取用户的角色
 *  @receiver userRoleService
 *  @param userId 用户ID
 *  @return roleIds 角色ID数组
 *  @return err
 */
func (userRoleService *UserRoleService) GetUserRoles(userId string) (roleIds []string, err error) {
	userId = userRoleService.GetUserId(userId)
	res, err := global.CSBEF.GetRolesForUser(userId)
	if err != nil {
		return nil, errors.New("获取用户具有的角色失败")
	}
	return res, err
}

// GetUserId
/**
 *  @Description: 为UserId加“user-”
 *  @receiver userRoleService
 *  @param userId
 *  @return id
 */
func (userRoleService *UserRoleService) GetUserId(userId string) (id string) {
	if strings.Contains(userId, "user-") {
		return userId
	} else {
		return "user-" + userId
	}
}

