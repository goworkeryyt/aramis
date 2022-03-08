/**
 * @Time: 2022/3/8 10:57
 * @Author: yt.yin
 */

package aramis

import (
	"github.com/gin-gonic/gin"
	"github.com/goworkeryyt/aramis/server/data"
	"github.com/goworkeryyt/aramis/server/menu"
	"github.com/goworkeryyt/aramis/server/merchant"
	"github.com/goworkeryyt/aramis/server/role"
	"github.com/goworkeryyt/aramis/server/source"
	"github.com/goworkeryyt/aramis/server/user"
	"github.com/goworkeryyt/go-core/global"
	"github.com/goworkeryyt/go-middle/record"
	"go.uber.org/zap"
)

// 数据初始化
func dbInit(){
	err := source.Admin.Init()
	if err != nil{
		global.LOG.Error("初始化admin账户异常:",zap.Any("err",err))
	}
	err = source.Menu.Init()
	if err != nil{
		global.LOG.Error("初始化系统菜单异常:",zap.Any("err",err))
	}
	err = source.Button.Init()
	if err != nil{
		global.LOG.Error("初始化按钮数据异常:",zap.Any("err",err))
	}
	err = source.Role.Init()
	if err != nil{
		global.LOG.Error("初始化角色数据异常:",zap.Any("err",err))
	}
	err = source.RoleMenu.Init()
	if err != nil{
		global.LOG.Error("初始化角色菜单异常:",zap.Any("err",err))
	}
	err = source.RoleButton.Init()
	if err != nil{
		global.LOG.Error("初始化角色按钮异常:",zap.Any("err",err))
	}
	err = source.ApiInfo.Init()
	if err != nil{
		global.LOG.Error("初始化数据资源异常:",zap.Any("err",err))
	}
	err = source.RoleApi.Init()
	if err != nil{
		global.LOG.Error("初始化角色资源异常:",zap.Any("err",err))
	}
}

// RouterRegister 注册阿拉姆斯所有路由
func RouterRegister(rGroup *gin.RouterGroup) {
	// 菜单
	menu.RouterRegister(rGroup)
	// 数据api
	data.RouterRegister(rGroup)
	// 角色
	role.RouterRegister(rGroup)
	// 用户账号
	user.RouterRegister(rGroup)
	// 商户
	merchant.RouterRegister(rGroup)
	// 访问记录
	record.RouterRegister(rGroup)
	// 初始化数据
	dbInit()
}
