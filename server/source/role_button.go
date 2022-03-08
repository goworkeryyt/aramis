/**
 * @Time: 2022/3/8 10:16
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

var RoleButton = new(roleButton)

type roleButton struct{}

var roleButtons = []rolemod.RoleButton{
	{RoleId: "90cb58fff9244754970711f2c4bd41a0", ButtonId: "061dc231534643c785790dcf306f4adf", MenuId: "651515f100e34fbc805fa2d7019bc335", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "90cb58fff9244754970711f2c4bd41a0", ButtonId: "0b49da0763564c5a8e44b0912cbae34a", MenuId: "86ab6b93e0a546978424008c736b29bd", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "90cb58fff9244754970711f2c4bd41a0", ButtonId: "331659be5b9c4d3f8185ba2a3c9083ab", MenuId: "651515f100e34fbc805fa2d7019bc335", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "90cb58fff9244754970711f2c4bd41a0", ButtonId: "40903ed28232423d99875e330b1c5f56", MenuId: "5123afd2a3ac489f95fbc4b7865ba3d3", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "90cb58fff9244754970711f2c4bd41a0", ButtonId: "4aa4b27eb6754399b2947e643aafb712", MenuId: "8b015470a20a47eaa660808a5b793d52", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "90cb58fff9244754970711f2c4bd41a0", ButtonId: "51320ea481c2463cbbb6deec97358dd9", MenuId: "648d2467026e481383fd8a62901ba22d", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "90cb58fff9244754970711f2c4bd41a0", ButtonId: "5a7f3b36b1c4470d8d5c27eac6ad0f4a", MenuId: "ece2a0a3c7c3448490d5ecafb9121ad7", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "90cb58fff9244754970711f2c4bd41a0", ButtonId: "5c5f49f0689d4216a24dbdd631e9d39f", MenuId: "e8566ed7deda4f88a6ad0a4b8974bda8", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "90cb58fff9244754970711f2c4bd41a0", ButtonId: "65446bfb9f0a4aa8951100dcae212e3d", MenuId: "ece2a0a3c7c3448490d5ecafb9121ad7", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "90cb58fff9244754970711f2c4bd41a0", ButtonId: "67a8cf20b1b34543a3a0edfd44a4d8ea", MenuId: "651515f100e34fbc805fa2d7019bc335", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "90cb58fff9244754970711f2c4bd41a0", ButtonId: "6918b1a9522c477f8e6bc96975f9f029", MenuId: "86ab6b93e0a546978424008c736b29bd", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "90cb58fff9244754970711f2c4bd41a0", ButtonId: "8a3c27cfc136408482d08428d088a54b", MenuId: "a8abac9ed1d741ac8a23bab26acef163", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "90cb58fff9244754970711f2c4bd41a0", ButtonId: "927f709cb44f4143b05d8b3d9deaa0c8", MenuId: "ece2a0a3c7c3448490d5ecafb9121ad7", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "90cb58fff9244754970711f2c4bd41a0", ButtonId: "9e1d5bc5a0d94a2991d58362dc74d75b", MenuId: "ece2a0a3c7c3448490d5ecafb9121ad7", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "90cb58fff9244754970711f2c4bd41a0", ButtonId: "b58e51ae3b324e5c85b85fbc836081d2", MenuId: "648d2467026e481383fd8a62901ba22d", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "90cb58fff9244754970711f2c4bd41a0", ButtonId: "b86d4f6d8e714e709c3c8e5942eddd73", MenuId: "5123afd2a3ac489f95fbc4b7865ba3d3", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "90cb58fff9244754970711f2c4bd41a0", ButtonId: "bb391d8b756e4993bbb29c75610f0ddf", MenuId: "5123afd2a3ac489f95fbc4b7865ba3d3", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "90cb58fff9244754970711f2c4bd41a0", ButtonId: "bb59c221d2744512b1ecc6b4074a8d63", MenuId: "a8abac9ed1d741ac8a23bab26acef163", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "90cb58fff9244754970711f2c4bd41a0", ButtonId: "be28a5e348184d0ba2d74bbbf06e4718", MenuId: "86ab6b93e0a546978424008c736b29bd", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "90cb58fff9244754970711f2c4bd41a0", ButtonId: "bfdd59c6fbb94f579fd4e40b297b3504", MenuId: "a8abac9ed1d741ac8a23bab26acef163", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "90cb58fff9244754970711f2c4bd41a0", ButtonId: "c37d95effd5d49bd9fa42b766120e86f", MenuId: "651515f100e34fbc805fa2d7019bc335", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "90cb58fff9244754970711f2c4bd41a0", ButtonId: "d3eb8a7b98ad42c9aacfb47a5b232325", MenuId: "648d2467026e481383fd8a62901ba22d", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "90cb58fff9244754970711f2c4bd41a0", ButtonId: "d4be75cb677c4128aaeb69929ccfbd10", MenuId: "648d2467026e481383fd8a62901ba22d", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "90cb58fff9244754970711f2c4bd41a0", ButtonId: "db0d7773e6604f9aae1c830df176d919", MenuId: "86ab6b93e0a546978424008c736b29bd", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "90cb58fff9244754970711f2c4bd41a0", ButtonId: "e00db52d9db742159286ee7ede5a76fd", MenuId: "648d2467026e481383fd8a62901ba22d", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "90cb58fff9244754970711f2c4bd41a0", ButtonId: "e9708e5f0a0142be96276355be4e828d", MenuId: "ece2a0a3c7c3448490d5ecafb9121ad7", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "90cb58fff9244754970711f2c4bd41a0", ButtonId: "fedf370edfe94c199d3ba086d968206f", MenuId: "5123afd2a3ac489f95fbc4b7865ba3d3", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "ab6da11b90524d48867eefb53ba89fc6", ButtonId: "061dc231534643c785790dcf306f4adf", MenuId: "651515f100e34fbc805fa2d7019bc335", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "ab6da11b90524d48867eefb53ba89fc6", ButtonId: "0b49da0763564c5a8e44b0912cbae34a", MenuId: "86ab6b93e0a546978424008c736b29bd", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "ab6da11b90524d48867eefb53ba89fc6", ButtonId: "331659be5b9c4d3f8185ba2a3c9083ab", MenuId: "651515f100e34fbc805fa2d7019bc335", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "ab6da11b90524d48867eefb53ba89fc6", ButtonId: "40903ed28232423d99875e330b1c5f56", MenuId: "5123afd2a3ac489f95fbc4b7865ba3d3", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "ab6da11b90524d48867eefb53ba89fc6", ButtonId: "4aa4b27eb6754399b2947e643aafb712", MenuId: "8b015470a20a47eaa660808a5b793d52", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "ab6da11b90524d48867eefb53ba89fc6", ButtonId: "5a7f3b36b1c4470d8d5c27eac6ad0f4a", MenuId: "ece2a0a3c7c3448490d5ecafb9121ad7", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "ab6da11b90524d48867eefb53ba89fc6", ButtonId: "5c5f49f0689d4216a24dbdd631e9d39f", MenuId: "e8566ed7deda4f88a6ad0a4b8974bda8", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "ab6da11b90524d48867eefb53ba89fc6", ButtonId: "65446bfb9f0a4aa8951100dcae212e3d", MenuId: "ece2a0a3c7c3448490d5ecafb9121ad7", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "ab6da11b90524d48867eefb53ba89fc6", ButtonId: "67a8cf20b1b34543a3a0edfd44a4d8ea", MenuId: "651515f100e34fbc805fa2d7019bc335", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "ab6da11b90524d48867eefb53ba89fc6", ButtonId: "6918b1a9522c477f8e6bc96975f9f029", MenuId: "86ab6b93e0a546978424008c736b29bd", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "ab6da11b90524d48867eefb53ba89fc6", ButtonId: "927f709cb44f4143b05d8b3d9deaa0c8", MenuId: "ece2a0a3c7c3448490d5ecafb9121ad7", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "ab6da11b90524d48867eefb53ba89fc6", ButtonId: "9e1d5bc5a0d94a2991d58362dc74d75b", MenuId: "ece2a0a3c7c3448490d5ecafb9121ad7", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "ab6da11b90524d48867eefb53ba89fc6", ButtonId: "b86d4f6d8e714e709c3c8e5942eddd73", MenuId: "5123afd2a3ac489f95fbc4b7865ba3d3", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "ab6da11b90524d48867eefb53ba89fc6", ButtonId: "bb391d8b756e4993bbb29c75610f0ddf", MenuId: "5123afd2a3ac489f95fbc4b7865ba3d3", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "ab6da11b90524d48867eefb53ba89fc6", ButtonId: "be28a5e348184d0ba2d74bbbf06e4718", MenuId: "86ab6b93e0a546978424008c736b29bd", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "ab6da11b90524d48867eefb53ba89fc6", ButtonId: "c37d95effd5d49bd9fa42b766120e86f", MenuId: "651515f100e34fbc805fa2d7019bc335", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "ab6da11b90524d48867eefb53ba89fc6", ButtonId: "db0d7773e6604f9aae1c830df176d919", MenuId: "86ab6b93e0a546978424008c736b29bd", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "ab6da11b90524d48867eefb53ba89fc6", ButtonId: "e9708e5f0a0142be96276355be4e828d", MenuId: "ece2a0a3c7c3448490d5ecafb9121ad7", CreateTime: carbon.Now().ToDateTimeString()},
	{RoleId: "ab6da11b90524d48867eefb53ba89fc6", ButtonId: "fedf370edfe94c199d3ba086d968206f", MenuId: "5123afd2a3ac489f95fbc4b7865ba3d3", CreateTime: carbon.Now().ToDateTimeString()},
}

// Init
/**
 *  @Description: sys_role_menu表初始化
 *  @receiver r
 *  @return error
 */
func (r *roleButton) Init() error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("role_id IN ?", []string{"ab6da11b90524d48867eefb53ba89fc6", "90cb58fff9244754970711f2c4bd41a0"}).Find(&[]rolemod.RoleButton{}).RowsAffected > 0 {
			color.Danger.Println("sys_role_button 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&roleButtons).Error; err != nil {
			// 遇到错误时回滚事务
			return err
		}
		color.Info.Println("sys_role_button 表初始数据成功!")
		return nil
	})
}

