/**
 * @Time: 2022/3/7 17:55
 * @Author: yt.yin
 */

package rolesrv

type ServiceGroup struct {
	RoleService
	RoleMenuService
	RoleButtonService
}

var ServiceGroupApp = new(ServiceGroup)
