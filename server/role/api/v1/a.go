/**
 * @Time: 2022/3/7 17:53
 * @Author: yt.yin
 */

package roleapi

import (
	"github.com/goworkeryyt/aramis/server/menu/service"
	"github.com/goworkeryyt/aramis/server/role/service"
)

type ApiGroup struct {
	RoleApi
	RoleMenuApi
	RoleButtonApi
}

var ApiGroupApp = new(ApiGroup)

var roleService = rolesrv.ServiceGroupApp.RoleService
var roleMenuService = rolesrv.ServiceGroupApp.RoleMenuService
var roleButtonService = rolesrv.ServiceGroupApp.RoleButtonService
var menuService = menusrv.ServiceGroupApp.MenuService