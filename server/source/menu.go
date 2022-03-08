/**
 * @Time: 2022/3/8 11:25
 * @Author: yt.yin
 */

package source

import (
	"github.com/golang-module/carbon/v2"
	"github.com/gookit/color"
	"github.com/goworkeryyt/aramis/server/menu/model"
	"github.com/goworkeryyt/go-core/global"
	"gorm.io/gorm"
)

var Menu = new(menu)

type menu struct{}

var menus = []menumod.Menu {
	{ID: "648d2467026e481383fd8a62901ba22d", CreateTime: carbon.Now().ToDateTimeString(), MenuName: "菜单管理", Sort:2, Url: "menu", RouterName: "menu", IsShow: 1, FilePath: "views/superAdmin/menu/menu.vue", MenuIcon: "", ParentId: "7f275e4796214dc48066d7714071a50b"},
	{ID: "5123afd2a3ac489f95fbc4b7865ba3d3", CreateTime: carbon.Now().ToDateTimeString(), MenuName: "用户管理", Sort:4, Url: "user", RouterName: "user", IsShow: 1, FilePath: "views/superAdmin/user/user.vue", MenuIcon: "", ParentId: "7f275e4796214dc48066d7714071a50b"},
	{ID: "e8566ed7deda4f88a6ad0a4b8974bda8", CreateTime: carbon.Now().ToDateTimeString(), MenuName: "操作历史", Sort:5, Url: "sysOperationRecord", RouterName: "sysOperationRecord", IsShow: 1, FilePath: "views/superAdmin/operation/sysOperationRecord.vue", MenuIcon: "", ParentId: "7f275e4796214dc48066d7714071a50b"},
	{ID: "8b015470a20a47eaa660808a5b793d52", CreateTime: carbon.Now().ToDateTimeString(), MenuName: "访问历史", Sort:6, Url: "accessRecord", RouterName: "accessRecord", IsShow: 1, FilePath: "views/superAdmin/operation/accessRecord.vue", MenuIcon: "", ParentId: "7f275e4796214dc48066d7714071a50b"},
	{ID: "ece2a0a3c7c3448490d5ecafb9121ad7", CreateTime: carbon.Now().ToDateTimeString(), MenuName: "角色管理", Sort:1, Url: "authority", RouterName: "authority", IsShow: 1, FilePath: "views/superAdmin/authority.vue", MenuIcon: "", ParentId: "7f275e4796214dc48066d7714071a50b"},
	{ID: "a8abac9ed1d741ac8a23bab26acef163", CreateTime: carbon.Now().ToDateTimeString(), MenuName: "api管理", Sort:3, Url: "api", RouterName: "api", IsShow: 1, FilePath: "views/superAdmin/api.vue", MenuIcon: "", ParentId: "7f275e4796214dc48066d7714071a50b"},
	{ID: "651515f100e34fbc805fa2d7019bc335", CreateTime: carbon.Now().ToDateTimeString(), MenuName: "商户管理", Sort:1, Url: "merchantList", RouterName: "merchantList", IsShow: 1, FilePath: "views/merchant/merchantList.vue", MenuIcon: "", ParentId: "a38526d5ee464957a88efd808a4a80e9"},
	{ID: "7f275e4796214dc48066d7714071a50b", CreateTime: carbon.Now().ToDateTimeString(), MenuName: "系统管理", Sort:1, Url: "superAdmin", RouterName: "superAdmin", IsShow: 1, FilePath: "views/superAdmin", MenuIcon: "s-tools", ParentId: ""},
	{ID: "a38526d5ee464957a88efd808a4a80e9", CreateTime: carbon.Now().ToDateTimeString(), MenuName: "商户信息", Sort:2, Url: "merchant", RouterName: "merchant", IsShow: 1, FilePath: "views/merchant", MenuIcon: "s-goods", ParentId: ""},
	{ID: "7e2bbef03e744b55899ba4d0ceb73163", CreateTime: carbon.Now().ToDateTimeString(), MenuName: "单位信息", Sort:3, Url: "corp", RouterName: "corp", IsShow: 1, FilePath: "views/corp", MenuIcon: "s-help", ParentId: ""},
}

// Init
/**
 *  @Description: sys_menu表初始化
 *  @receiver m
 *  @return error
 */
func (m *menu) Init() error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("menu_name IN ?", []string{"系统管理", "商户信息","单位信息"}).Find(&[]menumod.Menu{}).RowsAffected > 0 {
			color.Danger.Println("sys_menu 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&menus).Error; err != nil {
			// 遇到错误时回滚事务
			return err
		}
		color.Info.Println("sys_menu 表初始数据成功!")
		return nil
	})
}

