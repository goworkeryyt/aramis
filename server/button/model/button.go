/**
 * @Time: 2022/3/5 11:44
 * @Author: yt.yin
 */

package btnmod

// Button 按钮结构体
type Button struct {

	/** 主键id */
	ID                    string        `json:"id"                         gorm:"column:id;primary_key;type:varchar(36)"`

	/** 创建时间 */
	CreateTime            string        `json:"createTime"                 gorm:"column:create_time;index;type:varchar(20)"`

	/** 修改时间 */
	UpdateTime            string        `json:"updateTime"                 gorm:"column:update_time;type:varchar(20)"`

	/** 按钮名称 */
	ButtonName            string        `json:"buttonName"                 gorm:"column:button_name;comment:按钮名称;type:varchar(100);not null"`

	/** 按钮标识 */
	ButtonIdentity        string        `json:"buttonIdentity"             gorm:"column:button_identity;comment:按钮标识;type:varchar(255);not null"`

	/** 按钮完整标识 */
	ButtonFullIdentity    string        `json:"buttonFullIdentity"         gorm:"column:button_full_identity;comment:按钮完整标识;type:varchar(255);not null"`

	/** 页面名称 */
	PageName              string        `json:"pageName"                   gorm:"column:page_name;comment:页面名称;type:varchar(255)"`

	/** 页面路由 */
	PagePath              string        `json:"pagePath"                   gorm:"column:page_path;comment:页面路由;type:varchar(255)"`

	/** 菜单ID */
	MenuId                string        `json:"menuId"                     gorm:"column:menu_id;comment:菜单ID;"`

	/** 禁止操作 */
	Disable               bool          `json:"disable"                    gorm:"-"`
}

// TableName 自定义表名
func (Button) TableName() string {
	return "sys_button"
}
