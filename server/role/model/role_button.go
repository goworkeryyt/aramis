/**
 * @Time: 2022/3/5 15:01
 * @Author: yt.yin
 */

package rolemod

type RoleButton struct {

	/** 角色ID */
	RoleId            string        `json:"roleID"                     gorm:"column:role_id;comment:角色ID;primary_key;type:varchar(36)"`

	/** 按钮ID */
	ButtonId          string        `json:"buttonId"                   gorm:"column:button_id;comment:按钮ID;primary_key;type:varchar(36)"`

	/** 菜单ID */
	MenuId            string        `json:"menuId"                     gorm:"column:menu_id;comment:菜单ID;type:varchar(36)"`

	/** 创建时间 */
	CreateTime        string        `json:"createTime"                 gorm:"column:create_time;index;type:varchar(20)"`

	/** 修改时间 */
	UpdateTime        string        `json:"updateTime"                 gorm:"column:update_time;type:varchar(20)"`
}

// TableName 自定义表名
func (RoleButton) TableName() string {
	return "sys_role_button"
}
