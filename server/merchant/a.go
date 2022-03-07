/**
 * @Time: 2022/3/5 15:37
 * @Author: yt.yin
 */

package merchant

import (
	"encoding/json"
	"strings"
	"sync"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/gin-gonic/gin"
	"github.com/goworkeryyt/aramis/server/merchant/api/v1"
	"github.com/goworkeryyt/aramis/server/merchant/model"
	"github.com/goworkeryyt/go-core/global"
	"github.com/goworkeryyt/go-middle/record/operate"
	"go.uber.org/zap"

)

var (
	// Cache 商户缓存
	Cache *bigcache.BigCache
	lock  sync.Mutex
)

// 创建商户缓存对象
func createCache() {
	config := bigcache.Config{
		Shards:             1024,          // 存储的条目数量，值必须是2的幂
		LifeWindow:         4 * time.Hour, // 超时后条目被删除
		CleanWindow:        1 * time.Hour, // 清除间隔
		MaxEntriesInWindow: 0,             // 在 Life Window 中的最大数量，
		MaxEntrySize:       0,             // 条目最大尺寸，以字节为单位
		HardMaxCacheSize:   0,             // 设置缓存最大值，以MB为单位，超过了不在分配内存。0表示无限制分配
	}
	cache, err := bigcache.NewBigCache(config)
	if err != nil {
		global.LOG.Error("缓存对象创建异常:" + err.Error())
		return
	}
	lock.Lock()
	Cache = cache
	lock.Unlock()
}

// 缓存商户对象
func setMerchant(merchant mchtmod.Merchant) {
	if merchant.MerchantNo == "" {
		return
	}
	if Cache == nil {
		createCache()
	}
	jsonby, err := json.Marshal(merchant)
	if err != nil {
		global.LOG.Error("商户对象缓存时转json异常:" + err.Error())
		return
	}
	err = Cache.Set(merchant.MerchantNo, jsonby)
	if err != nil {
		global.LOG.Error("商户对象缓存异常:" + err.Error())
		return
	}
}

// GetMerchant 从缓存中取商户对象
func GetMerchant(merchantNo string) (merchant mchtmod.Merchant) {
	merchantNo = strings.TrimSpace(merchantNo)
	if merchantNo == "" {
		return
	}
	if Cache == nil {
		createCache()
	}
	entry, err := Cache.Get(merchantNo)
	if err == nil {
		err = json.Unmarshal(entry, &merchant)
		if err == nil {
			// 在缓存中找到 返回
			return merchant
		}
	}
	err = global.DB.Where("merchant_no = ?", merchantNo).First(&merchant).Error
	if err != nil{
		return
	}
	if merchant.ID != "" {
		setMerchant(merchant)
		return merchant
	}
	return
}

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

// RouterRegister 注册商户模块路由
func RouterRegister(rGroup *gin.RouterGroup) {
	// 初始化表
	autoCreateTables()
	// 创建路由
	{
		api := mchtapi.ApiGroupApp
		merchantGroup := rGroup.Group("merchant")
		merchantGroup.Use(
			operate.OperateRecordHandler(365),
		)
		{
			// 创建商户
			merchantGroup.POST("createMerchant", api.CreateMerchant)
			// 更新商户
			merchantGroup.POST("updateMerchant", api.UpdateMerchant)
			// 删除商户
			merchantGroup.POST("deleteMerchant", api.DeleteMerchant)
		}
		merchantGroupWithoutRecord := rGroup.Group("merchant")
		{
			// 查询商户（分页）
			merchantGroupWithoutRecord.GET("getMerchantPage", api.GetMerchantPage)
			// 查询商户（所有）
			merchantGroupWithoutRecord.GET("getAllMerchants", api.GetAllMerchants)
		}
	}
}
