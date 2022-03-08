/**
 * @Time: 2022/3/8 10:44
 * @Author: yt.yin
 */

package source

import (
	"github.com/gookit/color"
	"github.com/goworkeryyt/go-core/global"
)

var RoleApi = new(roleApi)

type roleApi struct{}

var rules = [][]string{
	{"90cb58fff9244754970711f2c4bd41a0", "/api/apiInfo/createApi", "POST"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/apiInfo/deleteApi", "POST"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/apiInfo/findApiPage", "GET"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/apiInfo/findApiTree", "GET"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/apiInfo/findRoleApi", "GET"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/apiInfo/roleApiBind", "POST"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/menu/createMenu", "POST"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/menu/deleteMenu", "POST"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/menu/findMenu", "GET"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/menu/getRoleMenu", "GET"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/menu/updateMenu", "POST"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/merchant/createMerchant", "POST"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/merchant/deleteMerchant", "POST"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/merchant/getAllMerchants", "GET"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/merchant/getMerchantPage", "GET"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/merchant/updateMerchant", "POST"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/record/getAccessRecordPage", "GET"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/record/getOperateRecordPage", "GET"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/role/createRole", "POST"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/role/deleteRole", "POST"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/role/findAllRoles", "GET"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/role/findRolePage", "GET"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/role/getRoleMenu", "GET"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/role/getSelfButtons", "GET"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/role/getSelfMenuTree", "GET"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/role/roleMenuBind", "POST"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/role/updateRole", "POST"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/user/captcha", "GET"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/user/changePassword", "POST"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/user/createUser", "POST"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/user/deleteUser", "POST"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/user/getSelfUserInfo", "GET"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/user/getUserPage", "GET"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/user/login", "POST"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/user/updateUser", "POST"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/user/updateUserPassword", "POST"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/user/updateUserRoles", "POST"},
	{"90cb58fff9244754970711f2c4bd41a0", "/api/user/updateUserStatus", "POST"},
	{"ab6da11b90524d48867eefb53ba89fc6", "/api/apiInfo/findApiTree", "GET"},
	{"ab6da11b90524d48867eefb53ba89fc6", "/api/apiInfo/findRoleApi", "GET"},
	{"ab6da11b90524d48867eefb53ba89fc6", "/api/apiInfo/roleApiBind", "POST"},
	{"ab6da11b90524d48867eefb53ba89fc6", "/api/menu/getRoleMenu", "GET"},
	{"ab6da11b90524d48867eefb53ba89fc6", "/api/merchant/createMerchant", "POST"},
	{"ab6da11b90524d48867eefb53ba89fc6", "/api/merchant/deleteMerchant", "POST"},
	{"ab6da11b90524d48867eefb53ba89fc6", "/api/merchant/getAllMerchants", "GET"},
	{"ab6da11b90524d48867eefb53ba89fc6", "/api/merchant/getMerchantPage", "GET"},
	{"ab6da11b90524d48867eefb53ba89fc6", "/api/merchant/updateMerchant", "POST"},
	{"ab6da11b90524d48867eefb53ba89fc6", "/api/record/getAccessRecordPage", "GET"},
	{"ab6da11b90524d48867eefb53ba89fc6", "/api/record/getOperateRecordPage", "GET"},
	{"ab6da11b90524d48867eefb53ba89fc6", "/api/role/createRole", "POST"},
	{"ab6da11b90524d48867eefb53ba89fc6", "/api/role/deleteRole", "POST"},
	{"ab6da11b90524d48867eefb53ba89fc6", "/api/role/findAllRoles", "GET"},
	{"ab6da11b90524d48867eefb53ba89fc6", "/api/role/findRolePage", "GET"},
	{"ab6da11b90524d48867eefb53ba89fc6", "/api/role/getRoleMenu", "GET"},
	{"ab6da11b90524d48867eefb53ba89fc6", "/api/role/getSelfButtons", "GET"},
	{"ab6da11b90524d48867eefb53ba89fc6", "/api/role/getSelfMenuTree", "GET"},
	{"ab6da11b90524d48867eefb53ba89fc6", "/api/role/roleMenuBind", "POST"},
	{"ab6da11b90524d48867eefb53ba89fc6", "/api/role/updateRole", "POST"},
	{"ab6da11b90524d48867eefb53ba89fc6", "/api/user/captcha", "GET"},
	{"ab6da11b90524d48867eefb53ba89fc6", "/api/user/changePassword", "POST"},
	{"ab6da11b90524d48867eefb53ba89fc6", "/api/user/createUser", "POST"},
	{"ab6da11b90524d48867eefb53ba89fc6", "/api/user/deleteUser", "POST"},
	{"ab6da11b90524d48867eefb53ba89fc6", "/api/user/getSelfUserInfo", "GET"},
	{"ab6da11b90524d48867eefb53ba89fc6", "/api/user/getUserPage", "GET"},
	{"ab6da11b90524d48867eefb53ba89fc6", "/api/user/login", "POST"},
	{"ab6da11b90524d48867eefb53ba89fc6", "/api/user/updateUser", "POST"},
	{"ab6da11b90524d48867eefb53ba89fc6", "/api/user/updateUserPassword", "POST"},
	{"ab6da11b90524d48867eefb53ba89fc6", "/api/user/updateUserRoles", "POST"},
	{"ab6da11b90524d48867eefb53ba89fc6", "/api/user/updateUserStatus", "POST"},
	{"ab6da11b90524d48867eefb53ba89fc6", "/health", "GET"},
}

// Init
/**
 *  @Description: 角色api权限初始化
 *  @receiver a
 *  @return error
 */
func (r *roleApi) Init() error {
	policies := global.CSBEF.GetPolicy()
	// 没有api权限才创建
	if len(policies) == 0 {
		_, err := global.CSBEF.AddPolicies(rules)
		if err != nil {
			color.Info.Println("角色api权限初始数据失败!")
			return err
		}
		color.Info.Println("角色api权限初始数据成功!")
		return nil
	}
	return nil
}

