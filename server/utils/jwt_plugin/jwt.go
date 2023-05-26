package jwt_plugin

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"

	"AirGo/global"
	timeTool "AirGo/utils/time"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.Server.JWT.SigningKey),
	}
}

type MyClaims struct {
	BaseClaims
	jwt.StandardClaims
}
type BaseClaims struct {
	// UUID        uuid.UUID
	ID       int
	UserName string
	// NickName    string
	// AuthorityId uint
}

func CreateClaims(baseClaims BaseClaims) MyClaims {
	ep, _ := timeTool.ParseDuration(global.Server.JWT.ExpiresTime)
	claims := MyClaims{
		BaseClaims: baseClaims,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,  // 签名生效时间
			ExpiresAt: time.Now().Add(ep).Unix(), // 过期时间 7天  配置文件
			Issuer:    global.Server.JWT.Issuer,  // 签名的发行者
		},
	}
	return claims
}

// 创建一个token
func CreateToken(claims MyClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(global.Server.JWT.SigningKey))
}

// 解析 token
func ParseToken(tokenString string) (*MyClaims, error) {
	signingKey := []byte(global.Server.JWT.SigningKey)
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return signingKey, nil
	})
	//log.Println("此时token err：", token, err)
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}

// 更新token
// func (j *JWT) RefreshToken(tokenString string) (string, error) {
//     jwt.TimeFunc = func() time.Time {
//         return time.Unix(0, 0)
//     }
//     token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
//         return j.SigningKey, nil
//     })
//     if err != nil {
//         return "", err
//     }
//     if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
//         jwt.TimeFunc = time.Now
//         claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
//         return j.CreateToken(*claims)
//     }
//     return "", TokenInvalid
// }
