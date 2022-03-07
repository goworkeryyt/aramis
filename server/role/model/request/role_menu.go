/**
 * @Time: 2022/3/7 18:34
 * @Author: yt.yin
 */

package rolereq

// RoleMenuBindRequest 角色菜单绑定请求参数
type RoleMenuBindRequest struct {

	/** 角色ID */
	RoleId           string           `json:"roleId"`

	/** 菜单ID列表 */
	MenuIds          []string         `json:"menuIds"`

	/** 按钮ID列表 */
	ButtonIds        []string         `json:"buttonIds"`
}
