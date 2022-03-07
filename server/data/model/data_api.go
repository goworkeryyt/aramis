/**
 * @Time: 2022/3/7 18:55
 * @Author: yt.yin
 */

package datamod

import "github.com/goworkeryyt/go-toolbox/validator"

var (
	// ApiVerify Api字段效验
	ApiVerify = validator.Rules{
		"Path":        {validator.NotEmpty()}, // Api路径
		"Description": {validator.NotEmpty()}, // 中文描述
		"ApiGroup":    {validator.NotEmpty()}, // Api组
		"Method":      {validator.NotEmpty()}, // 请求方法
	}

	// RoleApiBindVerify 角色API绑定
	RoleApiBindVerify = validator.Rules{
		"ID":         {validator.NotEmpty()}, // 角色ID
		"MerchantNo": {validator.NotEmpty()}, // 商户（域）
	}
)

type ApiInfo struct {

	/** 主键id */
	ID                string        `json:"id"                         gorm:"column:id;primary_key;type:varchar(36)"`

	/** 创建时间 */
	CreateTime        string        `json:"createTime"                 gorm:"column:create_time;index;type:varchar(20)"`

	/** 修改时间 */
	UpdateTime        string        `json:"updateTime"                 gorm:"column:update_time;type:varchar(20)"`

	/** api路径 */
	Path              string        `json:"path"                       gorm:"column:path;comment:api路径;type:varchar(255);not null"`

	/** api中文描述 */
	Description       string        `json:"description"                gorm:"column:description;comment:api中文描述;type:text;not null"`

	/** api组 */
	ApiGroup          string        `json:"apiGroup"                   gorm:"column:api_group;comment:api组;type:varchar(100);not null"`

	/** 方法:创建POST(默认)|查看GET */
	Method            string        `json:"method"                     gorm:"column:method;comment:方法;type:varchar(10);default:POST;"`
}

func (a ApiInfo) TableName() string {
	return "sys_api"
}




