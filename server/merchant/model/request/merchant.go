/**
 * @Time: 2022/3/5 15:43
 * @Author: yt.yin
 */

package mchtreq

// UpdateMerchantReq 更新商户请求参数
type UpdateMerchantReq struct {

	/** 商户ID */
	ID                     string          `json:"id"`

	/** 商户名称 */
	MerchantName           string          `json:"merchantName"`

	/** 地址 */
	Address                string          `json:"address"`

	/** 区域代码 */
	RegionCode             string          `json:"regionCode"`

	/** 区域名称 */
	RegionName             string          `json:"regionName"`

	/** 状态 */
	Status                 string          `json:"status"`

	/** 商户密钥 */
	SelfPriKey             string          `json:"selfPriKey"` // 商户密钥
}



