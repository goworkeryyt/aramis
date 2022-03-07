/**
 * @Time: 2022/3/7 20:25
 * @Author: yt.yin
 */

package userapi

import (
	"github.com/goworkeryyt/aramis/server/role/service"
	"github.com/goworkeryyt/aramis/server/user/service"
)

type ApiGroup struct {
	UserApi
	LoginApi
}

var ApiGroupApp = new(ApiGroup)

var userService = usersrv.ServiceGroupApp.UserService
var roleService = rolesrv.ServiceGroupApp.RoleService
var userRoleService = usersrv.ServiceGroupApp.UserRoleService

