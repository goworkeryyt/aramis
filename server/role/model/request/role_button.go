/**
 * @Time: 2022/3/7 18:32
 * @Author: yt.yin
 */

package rolereq

// RoleButtonBindRequest 角色按钮绑定请求参数
type RoleButtonBindRequest struct {

	/** 角色ID */
	RoleId          string        `json:"roleId"`

	/** 按钮ID */
	ButtonIds       []string      `json:"buttonIds"`
}
