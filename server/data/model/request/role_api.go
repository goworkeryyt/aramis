/**
 * @Time: 2022/3/7 19:01
 * @Author: yt.yin
 */

package datareq

// RoleApiBindReq 角色API绑定请求参数
type RoleApiBindReq struct {

	/** 角色ID */
	RoleId        string      `json:"roleId"`

	/** API id 集合 */
	ApiInfoIds   []string     `json:"apiIds"`
}
