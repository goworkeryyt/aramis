/**
 * @Time: 2022/3/4 12:18
 * @Author: yt.yin
 */

// Package usersrv 处理用户模块业务包 user service
package usersrv

import (
	"errors"
	"strings"

	"github.com/goworkeryyt/aramis/server/merchant/model"
	"github.com/goworkeryyt/aramis/server/role/model"
	"github.com/goworkeryyt/aramis/server/user/model"
	"github.com/goworkeryyt/aramis/server/user/model/request"
	"github.com/goworkeryyt/go-core/db"
	"github.com/goworkeryyt/go-core/global"
	"github.com/goworkeryyt/go-toolbox/array"
	"github.com/goworkeryyt/go-toolbox/page"
	"github.com/goworkeryyt/go-toolbox/sign"
	"github.com/goworkeryyt/go-toolbox/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type UserService struct{}


// Login
/**
 *  @Description: 登录验证
 *  @receiver userService
 *  @param u user对象
 *  @return err 异常
 *  @return userInter user对象
 */
func (userService *UserService) Login(u *usermod.User) (err error, userInter *usermod.User) {
	var user usermod.User
	u.Password = sign.SHA256(u.Password)
	err = global.DB.Where("user_name = ? AND password = ?", u.Username, u.Password).First(&user).Error
	// 用户状态判定
	if user.Status != usermod.ENABLE {
		return errors.New("该用户已停用"), &user
	}
	return err, &user
}

// ChangePassword
/**
 *  @Description: 修改密码
 *  @receiver userService
 *  @param u user对象
 *  @param newPassword 新的密码
 *  @return err 异常
 *  @return userInter 修改后的user
 */
func (userService *UserService) ChangePassword(u *usermod.User, newPassword string) (err error, userInter *usermod.User) {
	var user usermod.User
	u.Password = sign.SHA256(u.Password)
	err = global.DB.Where("user_name = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", sign.SHA256(newPassword)).Error
	return err, u
}

// GetUserInfo
/**
 *  @Description: 查询用户信息
 *  @receiver userService
 *  @param id
 *  @return user
 *  @return err
 */
func (userService *UserService) GetUserInfo(id string) (user usermod.User, err error) {
	err = global.DB.First(&user, "id = ?", id).Error
	// 为用户关联角色
	userRoleIds, _ := ServiceGroupApp.UserRoleService.GetUserRoles(user.ID)
	roles, errRole := roleService.FindRoleByIds(userRoleIds)
	if errRole != nil {
		return user, errors.New("用户级联角色失败")
	}
	user.Roles = roles
	return
}

// CreateUser
/**
 *  @Description: 创建用户
 *  @receiver userService
 *  @param user
 *  @return err
 */
func (userService *UserService) CreateUser(user usereq.CreateUserRequest) (err error) {
	// 开启数据库事务,使用defer，recover监听事务结束关闭事务
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 过滤空字符串
	user.RoleIds = array.RemoveEmptyStrInArray(user.RoleIds)
	var u usermod.User
	// 判断是否存在相同用户名
	if !errors.Is(global.DB.Where("user_name = ?", user.Username).First(&u).Error, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return errors.New("用户名已被占用")
	}
	// 新建用户实例
	var userNew usermod.User
	err = copier.Copy(&userNew, &user)
	if err != nil{
		tx.Rollback()
		return err
	}
	userNew.ID = uuid.UUID()
	userNew.Password = sign.SHA256(usermod.PASSWORD)
	userNew.Status = usermod.ENABLE
	// 不同用户类型进行不同处理
	if user.UserType == usermod.ADMI {
		// 超级管理员
	} else if user.UserType == usermod.MERT {
		var merchant mchtmod.Merchant
		err = tx.Where("merchant_no = ?", user.MerchantNo).Where("status = ?", usermod.ENABLE).First(&merchant).Error
		if err != nil || strings.TrimSpace(merchant.ID) == "" {
			tx.Rollback()
			return errors.New("该商户不存在或已停用")
		}

	} else if user.UserType == usermod.COMM {
		// 普通用户
	} else {
		tx.Rollback()
		return errors.New("非法的用户类型")
	}

	err = tx.Create(&userNew).Error
	if err != nil {
		tx.Rollback()
		return errors.New("创建用户失败")
	}

	if len(user.RoleIds) > 0 {
		// 效验角色
		var roles []rolemod.Role
		err = tx.Where("id in ?", user.RoleIds).Find(&roles).Error
		if err != nil {
			tx.Rollback()
			return errors.New("绑定的角色存在异常")
		}
		if len(roles) != len(user.RoleIds) {
			tx.Rollback()
			return errors.New("绑定的角色存在异常")
		}

		// 为用户批量添加角色
		for i := range user.RoleIds {
			roleId := user.RoleIds[i]
			_, err = global.CSBEF.AddRoleForUser("user-"+userNew.ID, roleId)
			if err != nil {
				tx.Rollback()
				return errors.New("为用户添加角色失败")
			}
		}
	}

	// 提交数据库事务
	return tx.Commit().Error
}

// GetUserPage
/**
 *  @Description: 分页查询用户
 *  @receiver userService
 *  @param pageInfo
 *  @return err
 *  @return pageBean
 */
func (userService *UserService) GetUserPage(pageInfo *page.PageInfo) (err error, pageBean *page.PageBean) {
	pageBean = &page.PageBean{Page: pageInfo.Current, PageSize: pageInfo.RowCount}
	rows := make([]*usermod.User, 0)
	err, pageBean = db.FindPage(&usermod.User{}, &rows, pageInfo)

	// 为用户关联角色
	for i, _ := range rows {
		user := *rows[i]
		userRoleIds, _ := ServiceGroupApp.UserRoleService.GetUserRoles(user.ID)
		roles, errRole := roleService.FindRoleByIds(userRoleIds)
		if errRole != nil {
			return errors.New("级联用户角色失败"), nil
		}
		user.Roles = roles
		*rows[i] = user
	}
	return
}

// DeleteUser
/**
 *  @Description: 删除用户
 *  @receiver userService
 *  @param id
 *  @return err
 */
func (userService *UserService) DeleteUser(id string) (err error) {
	var user usermod.User
	// 查询用户，获取对应的ID和商户（域），以便解绑角色
	err = global.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil
	}

	// 登录过的用户无法删除
	if user.LoginCount > 0 {
		return errors.New("登录过的用户无法删除")
	}

	if user.LoginCount <= 0 {
		roles, err := ServiceGroupApp.GetUserRoles(id)
		if err != nil {
			return errors.New("查询用户角色失败")
		}
		if roles != nil {
			err = ServiceGroupApp.UserRoleService.UpdateUserRoles(id, nil)
			if err != nil {
				return errors.New("解绑用户角色失败")
			}
		}

		err = global.DB.Where("id = ?", id).Delete(&usermod.User{}).Error
		if err != nil {
			return err
		}
		return nil
	}

	// 绑定过角色的用户无法删除
	userRoleIds, _ := ServiceGroupApp.UserRoleService.GetUserRoles(user.ID)
	if len(userRoleIds) > 0 {
		return errors.New("被角色绑定的用户无法删除")
	}

	// 删除用户
	err = global.DB.Where("id = ?", id).Delete(&usermod.User{}).Error
	if err != nil {
		return err
	}
	return err
}

// UpdateUserStatus
/**
 *  @Description: 更新用户状态
 *  @receiver userService
 *  @param userId
 *  @param status
 *  @return err
 */
func (userService *UserService) UpdateUserStatus(userId, status string) (err error) {
	// 判定用户是否存在
	var user usermod.User
	err = global.DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return errors.New("该用户不存在")
	}

	// 判定用户状态
	if status != usermod.ENABLE && status != usermod.DISABLE {
		return errors.New("非法的用户状态")
	}

	// 若商户存在则判定商户类型
	if user.MerchantNo != "" {
		// 查询用户所在商户状态
		var merchant mchtmod.Merchant
		err = global.DB.Where("merchant_no = ?", user.MerchantNo).First(&merchant).Error
		if err != nil {
			return errors.New("商户不存在")
		}
		// 停用的商户不能修改用户状态
		if merchant.Status == mchtmod.DISABLE {
			return errors.New("商户已停用，无法修改用户状态")
		}
	}

	// 更新商户状态
	err = global.DB.Model(&usermod.User{}).Where("id = ?", userId).Update("status", status).Error
	return
}

// UpdateUserPassword
/**
 *  @Description: 重置用户密码
 *  @receiver userService
 *  @param id
 *  @return err
 */
func (userService *UserService) UpdateUserPassword(id string) (err error) {
	err = global.DB.Model(&usermod.User{}).Where("id = ?", id).Update("password", sign.SHA256("bsit123456")).Error
	return
}

// UpdateUserRoles
/**
 *  @Description: 更新用户角色
 *  @receiver userService
 *  @param userId
 *  @param roleIds
 *  @return err
 */
func (userService *UserService) UpdateUserRoles(userId string, roleIds []string) (err error) {
	// 过滤空字符串
	roleIds = array.RemoveEmptyStrInArray(roleIds)

	// 判断用户是否存在
	var user usermod.User
	err = global.DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return errors.New("该用户不存在")
	}

	var roles []rolemod.Role
	err = global.DB.Where("id in ?", roleIds).Find(&roles).Error
	if err != nil {
		return errors.New("查询绑定的角色异常")
	}
	if len(roles) != len(roleIds) {
		return errors.New("绑定的角色存在异常")
	}

	// 更新用户角色
	err = ServiceGroupApp.UserRoleService.UpdateUserRoles(user.ID, roleIds)
	if err != nil {
		return errors.New("绑定部分角色失败")
	}
	return
}

// UpdateUser
/**
 *  @Description: 更新用户
 *  @receiver userService
 *  @param user
 *  @return err
 */
func (userService *UserService) UpdateUser(user usereq.UpdateUserRequest) (err error) {
	// 开启数据库事务,使用defer，recover监听事务结束关闭事务
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 过滤空字符串
	user.RoleIds = array.RemoveEmptyStrInArray(user.RoleIds)

	// 查询需要修改的用户
	var userNow usermod.User
	err = tx.Where("id = ?", user.ID).First(&userNow).Error
	if err != nil {
		tx.Rollback()
		return errors.New("该用户存在")
	}
	userNow.NickName = user.NickName
	userNow.HeaderImg = user.HeaderImg
	userNow.Phone = user.Phone
	userNow.Email = user.Email

	// 更新用户
	err = tx.Updates(&userNow).Error
	if err != nil {
		tx.Rollback()
		return errors.New("更新用户失败")
	}

	if len(user.RoleIds) > 0 {
		var roles []rolemod.Role
		err = tx.Where("id in ?", user.RoleIds).Find(&roles).Error
		if len(roles) != len(user.RoleIds) {
			tx.Rollback()
			return errors.New("存在非法的角色ID")
		}
		// 更新用户角色
		err = ServiceGroupApp.UserRoleService.UpdateUserRoles(userNow.ID, user.RoleIds)
		if err != nil {
			tx.Rollback()
			return errors.New("更新用户角色失败")
		}
	}

	// 提交数据库事务
	return tx.Commit().Error
}
