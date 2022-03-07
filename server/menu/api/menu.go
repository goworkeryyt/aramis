/**
 * @Time: 2022/3/7 11:24
 * @Author: yt.yin
 */

package menuapi

import (
	"github.com/gin-gonic/gin"
	"github.com/goworkeryyt/aramis/server/menu/model"
	"github.com/goworkeryyt/go-core/global"
	"github.com/goworkeryyt/go-toolbox/request"
	"github.com/goworkeryyt/go-toolbox/result"
	"github.com/goworkeryyt/go-toolbox/validator"
	"go.uber.org/zap"
)

type MenuApi struct{}

// CreateMenu 创建菜单
func (m *MenuApi) CreateMenu(c *gin.Context) {
	// json绑定
	var menu menumod.Menu
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		result.FailMsg("绑定的Json格式有误", c)
		return
	}
	// 数据效验
	if err = validator.Verify(menu, menumod.MenuCreateVerify); err != nil {
		result.FailMsg(err.Error(), c)
		return
	}
	if err = menuService.CreateMenu(menu); err != nil {
		global.LOG.Error("创建失败:", zap.Any("err", err))
		result.FailMsg("创建失败:"+err.Error(), c)
	} else {
		result.OkMsg("创建成功", c)
	}
}

// DeleteMenu 删除菜单(单个)
func (m *MenuApi) DeleteMenu(c *gin.Context) {
	// json绑定
	var rqId request.ID
	err := c.ShouldBindJSON(&rqId)
	if err != nil {
		result.FailMsg("绑定的Json格式有误", c)
		return
	}
	if err = menuService.DeleteMenu(rqId.ID); err != nil {
		global.LOG.Error("删除失败:", zap.Any("err", err))
		result.FailMsg("删除失败,"+err.Error(), c)
	} else {
		result.OkMsg("删除成功", c)
	}
}

// UpdateMenu 更新菜单
func (m *MenuApi) UpdateMenu(c *gin.Context) {
	// json绑定
	var menu menumod.Menu
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		result.FailMsg("绑定的Json格式有误", c)
		return
	}
	// 数据效验
	if err = validator.Verify(menu, menumod.MenuCreateVerify); err != nil {
		result.FailMsg(err.Error(), c)
		return
	}
	// 判断隐藏状态是否合法
	if menu.IsShow != menumod.SHOW && menu.IsShow != menumod.HIDE {
		result.FailMsg("非法的菜单隐藏标记", c)
		return
	}

	if err = menuService.UpdateMenu(menu); err != nil {
		global.LOG.Error("更新失败:", zap.Any("err", err))
		result.FailMsg("更新失败:"+err.Error(), c)
	} else {
		result.OkMsg("更新成功", c)
	}
}

// FindMenu 查询菜单(详情)
func (m *MenuApi) FindMenu(c *gin.Context) {
	// json绑定
	var rqId request.ID
	_ = c.ShouldBindQuery(&rqId)

	if menuInfo, err := menuService.GetMenuInfo(rqId.ID); err != nil {
		global.LOG.Error("查询失败:", zap.Any("err", err))
		result.FailMsg("查询失败:"+err.Error(), c)
	} else {
		result.OkData(menuInfo, c)
	}
}

