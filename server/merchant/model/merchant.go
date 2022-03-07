/**
 * @Time: 2022/3/5 15:32
 * @Author: yt.yin
 */

package mchtmod

import "github.com/goworkeryyt/go-toolbox/validator"

const (

	// ENABLE 商户为启用状态
	ENABLE  = "00"

	// DISABLE 商户为停用状态
	DISABLE = "99"
)

var (

	// MerchantCreateVerify 创建商户参数效验
	MerchantCreateVerify = validator.Rules{
		"MerchantNo":   {validator.NotEmpty()},
		"MerchantName": {validator.NotEmpty()},
	}

	// MerchantUpdateVerify 更新商户参数效验
	MerchantUpdateVerify = validator.Rules{
		"ID": {validator.NotEmpty()},
	}
)

// Merchant 商户结构体
type Merchant struct {

	/** 主键id */
	ID                string        `json:"id"                         gorm:"column:id;primary_key;type:varchar(36)"`

	/** 创建时间 */
	CreateTime        string        `json:"createTime"                 gorm:"column:create_time;index;type:varchar(20)"`

	/** 修改时间 */
	UpdateTime        string        `json:"updateTime"                 gorm:"column:update_time;type:varchar(20)"`

	/** 商户编号 */
	MerchantNo        string        `json:"merchantNo"                 gorm:"column:merchant_no;comment:商户编号;type:varchar(32);"`

	/** 商户名称 */
	MerchantName      string        `json:"merchantName"               gorm:"column:merchant_name;comment:商户名称;type:varchar(100);"`

	/** 地址 */
	Address           string        `json:"address"                    gorm:"column:address;comment:地址;type:varchar(255);"`

	/** 商户等级 */
	Level             uint          `json:"level"                      gorm:"column:level;comment:商户等级;default:1"`

	/** 区域代码 */
	RegionCode        string        `json:"regionCode"                 gorm:"column:region_code;comment:区域代码;type:varchar(32);"`

	/** 区域名称 */
	RegionName        string        `json:"regionName"                 gorm:"column:region_name;comment:区域名称;type:varchar(32);"`

	/** 状态(00:正常,99:停用) */
	Status            string        `json:"status"                     gorm:"column:status;comment:状态(00:正常,99:停用);type:varchar(4);default:00"`

	/** 父商户ID */
	ParentId          string        `json:"parentId"                   gorm:"column:parent_id;comment:父商户ID;type:varchar(32)"`

	/** 子商户列表 */
	ChildMerchants  []Merchant      `json:"childMerchants"             gorm:"-"`

	/** 父商户名称 */
	ParentName        string        `json:"parentName"                 gorm:"-"`

	/** 商户密钥 */
	SelfPriKey        string        `json:"selfPriKey"                 gorm:"column:self_pri_key;comment:商户密钥;type:text"`

}

// TableName Merchant 表名
func (Merchant) TableName() string {
	return "sys_merchant"
}