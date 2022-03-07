/**
 * @Time: 2022/3/7 16:10
 * @Author: yt.yin
 */

package mchtapi

import mchtsrv "github.com/goworkeryyt/aramis/server/merchant/service"

type ApiGroup struct {
	MerchantApi
}

var ApiGroupApp = new(ApiGroup)

var merchantService = mchtsrv.ServiceGroupApp.MerchantService
