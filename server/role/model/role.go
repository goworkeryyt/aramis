/**
 * @Time: 2022/3/5 11:32
 * @Author: yt.yin
 */

package rolemod

import "github.com/goworkeryyt/go-toolbox/validator"

var (
	// RoleVerify 创建角色效验
	RoleVerify = validator.Rules{
		"RoleName": {validator.NotEmpty(), validator.Le("32")}, // 角色名称
		"Remake":   {validator.Le("255")},                      // 角色描述
	}
)

// Role 角色结构体
type Role struct {

	/** 主键id */
	ID                string        `json:"id"                         gorm:"column:id;primary_key;type:varchar(36)"`

	/** 创建时间 */
	CreateTime        string        `json:"createTime"                 gorm:"column:create_time;index;type:varchar(20)"`

	/** 修改时间 */
	UpdateTime        string        `json:"updateTime"                 gorm:"column:update_time;type:varchar(20)"`

	/** 角色名称 */
	RoleName          string        `json:"roleName"                   gorm:"column:role_name;comment:角色名称;type:varchar(32);index;not null;"`

	/** 角色描述 */
	Remake            string        `json:"remark"                     gorm:"column:remark;comment:角色描述;type:varchar(255);"`

	/** 角色所属商户 */
	MerchantNo        string        `json:"merchantNo"                 gorm:"column:merchant_no;comment:所属商户;type:varchar(32);"`

	/** 商户名称 */
	MerchantName      string        `json:"merchantName"               gorm:"-"`
}

// TableName 自定义表名
func (Role) TableName() string {
	return "sys_role"
}


