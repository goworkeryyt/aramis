/**
 * @Time: 2022/3/8 10:11
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

var Role = new(role)

type role struct{}

var roles = []rolemod.Role{
	{
		ID:         "90cb58fff9244754970711f2c4bd41a0",
		CreateTime: carbon.Now().ToDateTimeString(),
		RoleName:   "超管角色",
		Remake:     "系统自动生成",
	},
	{
		ID:         "ab6da11b90524d48867eefb53ba89fc6",
		CreateTime: carbon.Now().ToDateTimeString(),
		RoleName:   "商户角色",
		Remake:     "系统自动生成",
	},
}

// Init
/**
 *  @Description: sys_menu表初始化
 *  @receiver m
 *  @return error
 */
func (r *role) Init() error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("role_name IN ?", []string{"超管角色", "商户角色"}).Find(&[]rolemod.Role{}).RowsAffected > 0 {
			color.Danger.Println("sys_role 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&roles).Error; err != nil {
			// 遇到错误时回滚事务
			return err
		}
		color.Info.Println("sys_role 表初始数据成功!")
		return nil
	})
}

