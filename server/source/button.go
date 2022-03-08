/**
 * @Time: 2022/3/8 11:17
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

var Button = new(button)

type button struct{}

var buttons = []menumod.Button{
	{ID: "061dc231534643c785790dcf306f4adf", CreateTime: carbon.Now().ToDateTimeString(), ButtonName: "删除", ButtonIdentity: "delete", ButtonFullIdentity: "merchant:merchantList:delete", PageName: "商户管理", PagePath: "merchant/merchantList", MenuId: "651515f100e34fbc805fa2d7019bc335"},
	{ID: "331659be5b9c4d3f8185ba2a3c9083ab", CreateTime: carbon.Now().ToDateTimeString(), ButtonName: "编辑", ButtonIdentity: "edit", ButtonFullIdentity: "merchant:merchantList:edit", PageName: "商户管理", PagePath: "merchant/merchantList", MenuId: "651515f100e34fbc805fa2d7019bc335"},
	{ID: "40903ed28232423d99875e330b1c5f56", CreateTime: carbon.Now().ToDateTimeString(), ButtonName: "编辑", ButtonIdentity: "edit", ButtonFullIdentity: "superAdmin:user:edit", PageName: "用户管理", PagePath: "superAdmin/user", MenuId: "5123afd2a3ac489f95fbc4b7865ba3d3"},
	{ID: "4aa4b27eb6754399b2947e643aafb712", CreateTime: carbon.Now().ToDateTimeString(), ButtonName: "查询", ButtonIdentity: "search", ButtonFullIdentity: "superAdmin:accessRecord:search", PageName: "访问历史", PagePath: "superAdmin/accessRecord", MenuId: "8b015470a20a47eaa660808a5b793d52"},
	{ID: "51320ea481c2463cbbb6deec97358dd9", CreateTime: carbon.Now().ToDateTimeString(), ButtonName: "添加子菜单", ButtonIdentity: "addChildMenu", ButtonFullIdentity: "superAdmin:menu:addChildMenu", PageName: "菜单管理", PagePath: "superAdmin/menu", MenuId: "648d2467026e481383fd8a62901ba22d"},
	{ID: "5a7f3b36b1c4470d8d5c27eac6ad0f4a", CreateTime: carbon.Now().ToDateTimeString(), ButtonName: "查询", ButtonIdentity: "search", ButtonFullIdentity: "superAdmin:authority:search", PageName: "角色管理", PagePath: "superAdmin/authority", MenuId: "ece2a0a3c7c3448490d5ecafb9121ad7"},
	{ID: "5c5f49f0689d4216a24dbdd631e9d39f", CreateTime: carbon.Now().ToDateTimeString(), ButtonName: "查询", ButtonIdentity: "search", ButtonFullIdentity: "superAdmin:sysOperationRecord:search", PageName: "操作历史", PagePath: "superAdmin/sysOperationRecord", MenuId: "e8566ed7deda4f88a6ad0a4b8974bda8"},
	{ID: "65446bfb9f0a4aa8951100dcae212e3d", CreateTime: carbon.Now().ToDateTimeString(), ButtonName: "删除", ButtonIdentity: "delete", ButtonFullIdentity: "superAdmin:authority:delete", PageName: "角色管理", PagePath: "superAdmin/authority", MenuId: "ece2a0a3c7c3448490d5ecafb9121ad7"},
	{ID: "67a8cf20b1b34543a3a0edfd44a4d8ea", CreateTime: carbon.Now().ToDateTimeString(), ButtonName: "查询", ButtonIdentity: "search", ButtonFullIdentity: "merchant:merchantList:search", PageName: "商户管理", PagePath: "merchant/merchantList", MenuId: "651515f100e34fbc805fa2d7019bc335"},
	{ID: "8a3c27cfc136408482d08428d088a54b", CreateTime: carbon.Now().ToDateTimeString(), ButtonName: "删除", ButtonIdentity: "delete", ButtonFullIdentity: "superAdmin:api:delete", PageName: "api管理", PagePath: "superAdmin/api", MenuId: "a8abac9ed1d741ac8a23bab26acef163"},
	{ID: "927f709cb44f4143b05d8b3d9deaa0c8", CreateTime: carbon.Now().ToDateTimeString(), ButtonName: "设置权限", ButtonIdentity: "setAuthority", ButtonFullIdentity: "superAdmin:authority:setAuthority", PageName: "角色管理", PagePath: "superAdmin/authority", MenuId: "ece2a0a3c7c3448490d5ecafb9121ad7"},
	{ID: "9e1d5bc5a0d94a2991d58362dc74d75b", CreateTime: carbon.Now().ToDateTimeString(), ButtonName: "新增", ButtonIdentity: "add", ButtonFullIdentity: "superAdmin:authority:add", PageName: "角色管理", PagePath: "superAdmin/authority", MenuId: "ece2a0a3c7c3448490d5ecafb9121ad7"},
	{ID: "b58e51ae3b324e5c85b85fbc836081d2", CreateTime: carbon.Now().ToDateTimeString(), ButtonName: "编辑", ButtonIdentity: "edit", ButtonFullIdentity: "superAdmin:menu:edit", PageName: "菜单管理", PagePath: "superAdmin/menu", MenuId: "648d2467026e481383fd8a62901ba22d"},
	{ID: "b86d4f6d8e714e709c3c8e5942eddd73", CreateTime: carbon.Now().ToDateTimeString(), ButtonName: "查询", ButtonIdentity: "search", ButtonFullIdentity: "superAdmin:user:search", PageName: "用户管理", PagePath: "superAdmin/user", MenuId: "5123afd2a3ac489f95fbc4b7865ba3d3"},
	{ID: "bb391d8b756e4993bbb29c75610f0ddf", CreateTime: carbon.Now().ToDateTimeString(), ButtonName: "删除", ButtonIdentity: "delete", ButtonFullIdentity: "superAdmin:user:delete", PageName: "用户管理", PagePath: "superAdmin/user", MenuId: "5123afd2a3ac489f95fbc4b7865ba3d3"},
	{ID: "bb59c221d2744512b1ecc6b4074a8d63", CreateTime: carbon.Now().ToDateTimeString(), ButtonName: "新增", ButtonIdentity: "add", ButtonFullIdentity: "superAdmin:api:add", PageName: "api管理", PagePath: "superAdmin/api", MenuId: "a8abac9ed1d741ac8a23bab26acef163"},
	{ID: "bfdd59c6fbb94f579fd4e40b297b3504", CreateTime: carbon.Now().ToDateTimeString(), ButtonName: "查询", ButtonIdentity: "search", ButtonFullIdentity: "superAdmin:api:search", PageName: "api管理", PagePath: "superAdmin/api", MenuId: "a8abac9ed1d741ac8a23bab26acef163"},
	{ID: "c37d95effd5d49bd9fa42b766120e86f", CreateTime: carbon.Now().ToDateTimeString(), ButtonName: "新增", ButtonIdentity: "add", ButtonFullIdentity: "merchant:merchantList:add", PageName: "商户管理", PagePath: "merchant/merchantList", MenuId: "651515f100e34fbc805fa2d7019bc335"},
	{ID: "d3eb8a7b98ad42c9aacfb47a5b232325", CreateTime: carbon.Now().ToDateTimeString(), ButtonName: "查询", ButtonIdentity: "search", ButtonFullIdentity: "superAdmin:menu:search", PageName: "菜单管理", PagePath: "superAdmin/menu", MenuId: "648d2467026e481383fd8a62901ba22d"},
	{ID: "d4be75cb677c4128aaeb69929ccfbd10", CreateTime: carbon.Now().ToDateTimeString(), ButtonName: "新增", ButtonIdentity: "add", ButtonFullIdentity: "superAdmin:menu:add", PageName: "菜单管理", PagePath: "superAdmin/menu", MenuId: "648d2467026e481383fd8a62901ba22d"},
	{ID: "e00db52d9db742159286ee7ede5a76fd", CreateTime: carbon.Now().ToDateTimeString(), ButtonName: "删除", ButtonIdentity: "delete", ButtonFullIdentity: "superAdmin:menu:delete", PageName: "菜单管理", PagePath: "superAdmin/menu", MenuId: "648d2467026e481383fd8a62901ba22d"},
	{ID: "e9708e5f0a0142be96276355be4e828d", CreateTime: carbon.Now().ToDateTimeString(), ButtonName: "编辑", ButtonIdentity: "edit", ButtonFullIdentity: "superAdmin:authority:edit", PageName: "角色管理", PagePath: "superAdmin/authority", MenuId: "ece2a0a3c7c3448490d5ecafb9121ad7"},
	{ID: "fedf370edfe94c199d3ba086d968206f", CreateTime: carbon.Now().ToDateTimeString(), ButtonName: "新增", ButtonIdentity: "add", ButtonFullIdentity: "superAdmin:user:add", PageName: "用户管理", PagePath: "superAdmin/user", MenuId: "5123afd2a3ac489f95fbc4b7865ba3d3"},
}

// Init
/**
 *  @Description: sys_button表初始化
 *  @receiver b
 *  @return error
 */
func (b *button) Init() error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("page_name IN ?", []string{"商户管理","用户管理","访问历史","菜单管理","角色管理","操作历史","api管理"}).Find(&[]menumod.Button{}).RowsAffected > 0 {
			color.Danger.Println("sys_button 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&buttons).Error; err != nil {
			// 遇到错误时回滚事务
			return err
		}
		color.Info.Println("sys_button 表初始数据成功!")
		return nil
	})
}
