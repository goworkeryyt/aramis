/**
 * @Time: 2022/3/7 19:20
 * @Author: yt.yin
 */

package datasrv

type ServiceGroup struct {
	ApiInfoService
	RoleApiService
}

var ServiceGroupApp = new(ServiceGroup)
