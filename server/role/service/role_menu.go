/**
 * @Time: 2022/3/7 18:05
 * @Author: yt.yin
 */

package rolesrv

import (
	"errors"
	"github.com/goworkeryyt/aramis/server/menu/model"
	"github.com/goworkeryyt/aramis/server/role/model"
	"github.com/goworkeryyt/go-core/global"
	"github.com/goworkeryyt/go-toolbox/array"
)

type RoleMenuService struct{}

// RoleMenuBind
/**
 *  @Description: 角色菜单及按钮权限绑定
 *  @receiver roleMenuService
 *  @param roleId 角色ID
 *  @param menuIds 菜单ID列表
 *  @param buttonIds 按钮ID列表
 *  @return err
 */
func (roleMenuService RoleMenuService) RoleMenuBind(roleId string, menuIds []string, buttonIds []string) (err error) {
	// 过滤空字符串
	menuIds = array.RemoveEmptyStrInArray(menuIds)
	buttonIds = array.RemoveEmptyStrInArray(buttonIds)

	// 效验角色是否存在
	var role rolemod.Role
	err = global.DB.Where("id = ?", roleId).First(&role).Error
	if err != nil {
		return errors.New("该角色不存在")
	}

	// 效验按钮是否存在
	var buttons []menumod.Button
	err = global.DB.Where("id in ?", buttonIds).Find(&buttons).Error
	if err != nil {
		return errors.New("查询按钮异常")
	}
	if len(buttons) != len(buttonIds) {
		return errors.New("存在错误的按钮ID")
	}

	// 效验菜单是否存在
	var menus []menumod.Menu
	err = global.DB.Where("id in ?", menuIds).Find(&menus).Error
	if err != nil {
		return errors.New("查询菜单异常")
	}
	if len(menus) != len(menuIds) {
		return errors.New("存在错误的按钮ID")
	}

	// 开启数据库事务,使用defer，recover监听事务结束关闭事务
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 解绑原有菜单权限
	err = tx.Where("role_id = ?", role.ID).Delete(&rolemod.RoleMenu{}).Error
	if err != nil {
		tx.Rollback()
		return errors.New("解绑原有菜单权限失败")
	}
	// 解绑原有按钮权限
	err = tx.Where("role_id = ?", role.ID).Delete(&rolemod.RoleButton{}).Error
	if err != nil {
		tx.Rollback()
		return errors.New("解绑原有按钮权限失败")
	}

	var roleMenus []rolemod.RoleMenu
	for _, menu := range menus {
		roleMenu := rolemod.RoleMenu{
			RoleID: role.ID,
			MenuID: menu.ID,
		}
		roleMenus = append(roleMenus, roleMenu)
	}
	if len(roleMenus) > 0 {
		err = tx.Create(&roleMenus).Error
		if err != nil {
			tx.Rollback()
			return errors.New("菜单权限绑定失败")
		}
	}

	var roleButtons []rolemod.RoleButton
	for _, button := range buttons {
		roleButton := rolemod.RoleButton{
			RoleId:   role.ID,
			ButtonId: button.ID,
			MenuId:   button.MenuId,
		}
		roleButtons = append(roleButtons, roleButton)
	}
	if len(roleButtons) > 0 {
		err = tx.Create(&roleButtons).Error
		if err != nil {
			tx.Rollback()
			return errors.New("按钮权限绑定失败")
		}
	}

	// 提交数据库事务
	return tx.Commit().Error
}

// GetRoleMenu
/**
 *  @Description: 获取该角色拥有的菜单权限
 *  @receiver roleMenuService
 *  @param roleId 角色ID
 *  @return menuIds 菜单ID列表
 *  @return err
 */
func (roleMenuService RoleMenuService) GetRoleMenu(roleId string) (menuIds []string, err error) {
	menuIds = make([]string, 0)
	var roleMenus []rolemod.RoleMenu
	if err = global.DB.Select("menu_id").Where("role_id = ?", roleId).Find(&roleMenus).Error; err != nil {
		return nil, errors.New("获取角色菜单失败")
	}
	for _, rm := range roleMenus {
		menuIds = append(menuIds, rm.MenuID)
	}
	return
}

// GetRolesMenu
/**
 *  @Description: 获取多个角色的菜单
 *  @receiver roleMenuService
 *  @param roleIds 角色ID列表
 *  @return menuIds 菜单ID列表
 *  @return err
 */
func (roleMenuService RoleMenuService) GetRolesMenu(roleIds []string) (menuIds []string, err error) {
	var roleMenus []rolemod.RoleMenu
	if err = global.DB.Select("menu_id").Where("role_id in ?", roleIds).Find(&roleMenus).Error; err != nil {
		return nil, errors.New("获取角色菜单失败")
	}
	for _, rm := range roleMenus {
		menuIds = append(menuIds, rm.MenuID)
	}
	return
}

