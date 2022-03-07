/**
 * @Time: 2022/3/7 16:39
 * @Author: yt.yin
 */

package rolesrv

import (
	"errors"
	"strings"

	"github.com/goworkeryyt/aramis/server/merchant/model"
	"github.com/goworkeryyt/aramis/server/role/model"
	"github.com/goworkeryyt/go-core/db"
	"github.com/goworkeryyt/go-core/global"
	"github.com/goworkeryyt/go-toolbox/page"
	"gorm.io/gorm"
)

type RoleService struct{}

// CreateRole
/**
 *  @Description: 新建角色
 *  @receiver roleService
 *  @param role role实例
 *  @return err 异常
 */
func (roleService RoleService) CreateRole(role rolemod.Role) (err error) {
	// 角色名称重名效验
	var r rolemod.Role
	if !errors.Is(global.DB.Where("role_name = ? AND merchant_no = ?", role.RoleName, role.MerchantNo).First(&r).Error, gorm.ErrRecordNotFound) {
		return errors.New("角色名称被占用")
	}
	err = global.DB.Create(&role).Error
	return err
}

// DeleteRole
/**
 *  @Description: 删除角色
 *  @receiver roleService
 *  @param id 角色ID
 *  @param operatorId 操作员ID
 *  @return err
 */
func (roleService RoleService) DeleteRole(id string) (err error) {
	var role rolemod.Role
	// 不存在则不需要删除
	err = global.DB.Where("id = ?", id).First(&role).Error
	if err != nil {
		return nil
	}
	// 与用户绑定的角色不可删除
	res, err := global.CSBEF.GetUsersForRole(role.ID)
	if err != nil {
		return errors.New("查询用户角色绑定关系失败")
	}
	if len(res) > 0 {
		return errors.New("绑定用户的角色不可删除")
	}

	// 与菜单绑定的角色不可删除
	var roleMenus []rolemod.RoleMenu
	err = global.DB.Where("role_id = ?", id).Find(&roleMenus).Error
	if err != nil {
		return errors.New("查询角色菜单绑定关系失败")
	}
	if len(roleMenus) > 0 {
		return errors.New("绑定菜单的角色不可删除")
	}

	// 与按钮绑定的角色不可删除
	var roleButtons []rolemod.RoleButton
	err = global.DB.Where("role_id").Find(&roleButtons).Error
	if err != nil {
		return errors.New("查询角色按钮绑定关系失败")
	}
	if len(roleMenus) > 0 {
		return errors.New("绑定按钮的角色不可删除")
	}

	// 与API绑定的角色不可删除
	answer := global.CSBEF.GetFilteredPolicy(0, role.ID, "", "")
	if len(answer) > 0 {
		return errors.New("与API绑定的角色不可删除")
	}

	// 删除角色
	err = global.DB.Where("id = ?", id).Delete(&rolemod.Role{}).Error
	if err != nil {
		return errors.New("删除角色失败")
	}
	return
}

// UpdateRole
/**
 *  @Description: 更新角色
 *  @receiver roleService
 *  @param role
 *  @return err
 */
func (roleService RoleService) UpdateRole(role rolemod.Role) (err error) {
	// 判定角色是否存在
	var roleFind rolemod.Role
	err = global.DB.Where("id = ?", role.ID).First(&roleFind).Error
	if err != nil {
		return errors.New("该角色不存在")
	}

	// 判定角色是否重名
	var r rolemod.Role
	global.DB.Where("role_name = ? AND merchant_no = ?", role.RoleName, role.MerchantNo).First(&r)
	if r.ID != "" && r.ID != role.ID {
		return errors.New("该角色名称被占用")
	}

	// 修改角色
	roleFind.RoleName = role.RoleName
	roleFind.Remake = role.Remake
	err = global.DB.Updates(&roleFind).Error
	return err
}

// FindRoleByID
/**
 *  @Description: 通过ID查询角色
 *  @receiver roleService
 *  @param roleId 角色ID
 *  @return role
 *  @return err
 */
func (roleService RoleService) FindRoleByID(roleId string) (role rolemod.Role, err error) {
	err = global.DB.Where("id = ?", roleId).First(&role).Error
	if strings.TrimSpace(role.MerchantNo) != "" {
		// 为角色填充商户名称
		var merchant mchtmod.Merchant
		err = global.DB.Where("merchant_no = ?", role.MerchantNo).First(&merchant).Error
		if err != nil {
			return rolemod.Role{}, errors.New("绑定商户名称失败")
		}
		role.MerchantName = merchant.MerchantName
	}
	return
}

// FindAllRoles
/**
 *  @Description: 查询所有角色
 *  @receiver roleService
 *  @param merchantNo 商户编号
 *  @param isCreate 是否创建用户时使用
 *  @return roles
 *  @return err
 */
func (roleService RoleService) FindAllRoles(merchantNo, isCreate string) (roles []*rolemod.Role, err error) {
	roles = make([]*rolemod.Role, 0)
	// 查询角色
	if merchantNo != "" {
		err = global.DB.Order("create_time desc").Where("merchant_no = ?", merchantNo).Find(&roles).Error
	} else {
		if isCreate != "0" {
			err = global.DB.Order("create_time desc").Where("merchant_no = ?", "").Find(&roles).Error
		} else {
			err = global.DB.Order("create_time desc").Find(&roles).Error
		}
	}
	// 为角色填充商户名称
	merchantNos := make([]string, 0)
	for _, role := range roles {
		merchantNos = append(merchantNos, role.MerchantNo)
	}
	merchants := make([]mchtmod.Merchant, 0)
	err = global.DB.Where("merchant_no in ?", merchantNos).Find(&merchants).Error
	if err != nil {
		return nil, errors.New("查询商户失败")
	}

	if len(merchants) > 0 {
		merchantMap := make(map[string]string)
		for _, merchant := range merchants {
			merchantMap[merchant.MerchantNo] = merchant.MerchantName
		}
		for _, role := range roles {
			role.MerchantName = merchantMap[role.MerchantNo]
		}
	}
	return
}

// GetRolePage
/**
 *  @Description: 分页查询角色列表
 *  @receiver roleService
 *  @param pageInfo
 *  @return err
 *  @return pageBean
 */
func (roleService RoleService) GetRolePage(pageInfo *page.PageInfo) (err error, pageBean *page.PageBean) {
	pageBean = &page.PageBean{Page: pageInfo.Current, PageSize: pageInfo.RowCount}
	rows := make([]*rolemod.Role, 0)
	err, pageBean = db.FindPage(&rolemod.Role{}, &rows, pageInfo)

	// 为角色填充商户名称
	merchantNos := make([]string, 0)
	for _, row := range rows {
		merchantNos = append(merchantNos, row.MerchantNo)
	}
	merchants := make([]mchtmod.Merchant, 0)
	err = global.DB.Where("merchant_no in ?", merchantNos).Find(&merchants).Error
	if err != nil {
		return errors.New("查询商户失败"), pageBean
	}

	if len(merchants) > 0 {
		merchantMap := make(map[string]string)
		for _, merchant := range merchants {
			merchantMap[merchant.MerchantNo] = merchant.MerchantName
		}
		roles := make([]*rolemod.Role, len(rows))
		for i, row := range rows {
			row.MerchantName = merchantMap[row.MerchantNo]
			roles[i] = row
		}
		rows = roles
	}
	return
}

// FindRoleByIds
/**
 *  @Description: 根据角色ID列表查询角色
 *  @receiver roleService
 *  @param roleIds
 *  @return roles
 *  @return err
 */
func (roleService RoleService) FindRoleByIds(roleIds []string) (roles []rolemod.Role, err error) {
	err = global.DB.Where("id in ?", roleIds).Find(&roles).Error
	return
}



