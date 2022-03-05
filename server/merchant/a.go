/**
 * @Time: 2022/3/5 15:37
 * @Author: yt.yin
 */

package merchant

import (
	"github.com/goworkeryyt/aramis/server/merchant/model"
	"github.com/goworkeryyt/go-core/global"
	"go.uber.org/zap"
)

// 引入的时候自动创建表
func autoCreateTables() {
	if global.DB != nil {
		// 数据库自动迁移
		err := global.DB.AutoMigrate(
			mchtmod.Merchant{},
		)
		if err != nil && global.LOG != nil {
			global.LOG.Error("初始化表时异常：", zap.Any("err", err))
		}
	}
}
