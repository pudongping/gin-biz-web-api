// jwt 认证
package jwt

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	jwtPkg "github.com/golang-jwt/jwt"

	"gin-biz-web-api/pkg/app"
	"gin-biz-web-api/pkg/config"
	"gin-biz-web-api/pkg/logger"
)

var (
	ErrTokenExpired           = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh = errors.New("令牌已过最大刷新时间")
	ErrTokenMalformed         = errors.New("请求令牌格式有误")
	ErrTokenInvalid           = errors.New("请求令牌无效")
	ErrTokenNotFound          = errors.New("无法找到令牌")
)

// JWT 定义一个 jwt 对象
type JWT struct {

	// 密钥，用以加密 JWT，从配置文件 config/jwt.go 中读取
	Key []byte

	// 刷新 token 的最大过期时间
	MaxRefresh time.Duration
}

// JWTCustomClaims 自定义载荷
type JWTCustomClaims struct {
	UserID       string `json:"user_id"`     // 当前登录的用户 id
	ExpireAtTime int64  `json:"expire_time"` // 过期时间

	// StandardClaims 结构体实现了 Claims 接口继承了  Valid() 方法
	// JWT 规定了7个官方字段，提供使用:
	// - iss (issuer)：发布者
	// - sub (subject)：主题
	// - iat (Issued At)：生成签名的时间
	// - exp (expiration time)：签名过期时间
	// - aud (audience)：观众，相当于接受者
	// - nbf (Not Before)：生效时间
	// - jti (JWT ID)：编号
	jwtPkg.StandardClaims
}

func NewJWT() *JWT {
	return &JWT{
		Key:        []byte(config.GetString("jwt.key")),                                  // 密钥
		MaxRefresh: time.Duration(config.GetInt64("jwt.max_refresh_time")) * time.Minute, // 允许刷新时间
	}
}

// ParseToken 解析 token
func (j *JWT) ParseToken(c *gin.Context, userToken ...string) (*JWTCustomClaims, error) {
	var (
		tokenStr string
		err      error
	)

	if len(userToken) > 0 {
		tokenStr = userToken[0]
	} else {
		// 获取 token
		tokenStr, err = j.GetToken(c)
		if err != nil {
			return nil, err
		}
	}

	// 解析用户 token
	token, err := j.parseTokenString(tokenStr)

	// 解析出错时
	if err != nil {
		validationErr, ok := err.(*jwtPkg.ValidationError)
		if ok {
			switch validationErr.Errors {
			case jwtPkg.ValidationErrorMalformed:
				return nil, ErrTokenMalformed
			case jwtPkg.ValidationErrorExpired:
				return nil, ErrTokenExpired
			}
		}
		return nil, ErrTokenInvalid
	}

	// 将 token 中的 claims 信息解析出来和 JWTCustomClaims 数据结构进行校验
	// Valid 验证基于时间的声明，例如：过期时间（ExpiresAt）、签发者（Issuer）、生效时间（Not Before），
	// 需要注意的是，如果没有任何声明在令牌中，仍然会被认为是有效的
	if claims, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}

// GetTTL 计算出 token 还剩多少秒后过期
func (j *JWT) GetTTL(c *gin.Context, userToken ...string) (int64, error) {

	claims, err := j.ParseToken(c, userToken...)

	if err != nil {
		// 此时已经过期，或者出现 token 解析失败
		return 0, err
	}

	// 此时的 token 一定是没有过期的，否则上一步 ParseToken 就已经报错了
	ttl := claims.ExpiresAt - app.TimeNowInTimezone().Unix()

	return ttl, nil
}

// RefreshToken 刷新 token
func (j *JWT) RefreshToken(c *gin.Context) (string, error) {
	// 获取 token
	tokenStr, err := j.GetToken(c)
	if err != nil {
		return "", err
	}

	// 解析用户 token
	token, err := j.parseTokenString(tokenStr)

	// 解析出错时（未报错证明是合法的 token 或者未到过期时间）
	if err != nil {
		validationErr, ok := err.(*jwtPkg.ValidationError)
		// 如果满足刷新 token 的条件，就继续往下走下一步（只要是单一的 ValidationErrorExpired 报错就认为是）
		if !ok || validationErr.Errors != jwtPkg.ValidationErrorExpired {
			return "", err
		}
	}

	// 解析出自定义的载荷信息 JWTCustomClaims
	claims := token.Claims.(*JWTCustomClaims)

	// 检查是否过了【最大允许刷新的时间】
	// 首次签名时间 + 最大允许刷新时间区间 > 当前时间 ====> 首次签名时间 > 当前时间 - 最大允许刷新时间区间
	if claims.IssuedAt > app.TimeNowInTimezone().Add(-j.MaxRefresh).Unix() {
		// 此时并没有过最大允许刷新时间，因此可以重新颁发 token
		claims.StandardClaims.ExpiresAt = j.expireAtTime()
		return j.createToken(*claims)
	}

	// 当前时间过了最大允许刷新的时间
	return "", ErrTokenExpiredMaxRefresh
}

// GenerateToken 生成 token
func (j *JWT) GenerateToken(userId string) string {
	// 构造用户 claims 信息（负荷）
	expireAtTime := j.expireAtTime()
	claims := JWTCustomClaims{
		UserID:       userId,
		ExpireAtTime: expireAtTime,
		StandardClaims: jwtPkg.StandardClaims{
			NotBefore: app.TimeNowInTimezone().Unix(), // 签名生效时间
			IssuedAt:  app.TimeNowInTimezone().Unix(), // 首次签名时间（后续刷新 token 不会更新）
			ExpiresAt: expireAtTime,                   // 签名过期时间
			Issuer:    config.GetString("cfg.app.name"),   // 签名颁发者
		},
	}

	// 根据 claims 生成 token
	token, err := j.createToken(claims)
	if err != nil {
		logger.LogErrorIf(err)
		return ""
	}

	return token
}

// createToken 创建 token，用于内部调用
func (j *JWT) createToken(claims JWTCustomClaims) (string, error) {
	// 使用 HS256 算法生成的 token
	tokenClaims := jwtPkg.NewWithClaims(jwtPkg.SigningMethodHS256, claims)
	// 生成签名字符串
	return tokenClaims.SignedString(j.Key)
}

// expireAtTime 获取过期时间点
func (j *JWT) expireAtTime() int64 {
	timeNow := app.TimeNowInTimezone() // 获取当前时区的时间

	var expireTime int64

	if app.IsLocal() {
		// 调试模式时，使用调试模式的过期时间
		expireTime = config.GetInt64("jwt.local_expire_time")
	} else {
		expireTime = config.GetInt64("jwt.expire_time")
	}

	expire := time.Duration(expireTime) * time.Minute
	// 返回加过期时间区间后的时间点
	return timeNow.Add(expire).Unix()
}

// parseTokenString 使用 jwtpkg.ParseWithClaims 解析 Token
func (j *JWT) parseTokenString(tokenStr string) (*jwtPkg.Token, error) {
	// ParseWithClaims 用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回 *jwtPkg.Token
	return jwtPkg.ParseWithClaims(tokenStr, &JWTCustomClaims{}, func(token *jwtPkg.Token) (interface{}, error) {
		return j.Key, nil
	})
}

// GetToken 获取请求中的 token 参数
func (j *JWT) GetToken(c *gin.Context) (string, error) {
	var token string

	if query, exists := c.GetQuery("token"); exists && "" != query {
		token = query
	} else if post, exists := c.GetPostForm("token"); exists && "" != post {
		token = post
	} else {
		token = c.GetHeader("token")
	}

	if "" == token {
		return "", ErrTokenNotFound
	}

	return token, nil
}
