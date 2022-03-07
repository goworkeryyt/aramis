/**
 * @Time: 2022/3/7 20:02
 * @Author: yt.yin
 */

package usersrv

import  "github.com/goworkeryyt/aramis/server/role/service"

type ServiceGroup struct {
	UserService
	UserRoleService
}

var ServiceGroupApp = new(ServiceGroup)

var roleService = rolesrv.ServiceGroupApp.RoleService