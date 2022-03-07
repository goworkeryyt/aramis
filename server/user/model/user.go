/**
 * @Time: 2022/3/4 12:03
 * @Author: yt.yin
 */

// Package usermod 用户模块包 user model
package usermod

import (
	"github.com/goworkeryyt/aramis/server/role/model"
	"github.com/goworkeryyt/go-toolbox/validator"
)

const (

	// ADMI 管理员用户类型
	ADMI    = "ADMI";

	// MERT 商户用户类型
	MERT    = "MERT";

	// COMM 普通用户类型
	COMM    = "COMM"
)

const (

	// ENABLE 用户账号为启用状态
	ENABLE  = "00"

	// DISABLE 用户账号为停用状态
	DISABLE = "99"
)

const (

	// PASSWORD 初始密码
	PASSWORD = "123456"
)

var (
	// LoginVerify 登录参数效验
	LoginVerify = validator.Rules{
		"Username":  {validator.NotEmpty()},
		"Password":  {validator.NotEmpty()},
		"CaptchaId": {validator.NotEmpty()},
		"Captcha":   {validator.NotEmpty()},
	}

	// ChangePasswordVerify 更改密码校验
	ChangePasswordVerify = validator.Rules{
		"Username":    {validator.NotEmpty()},
		"Password":    {validator.NotEmpty()},
		"NewPassword": {validator.NotEmpty()},
	}

	// CreateUserVerify 新增用户参数效验
	CreateUserVerify = validator.Rules{
		"Username": {validator.NotEmpty()}, // 用户名
		"UserType": {validator.NotEmpty()}, // 用户类型
		"RoleID":   {validator.NotEmpty()}, // 角色ID
	}

	// UpdateUserVerify 更新用户参数效验
	UpdateUserVerify = validator.Rules{
		"ID":       {validator.NotEmpty()}, // 用户ID
		"Username": {validator.NotEmpty()}, // 用户名
		"UserType": {validator.NotEmpty()}, // 用户类型
		"RoleID":   {validator.NotEmpty()}, // 角色ID
	}

	// UpdateUserStatusVerify 更新用户状态参数效验
	UpdateUserStatusVerify = validator.Rules{
		"UserId": {validator.NotEmpty()}, // 用户ID
		"Status": {validator.NotEmpty()}, // 用户状态
	}

	// UpdateUserRolesVerify 更新用户角色参数效验
	UpdateUserRolesVerify = validator.Rules{
		"UserId": {validator.NotEmpty()}, // 用户ID
	}
)

// User 用户信息实体
type User struct {

	/** 主键id */
	ID                string                `json:"id"                         gorm:"column:id;primary_key;type:varchar(36)"`

	/** 创建时间 */
	CreateTime        string                `json:"createTime"                 gorm:"column:create_time;index;type:varchar(20)"`

	/** 修改时间 */
	UpdateTime        string                `json:"updateTime"                 gorm:"column:update_time;type:varchar(20)"`

	/** 用户名 */
	Username          string                `json:"username"                   gorm:"column:user_name;comment:用户名;type:varchar(32);unique;index;not null;"`

	/** 用户昵称 */
	NickName          string                `json:"nickName"                   gorm:"column:nick_name;comment:用户昵称;type:varchar(20);default:系统用户;"`

	/** 用户头像 */
	HeaderImg         string                `json:"headerImg"                  gorm:"column:header_img;comment:用户头像;type:varchar(255);default:boy;"`

	/** 密码 */
	Password          string                `json:"-"                          gorm:"column:password;comment:密码;type:varchar(256);not null;"`

	/** 用户类型(管理员，商户，运营单位商户，供应商，普通用户) */
	UserType          string                `json:"userType"                   gorm:"column:user_type;comment:用户类型;type:varchar(8);default:COMM"`

	/** 电话 */
	Phone             string                `json:"phone"                      gorm:"column:phone;comment:电话;type:varchar(11);"`

	/** 邮箱 */
	Email             string                `json:"email"                      gorm:"column:email;comment:邮箱;type:varchar(64);"`

	/** 商户号 */
	MerchantNo        string                `json:"merchantNo"                 gorm:"column:merchant_no;comment:商户号;type:varchar(32);"`

	/** 商户名称 */
	MerchantName      string                `json:"merchantName"               gorm:"-"`

	/** 登录次数 */
	LoginCount        uint                  `json:"loginCount"                 gorm:"column:login_count;comment:登录次数;"`

	/** 用户状态(00:正常,99:停用) */
	Status            string                `json:"status"                     gorm:"column:status;comment:用户状态(00:正常,99:停用);type:varchar(4);default:00"`

	/** 最后一次登陆时间 */
	LoginTime         string                `json:"loginTime"                  gorm:"column:login_time;comment:上次登录时间;type:varchar(20);"`

	/** 创建者类型(管理员，商户，普通用户)(查询用户时筛选使用) */
	CreatorType       string                `json:"creatorType"                gorm:"column:creator_type;comment:创建者类型;type:varchar(8);default:ADMI"`

	/** 创建者ID（查询用户时使用） */
	CreatorId         string                `json:"creatorId"                  gorm:"column:creator_id;comment:创建者ID;type:varchar(32);"`

	/** 用户角色 */
	Roles             []rolemod.Role        `json:"roles"                      gorm:"-"`

}

// TableName 自定义表名
func (User) TableName() string {
	return "sys_user"
}

