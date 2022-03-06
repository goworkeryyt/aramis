/**
 * @Time: 2022/3/5 15:19
 * @Author: yt.yin
 */

package button

import (
	"github.com/goworkeryyt/aramis/server/button/model"
	"github.com/goworkeryyt/go-core/global"
	"go.uber.org/zap"
)

// 引入的时候自动创建表
func autoCreateTables() {
	if global.DB != nil {
		// 数据库自动迁移
		err := global.DB.AutoMigrate(
			btnmod.Button{},
		)
		if err != nil && global.LOG != nil {
			global.LOG.Error("初始化表时异常：", zap.Any("err", err))
		}
	}
}

