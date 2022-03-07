/**
 * @Time: 2022/3/7 19:33
 * @Author: yt.yin
 */

package datapi

import  "github.com/goworkeryyt/aramis/server/data/service"

type ApiGroup struct {
	DataApi
	RoleDataApi
}

var ApiGroupApp = new(ApiGroup)

var apiInfoService = datasrv.ServiceGroupApp.ApiInfoService
var roleApiService = datasrv.ServiceGroupApp.RoleApiService
