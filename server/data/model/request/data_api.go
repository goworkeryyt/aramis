/**
 * @Time: 2022/3/7 18:59
 * @Author: yt.yin
 */

package datareq

// UpdateApiInfoReq 更新API的请求参数
type UpdateApiInfoReq struct {

	/** 记录id */
	ID               string      `json:"id"`

	/** API中文描述 */
	Description      string      `json:"description"`

	/** API组 */
	ApiGroup         string      `json:"apiGroup"`
}




