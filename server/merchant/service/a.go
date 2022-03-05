/**
 * @Time: 2022/3/5 15:38
 * @Author: yt.yin
 */

package mchtsrv

type ServiceGroup struct {
	MerchantService
}

var ServiceGroupApp = new(ServiceGroup)
