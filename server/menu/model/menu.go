/**
 * @Time: 2022/3/5 11:53
 * @Author: yt.yin
 */

package menumod

import (
	"github.com/goworkeryyt/aramis/server/button/model"
	"github.com/goworkeryyt/go-toolbox/validator"
)

const (

	// SHOW 菜单展示
	SHOW = 1

	// HIDE 菜单隐藏
	HIDE = -1
)

var (
	// MenuCreateVerify 创建菜单参数效验
	MenuCreateVerify = validator.Rules{
		"MenuName":   {validator.NotEmpty(), validator.Le("100")},
		"Url":        {validator.NotEmpty(), validator.Le("255")},
		"MenuIcon":   {validator.Le("100")},
		"RouterName": {validator.NotEmpty(), validator.Le("255")},
		"IsShow":     {validator.NotEmpty()},
		"FilePath":   {validator.Le("255")},
	}
)

// Menu 菜单结构体
type Menu struct {

	/** 主键id */
	ID                string           `json:"id"                         gorm:"column:id;primary_key;type:varchar(36)"`

	/** 创建时间 */
	CreateTime        string           `json:"createTime"                 gorm:"column:create_time;index;type:varchar(20)"`

	/** 修改时间 */
	UpdateTime        string           `json:"updateTime"                 gorm:"column:update_time;type:varchar(20)"`

	/** 菜单名称 */
	MenuName          string           `json:"menuName"                   gorm:"column:menu_name;comment:菜单名称;type:varchar(100);not null"`

	/** 菜单排序 */
	Sort              uint             `json:"sort"                       gorm:"column:sort;comment:菜单排序;default:1"`

	/** 菜单标记/URL */
	Url               string           `json:"url"                        gorm:"column:url;comment:菜单标记/URL;type:varchar(255);unique;not null"`

	/** 路由名称 */
	RouterName        string           `json:"routerName"                 gorm:"column:router_name;comment:路由名称;type:varchar(255);"`

	/** 是否展示(1展示,-1隐藏) */
	IsShow            int              `json:"isShow"                     gorm:"column:is_show;comment:是否展示(1展示,-1隐藏);default:1"`

	/** 文件路径 */
	FilePath          string           `json:"filePath"                   gorm:"column:file_path;comment:文件路径;type:varchar(255);"`

	/** 菜单图标 */
	MenuIcon          string           `json:"menuIcon"                   gorm:"column:menu_icon;comment:菜单图标;type:varchar(100)"`

	/** 上级菜单ID */
	ParentId          string           `json:"parentId"                   gorm:"column:parent_id;comment:父菜单ID;type:varchar(36)"`

	/** 子菜单项 */
	Children          []Menu           `json:"children"                   gorm:"-"`

	/** 按钮列表(创建菜单时用) */
	Buttons           []btnmod.Button  `json:"buttons"                    gorm:"-"`

	/** 按钮ID */
	ButtonId          string           `json:"buttonId"                   gorm:"-"`

	/** 按钮名称 */
	ButtonName        string           `json:"buttonName"                 gorm:"-"`

	/** 按钮标识 */
	ButtonIdentity    string           `json:"buttonIdentity"             gorm:"-"`
}

// TableName 自定义表名
func (Menu) TableName() string {
	return "sys_menu"
}


