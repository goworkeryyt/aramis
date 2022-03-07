/**
 * @Time: 2022/3/4 15:28
 * @Author: yt.yin
 */

package usereq

// LoginRequest 登录参数提交
type LoginRequest struct {

	/** 用户名 */
	Username        string         `json:"username"`

	/** 密码 */
	Password        string         `json:"password"`

	/** 验证码 */
	Captcha         string         `json:"captcha"`

	/** 验证码ID */
	CaptchaId       string         `json:"captchaId"`
}

