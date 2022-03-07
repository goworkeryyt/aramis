/**
 * @Time: 2022/3/7 20:30
 * @Author: yt.yin
 */

package userapi

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"github.com/gookit/color"
	"github.com/goworkeryyt/aramis/server/user/model"
	"github.com/goworkeryyt/aramis/server/user/model/request"
	"github.com/goworkeryyt/aramis/server/user/model/response"
	"github.com/goworkeryyt/go-core/global"
	"github.com/goworkeryyt/go-core/jwt"
	"github.com/goworkeryyt/go-toolbox/result"
	"github.com/goworkeryyt/go-toolbox/uuid"
	"github.com/goworkeryyt/go-toolbox/validator"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"time"
)

type LoginApi struct{}

// 当开启多服务器部署时，替换下面的配置，使用redis共享存储验证码
// var store = captcha.NewDefaultRedisStore()

var store = base64Captcha.DefaultMemStore

// Captcha 验证码
func (l *LoginApi) Captcha(c *gin.Context) {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80)
	// cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // v8下使用redis
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		global.LOG.Error("验证码获取失败:", zap.Any("err", err))
		result.FailMsg("验证码获取失败:"+err.Error(), c)
	} else {
		result.OkDataMsg(gin.H{"captchaId": id, "picPath": b64s}, "验证码获取成功", c)
	}
	return
}

// Login 登录
func (l *LoginApi) Login(c *gin.Context) {
	var lReq usereq.LoginRequest
	// 绑定Json
	err := c.ShouldBindJSON(&lReq)
	if err != nil {
		result.FailMsg("绑定的Json格式有误", c)
		return
	}

	// 参数格式效验
	if err = validator.Verify(lReq, usermod.LoginVerify); err != nil {
		result.FailMsg(err.Error(), c)
		return
	}

	// 验证码正确，则进入密码效验，否则直接返回前端
	if store.Verify(lReq.CaptchaId, lReq.Captcha, true) {
		u := &usermod.User{Username: lReq.Username, Password: lReq.Password}
		if err, user := userService.Login(u); err != nil {
			global.LOG.Error("登陆失败! 用户名不存在或者密码错误:", zap.Any("err", err))
			result.FailMsg("用户名不存在或者密码错误("+err.Error()+")", c)
		} else {
			// 获取用户具有的角色
			roleIds, _ := userRoleService.GetUserRoles(user.ID)
			roles, _ := roleService.FindRoleByIds(roleIds)
			// 角色赋值
			user.Roles = roles
			// 签发token
			l.tokenNext(c, *user)
		}
	} else {
		result.FailMsg("验证码错误", c)
	}
}

// ChangePassword 修改密码
func (l *LoginApi) ChangePassword(c *gin.Context) {
	var user usereq.ChangePasswordRequest
	// json绑定
	err := c.ShouldBindJSON(&user)
	if err != nil {
		result.FailMsg("绑定的Json格式有误", c)
		return
	}

	// 校验密码
	if err = validator.Verify(user, usermod.ChangePasswordVerify); err != nil {
		result.FailMsg(err.Error(), c)
		return
	}

	u := &usermod.User{Username: user.Username, Password: user.Password}
	if err, _ := userService.ChangePassword(u, user.NewPassword); err != nil {
		global.LOG.Error("修改失败:", zap.Any("err", err))
		result.FailMsg("修改失败，原密码与当前账户不符", c)
	} else {
		result.OkMsg("修改成功", c)
	}
}

// tokenNext 登录以后签发jwt
func (l *LoginApi) tokenNext(c *gin.Context, user usermod.User) {
	var roleIdsStr string
	for i, role := range user.Roles {
		if i == len(user.Roles)-1 {
			roleIdsStr = roleIdsStr + role.ID
		} else {
			roleIdsStr = roleIdsStr + role.ID + ","
		}
	}

	// 唯一签名
	j := &jwt.JWT{SigningKey: []byte(global.CONFIG.JWT.SigningKey)}
	// 过期时间 单位秒
	expiresAt := time.Now().Unix() + global.CONFIG.JWT.ExpiresTime
	claims := jwt.CustomClaims{
		TokenId:       uuid.UUID(),
		UserId:        user.ID,
		Username:      user.Username,
		AuthorityId:   roleIdsStr,
		MerchantNo:    user.MerchantNo,
		Phone:         user.Phone,
		UserType:      user.UserType,
		BufferTime:    global.CONFIG.JWT.BufferTime,
		// 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		RegisteredClaims:jwt.RegisteredClaims("yt.yin",expiresAt),
	}

	token, err := j.CreateToken(claims)
	if err != nil {
		color.Danger.Println(err)
		global.LOG.Error("获取token失败!", zap.Any("err", err))
		result.FailMsg("获取token失败", c)
		return
	}

	// 登录成功，记录登录时间，登录次数+1
	user.LoginCount = user.LoginCount + 1
	user.LoginTime = carbon.Now().ToDateTimeString()
	err = global.DB.Updates(&user).Error
	if err != nil {
		global.LOG.Error("登录失败!", zap.Any("err", err))
		result.FailMsg("登录失败:"+err.Error(), c)
		return
	}
	result.OkDataMsg(response.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: expiresAt * 1000,
	}, "登录成功", c)
	return
}

// GetUserInfo 获取用户自身信息
func (l *LoginApi) GetUserInfo(c *gin.Context) {
	id := jwt.GetUserID(c)
	if ReqUser, err := userService.GetUserInfo(id); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		result.FailMsg("获取失败:"+err.Error(), c)
	} else {
		result.OkDataMsg(gin.H{"userInfo": ReqUser}, "获取成功", c)
	}
}

