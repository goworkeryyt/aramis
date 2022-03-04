/**
 * @Time: 2022/3/4 15:43
 * @Author: yt.yin
 */

package request

// UpdateUserStatusReq 更新用户状态请求参数
type UpdateUserStatusReq struct {

	/** 用户账号记录ID */
	UserId          string      `json:"id"`

	/** 用户账号状态 */
	Status          string      `json:"status"`
}


// ChangePasswordRequest 修改密码参数提交
type ChangePasswordRequest struct {

	/** 用户名 */
	Username        string      `json:"username"`

	/** 密码 */
	Password        string      `json:"password"`

	/** 新密码 */
	NewPassword     string      `json:"newPassword"`
}

// CreateUserRequest 新建用户
type CreateUserRequest struct {

	/** 用户账号记录ID */
	ID              string      `json:"id"`

	/** 用户名 */
	Username        string      `json:"username"`

	/** 用户昵称 */
	NickName        string      `json:"nickName"`

	/** 用户头像 */
	HeaderImg       string      `json:"headerImg"`

	/** 密码 */
	Password        string      `json:"password"`

	/** 用户类型(管理员，商户，普通用户) */
	UserType        string      `json:"userType"`

	/** 电话 */
	Phone           string      `json:"phone"`

	/** 邮件 */
	Email           string      `json:"email"`

	/** 商户号 */
	Status          string      `json:"status"`

	/** 用户角色ID  */
	RoleIds       []string      `json:"roleIds"`

	/** 商户号 */
	MerchantNo      string      `json:"merchantNo"`

	/** 创建者类型（从Token中获取） */
	CreatorType     string

	/** 创建者ID（从Token中获取） */
	CreatorId       string
}

// UpdateUserRequest 更新用户请求参数
type UpdateUserRequest struct {

	/** 用户账号记录ID */
	ID              string      `json:"id"`

	/** 用户昵称 */
	NickName        string      `json:"nickName"`

	/** 用户头像 */
	HeaderImg       string      `json:"headerImg"`

	/** 密码 */
	Password        string      `json:"password"`

	/** 电话 */
	Phone           string      `json:"phone"`

	/** 邮件 */
	Email           string      `json:"email"`

	/** 商户号 */
	Status          string      `json:"status"`

	/** 用户角色ID */
	RoleIds       []string      `json:"roleIds"`
}

// UpdateUserRolesReq 更新用户角色请求参数
type UpdateUserRolesReq struct {

	/** 用户账号记录ID */
	UserId          string      `json:"id"`

	/** 用户角色ID列表 */
	RoleIds       []string      `json:"roleIds"`
}