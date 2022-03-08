/**
 * @Time: 2022/3/8 10:28
 * @Author: yt.yin
 */

package source

import (
	"github.com/golang-module/carbon/v2"
	"github.com/gookit/color"
	"github.com/goworkeryyt/aramis/server/data/model"
	"github.com/goworkeryyt/go-core/global"
	"gorm.io/gorm"
)

var ApiInfo = new(apiInfo)

type apiInfo struct{}

var apiInfoSource = []datamod.ApiInfo{
	{ID: "197a9a67271746a89922a6f0df57d867", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/apiInfo/createApi",           Method: "POST", ApiGroup: "API管理", Description: "创建API"},
	{ID: "197a9a67271746a89922a6f0df57d868", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/apiInfo/deleteApi",           Method: "POST", ApiGroup: "API管理", Description: "删除API(单个)"},
	{ID: "197a9a67271746a89922a6f0df57d869", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/apiInfo/findApiPage",         Method: "GET",  ApiGroup: "API管理", Description: "查询API(分页)"},
	{ID: "197a9a67271746a89922a6f0df57d870", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/apiInfo/findApis",            Method: "GET",  ApiGroup: "未用到",   Description: "查询API(所有)"},
	{ID: "197a9a67271746a89922a6f0df57d871", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/apiInfo/findApiTree",         Method: "GET",  ApiGroup: "角色管理", Description: "查询API(树)"},
	{ID: "197a9a67271746a89922a6f0df57d872", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/apiInfo/findRoleApi",         Method: "GET",  ApiGroup: "角色管理", Description: "查询API(角色)"},
	{ID: "197a9a67271746a89922a6f0df57d873", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/apiInfo/roleApiBind",         Method: "POST", ApiGroup: "角色管理", Description: "角色API绑定"},
	{ID: "197a9a67271746a89922a6f0df57d880", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/menu/createMenu",             Method: "POST", ApiGroup: "菜单管理", Description: "创建菜单"},
	{ID: "197a9a67271746a89922a6f0df57d881", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/menu/deleteMenu",             Method: "POST", ApiGroup: "菜单管理", Description: "删除菜单(单个)"},
	{ID: "197a9a67271746a89922a6f0df57d882", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/menu/findMenu",               Method: "GET",  ApiGroup: "菜单管理", Description: "查询菜单(单个)"},
	{ID: "197a9a67271746a89922a6f0df57d883", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/menu/findMenuTreeSelf",       Method: "GET",  ApiGroup: "未用到",   Description: "查询菜单树(当前用户)"},
	{ID: "197a9a67271746a89922a6f0df57d884", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/menu/getRoleMenu",            Method: "GET",  ApiGroup: "角色管理", Description: "查询菜单(角色)"},
	{ID: "197a9a67271746a89922a6f0df57d885", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/menu/updateMenu",             Method: "POST", ApiGroup: "菜单管理", Description: "更新菜单"},
	{ID: "197a9a67271746a89922a6f0df57d886", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/merchant/createMerchant",     Method: "POST", ApiGroup: "商户管理", Description: "创建商户"},
	{ID: "197a9a67271746a89922a6f0df57d887", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/merchant/deleteMerchant",     Method: "POST", ApiGroup: "商户管理", Description: "删除商户(单个)"},
	{ID: "197a9a67271746a89922a6f0df57d888", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/merchant/getAllMerchants",    Method: "GET",  ApiGroup: "商户管理", Description: "查询商户(所有)"},
	{ID: "197a9a67271746a89922a6f0df57d889", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/merchant/getMerchantPage",    Method: "GET",  ApiGroup: "商户管理", Description: "查询商户(分页)"},
	{ID: "197a9a67271746a89922a6f0df57d890", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/merchant/updateMerchant",     Method: "POST", ApiGroup: "商户管理", Description: "更新商户"},
	{ID: "197a9a67271746a89922a6f0df57d891", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/record/getAccessRecordPage",  Method: "GET",  ApiGroup: "日志",     Description: "访问日志"},
	{ID: "197a9a67271746a89922a6f0df57d892", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/record/getOperateRecordPage", Method: "GET",  ApiGroup: "日志",     Description: "操作日志"},
	{ID: "197a9a67271746a89922a6f0df57d893", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/role/createRole",             Method: "POST", ApiGroup: "角色管理", Description: "创建角色"},
	{ID: "197a9a67271746a89922a6f0df57d894", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/role/deleteRole",             Method: "POST", ApiGroup: "角色管理", Description: "删除角色"},
	{ID: "197a9a67271746a89922a6f0df57d895", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/role/findAllRoles",           Method: "GET",  ApiGroup: "角色管理", Description: "查询角色(所有)"},
	{ID: "197a9a67271746a89922a6f0df57d896", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/role/findRole",               Method: "GET",  ApiGroup: "未用到",   Description: "查询角色(详情)"},
	{ID: "197a9a67271746a89922a6f0df57d897", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/role/findRolePage",           Method: "GET",  ApiGroup: "角色管理", Description: "查询角色(分页)"},
	{ID: "197a9a67271746a89922a6f0df57d898", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/role/getRoleMenu",            Method: "GET",  ApiGroup: "角色管理", Description: "获取角色菜单"},
	{ID: "197a9a67271746a89922a6f0df57d899", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/role/getSelfButtons",         Method: "GET",  ApiGroup: "角色管理", Description: "获取自身按钮权限"},
	{ID: "197a9a67271746a89922a6f0df57d900", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/role/getSelfMenuTree",        Method: "GET",  ApiGroup: "菜单管理", Description: "获取用户菜单权限树及按钮"},
	{ID: "197a9a67271746a89922a6f0df57d901", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/role/roleMenuBind",           Method: "POST", ApiGroup: "角色管理", Description: "角色菜单及按钮绑定"},
	{ID: "197a9a67271746a89922a6f0df57d902", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/role/updateRole",             Method: "POST", ApiGroup: "角色管理", Description: "更新角色"},
	{ID: "197a9a67271746a89922a6f0df57d903", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/user/captcha",                Method: "GET",  ApiGroup: "登录",     Description: "获取验证码"},
	{ID: "197a9a67271746a89922a6f0df57d904", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/user/changePassword",         Method: "POST", ApiGroup: "登录",     Description: "修改密码"},
	{ID: "197a9a67271746a89922a6f0df57d905", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/user/createUser",             Method: "POST", ApiGroup: "用户管理", Description: "创建用户"},
	{ID: "197a9a67271746a89922a6f0df57d906", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/user/deleteUser",             Method: "POST", ApiGroup: "用户管理", Description: "删除用户(单个)"},
	{ID: "197a9a67271746a89922a6f0df57d907", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/user/getSelfUserInfo",        Method: "GET",  ApiGroup: "登录",     Description: "查询用户(当前)"},
	{ID: "197a9a67271746a89922a6f0df57d908", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/user/getUserPage",            Method: "GET",  ApiGroup: "用户管理", Description: "查询用户(分页)"},
	{ID: "197a9a67271746a89922a6f0df57d909", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/user/login",                  Method: "POST", ApiGroup: "登录",    Description: "登录"},
	{ID: "197a9a67271746a89922a6f0df57d910", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/user/updateUser",             Method: "POST", ApiGroup: "用户管理", Description: "更新用户"},
	{ID: "197a9a67271746a89922a6f0df57d911", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/user/updateUserPassword",     Method: "POST", ApiGroup: "用户管理", Description: "更新用户(密码)"},
	{ID: "197a9a67271746a89922a6f0df57d912", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/user/updateUserRoles",        Method: "POST", ApiGroup: "用户管理", Description: "更新用户(角色)"},
	{ID: "197a9a67271746a89922a6f0df57d913", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/user/updateUserStatus",       Method: "POST", ApiGroup: "用户管理", Description: "更新用户(状态)"},
	{ID: "197a9a67271746a89922a6f0df57d914", CreateTime: carbon.Now().ToDateTimeString(), Path: "/health",                          Method: "GET",  ApiGroup: "健康监测", Description: "健康监测"},
	{ID: "197a9a67271746a89922a6f0df57d915", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/merchant/getAllMerchants",    Method: "GET",  ApiGroup: "商户管理", Description: "查询商户(所有)"},
	{ID: "197a9a67271746a89922a6f0df57d916", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/merchant/getAllMerchants",    Method: "GET",  ApiGroup: "登录",    Description: "查询商户(所有)"},
	{ID: "197a9a67271746a89922a6f0df57d917", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/merchant/getAllMerchants",    Method: "GET",  ApiGroup: "用户管理", Description: "查询商户(所有)"},
	{ID: "197a9a67271746a89922a6f0df57d918", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/role/findAllRoles",           Method: "GET",  ApiGroup: "用户管理", Description: "查询角色(所有)"},
	{ID: "197a9a67271746a89922a6f0df57d919", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/role/getSelfMenuTree",        Method: "GET",  ApiGroup: "角色管理", Description: "获取用户菜单权限树及按钮"},
	{ID: "197a9a67271746a89922a6f0df57d920", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/role/getSelfMenuTree",        Method: "GET",  ApiGroup: "登录",    Description: "获取用户菜单权限树及按钮"},
	{ID: "b6b47ea157e14b1ea9c6fca64f9322da", CreateTime: carbon.Now().ToDateTimeString(), Path: "/api/role/getSelfButtons",         Method: "GET",  ApiGroup: "登录",    Description: "获取自身按钮权限"},
}

// Init
/**
 *  @Description: sys_user表初始化
 *  @receiver a
 *  @return error
 */
func (a *apiInfo) Init() error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("api_group IN ?", []string{"API管理", "角色管理", "初始化", "菜单管理", "商户管理", "日志", "登录", "用户管理", "健康监测"}).Find(&[]datamod.ApiInfo{}).RowsAffected > 0 {
			color.Danger.Println("api 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&apiInfoSource).Error; err != nil {
			return err
		}
		color.Info.Println("api 表初始数据成功!")
		return nil
	})
}

