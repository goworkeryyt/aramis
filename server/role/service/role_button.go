/**
 * @Time: 2022/3/7 18:01
 * @Author: yt.yin
 */

package rolesrv

import (
	"errors"
	"strings"

	"github.com/goworkeryyt/aramis/server/button/model"
	"github.com/goworkeryyt/aramis/server/role/model"
	"github.com/goworkeryyt/go-core/global"
	"github.com/goworkeryyt/go-toolbox/array"
	"go.uber.org/zap"

)

type RoleButtonService struct{}

// UpdateRoleButton
/**
 *  @Description: 更新角色按钮
 *  @receiver roleButtonService
 *  @param roleId 角色ID
 *  @param buttonIds 按钮Id列表
 *  @return err
 */
func (roleButtonService *RoleButtonService) UpdateRoleButton(roleId string, buttonIds []string) (err error) {
	// 过滤空字符串
	buttonIds = array.RemoveEmptyStrInArray(buttonIds)

	// 查询角色是否存在
	var role rolemod.Role
	err = global.DB.Where("id = ?", roleId).First(&role).Error
	if err != nil {
		global.LOG.Error("查询按钮失败:", zap.Any("err:", err))
		return errors.New("查询角色失败")
	}

	// 查询按钮是否存在
	var buttons []btnmod.Button
	err = global.DB.Where("id in ?", buttonIds).Find(&buttons).Error
	if err != nil {
		global.LOG.Error("查询按钮失败:", zap.Any("err:", err))
		return errors.New("查询按钮失败")
	}
	// 开启数据库事务,使用defer，recover监听事务结束关闭事务
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 先解绑角色所有按钮
	err = tx.Where("role_id = ?", roleId).Delete(&rolemod.RoleButton{}).Error
	if err != nil {
		tx.Rollback()
		global.LOG.Error("查询按钮失败:", zap.Any("err:", err))
		return errors.New("角色按钮解绑失败")
	}

	// 绑定按钮
	var roleButtons []rolemod.RoleButton
	for _, button := range buttons {
		var roleButton rolemod.RoleButton
		roleButton.RoleId = role.ID
		roleButton.ButtonId = button.ID
		roleButtons = append(roleButtons, roleButton)
	}
	if len(roleButtons) > 0 {
		err = tx.Create(&roleButtons).Error
		if err != nil {
			tx.Rollback()
			return errors.New("角色按钮绑定失败")
		}
	}

	// 提交数据库事务
	return tx.Commit().Error
}

// GetRoleButtons
/**
 *  @Description: 查询角色按钮权限
 *  @receiver roleButtonService
 *  @param roleIds 角色ID列表
 *  @param url 当前页面的URL
 *  @return buttonFullIdentity 前端可用的按钮完整标识
 *  @return err
 */
func (roleButtonService *RoleButtonService) GetRoleButtons(roleIds []string, url string) (buttonFullIdentity map[string]string, err error) {
	var roleButtons []rolemod.RoleButton
	err = global.DB.Where("role_id in ?", roleIds).Find(&roleButtons).Error
	if err != nil {
		return nil, errors.New("查询按钮权限失败")
	}
	var buttonIds []string
	for _, button := range roleButtons {
		buttonIds = append(buttonIds, button.ButtonId)
	}

	var buttons []btnmod.Button
	if strings.TrimSpace(url) != "" {
		err = global.DB.Where("id in ?", buttonIds).Where("page_path = ?", url).Find(&buttons).Error
	} else {
		err = global.DB.Where("id in ?", buttonIds).Find(&buttons).Error
	}
	if err != nil {
		return nil, errors.New("查询按钮失败")
	}

	buttonFullIdentity = make(map[string]string, len(buttons))
	for _, button := range buttons {
		buttonFullIdentity[button.ButtonFullIdentity] = button.ButtonName
	}
	return buttonFullIdentity, nil
}

// GetAllButtons
/**
 *  @Description: 查询所有按钮权限
 *  @receiver roleButtonService
 *  @param url 当前页面的URL
 *  @return buttonFullIdentity 前端可用的按钮完整标识
 *  @return err
 */
func (roleButtonService *RoleButtonService) GetAllButtons(url string) (buttonFullIdentity map[string]string, err error) {
	var buttons []btnmod.Button
	if strings.TrimSpace(url) != "" {
		err = global.DB.Where("page_path = ?", url).Find(&buttons).Error
	} else {
		err = global.DB.Find(&buttons).Error
	}
	if err != nil {
		return nil, errors.New("查询按钮失败")
	}

	buttonFullIdentity = make(map[string]string, len(buttons))
	for _, button := range buttons {
		buttonFullIdentity[button.ButtonFullIdentity] = button.ButtonName
	}
	return buttonFullIdentity, nil
}

