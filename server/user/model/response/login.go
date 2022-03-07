/**
 * @Time: 2022/3/4 18:42
 * @Author: yt.yin
 */

package response

import  "github.com/goworkeryyt/aramis/server/user/model"

// LoginResponse 登录返回结果
type LoginResponse struct {

	/** 用户信息 */
	User             usermod.User       `json:"user"`

	/** 用户登陆的 token */
	Token            string             `json:"token"`

	/** toekn失效时间 */
	ExpiresAt        int64              `json:"expiresAt"`
}
