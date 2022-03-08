/**
 * @Time: 2022/3/8 10:25
 * @Author: yt.yin
 */

package source

import (
	"github.com/golang-module/carbon/v2"
	"github.com/gookit/color"
	"github.com/goworkeryyt/aramis/server/role/model"
	"github.com/goworkeryyt/go-core/global"
	"gorm.io/gorm"
)

var RoleMenu = new(roleMenu)

type roleMenu struct{}

var roleMenus = []rolemod.RoleMenu{
	{RoleID: "90cb58fff9244754970711f2c4bd41a0", MenuID: "5123afd2a3ac489f95fbc4b7865ba3d3", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleID: "90cb58fff9244754970711f2c4bd41a0", MenuID: "648d2467026e481383fd8a62901ba22d", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleID: "90cb58fff9244754970711f2c4bd41a0", MenuID: "651515f100e34fbc805fa2d7019bc335", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleID: "90cb58fff9244754970711f2c4bd41a0", MenuID: "7e2bbef03e744b55899ba4d0ceb73163", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleID: "90cb58fff9244754970711f2c4bd41a0", MenuID: "7f275e4796214dc48066d7714071a50b", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleID: "90cb58fff9244754970711f2c4bd41a0", MenuID: "86ab6b93e0a546978424008c736b29bd", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleID: "90cb58fff9244754970711f2c4bd41a0", MenuID: "8b015470a20a47eaa660808a5b793d52", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleID: "90cb58fff9244754970711f2c4bd41a0", MenuID: "a38526d5ee464957a88efd808a4a80e9", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleID: "90cb58fff9244754970711f2c4bd41a0", MenuID: "a8abac9ed1d741ac8a23bab26acef163", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleID: "90cb58fff9244754970711f2c4bd41a0", MenuID: "e8566ed7deda4f88a6ad0a4b8974bda8", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleID: "90cb58fff9244754970711f2c4bd41a0", MenuID: "ece2a0a3c7c3448490d5ecafb9121ad7", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleID: "ab6da11b90524d48867eefb53ba89fc6", MenuID: "5123afd2a3ac489f95fbc4b7865ba3d3", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleID: "ab6da11b90524d48867eefb53ba89fc6", MenuID: "651515f100e34fbc805fa2d7019bc335", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleID: "ab6da11b90524d48867eefb53ba89fc6", MenuID: "7e2bbef03e744b55899ba4d0ceb73163", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleID: "ab6da11b90524d48867eefb53ba89fc6", MenuID: "7f275e4796214dc48066d7714071a50b", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleID: "ab6da11b90524d48867eefb53ba89fc6", MenuID: "86ab6b93e0a546978424008c736b29bd", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleID: "ab6da11b90524d48867eefb53ba89fc6", MenuID: "8b015470a20a47eaa660808a5b793d52", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleID: "ab6da11b90524d48867eefb53ba89fc6", MenuID: "a38526d5ee464957a88efd808a4a80e9", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleID: "ab6da11b90524d48867eefb53ba89fc6", MenuID: "e8566ed7deda4f88a6ad0a4b8974bda8", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleID: "ab6da11b90524d48867eefb53ba89fc6", MenuID: "ece2a0a3c7c3448490d5ecafb9121ad7", CreateTime: carbon.Now().ToDateTimeString()},
}

// Init
/**
 *  @Description: sys_role_menu表初始化
 *  @receiver r
 *  @return error
 */
func (r *roleMenu) Init() error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("role_id IN ?", []string{"ab6da11b90524d48867eefb53ba89fc6", "90cb58fff9244754970711f2c4bd41a0"}).Find(&[]rolemod.RoleMenu{}).RowsAffected > 0 {
			color.Danger.Println("sys_role_menu 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&roleMenus).Error; err != nil {
			// 遇到错误时回滚事务
			return err
		}
		color.Info.Println("sys_role_menu 表初始数据成功!")
		return nil
	})
}

