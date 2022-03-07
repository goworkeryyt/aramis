/**
 * @Time: 2022/3/7 11:25
 * @Author: yt.yin
 */

package menuapi

import "github.com/goworkeryyt/aramis/server/menu/service"

type ApiGroup struct {
	MenuApi
}

var ApiGroupApp = new(ApiGroup)

var menuService = menusrv.ServiceGroupApp.MenuService
