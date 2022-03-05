/**
 * @Time: 2022/3/5 12:18
 * @Author: yt.yin
 */

package rolemod

// RoleMenu 角色菜单关系表
type RoleMenu struct {

	/** 角色ID */
	RoleID            string        `json:"roleID"                     gorm:"column:role_id;comment:角色ID;primary_key;type:varchar(36)"`

	/** 菜单项ID */
	MenuID            string        `json:"menuID"                     gorm:"column:menu_id;comment:菜单项ID;primary_key;type:varchar(36)"`

	/** 创建时间 */
	CreateTime        string        `json:"createTime"                 gorm:"column:create_time;index;type:varchar(20)"`

	/** 修改时间 */
	UpdateTime        string        `json:"updateTime"                 gorm:"column:update_time;type:varchar(20)"`
}

// TableName 自定义表名
func (RoleMenu) TableName() string {
	return "sys_role_menu"
}