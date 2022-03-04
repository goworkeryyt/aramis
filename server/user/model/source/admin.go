/**
 * @Time: 2022/3/4 18:08
 * @Author: yt.yin
 */

package source

import (
	"aramis/server/user/model"
	"github.com/golang-module/carbon/v2"
	"github.com/gookit/color"
	"github.com/goworkeryyt/go-core/global"
	"github.com/goworkeryyt/go-toolbox/sign"
	"github.com/goworkeryyt/go-toolbox/uuid"
	"gorm.io/gorm"
)

type admin struct{}

var Admin = new(admin)

var admins = []usermod.User{
	{
		ID:         uuid.UUID(),
		CreateTime: carbon.Now().ToDateTimeString(),
		UpdateTime: carbon.Now().ToDateTimeString(),
		Username: "admin",
		Password: sign.SHA256("123456"),
		NickName: "超级管理员",
		UserType: usermod.ADMI,
	},
}

// Init
/**
 *  @Description: sys_user表初始化
 *  @receiver a
 *  @return error
 */
func (a *admin) Init() error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("user_name IN ?", []string{"admin"}).Find(&[]usermod.User{}).RowsAffected > 0 {
			color.Danger.Println("sys_user 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&admins).Error; err != nil {
			// 遇到错误时回滚事务
			return err
		}
		color.Info.Println("sys_user 表初始数据成功!")
		return nil
	})
}