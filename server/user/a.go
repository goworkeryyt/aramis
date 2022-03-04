/**
 * @Time: 2022/3/4 11:12
 * @Author: yt.yin
 */

package user

import (
	"aramis/server/user/model"
	"aramis/server/user/model/source"
	"github.com/goworkeryyt/go-core/global"
	"go.uber.org/zap"
)

// 引入的时候自动创建表（用户表，角色表）
func autoCreateTables() {
	if global.DB != nil {
		// 数据库自动迁移
		err := global.DB.AutoMigrate(
			usermod.User{},
		)
		if err != nil && global.LOG != nil {
			global.LOG.Error("初始化表时异常：", zap.Any("err", err))
		}
		// 初始化数据
		err = source.Admin.Init()
		if err != nil {
			global.LOG.Error("初始化数据异常：", zap.Any("err", err))
		}
	}
}
