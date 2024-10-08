package jwt

import (
	"backend/pkg/config"
	"github.com/dgrijalva/jwt-go"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const JwtName = "Authorization"

type Claims struct {
	Id     uint `json:"id"`
	RoleId int  `json:"roleId"`
	jwt.StandardClaims
}

// MakeToken 生成 jwt 令牌
func MakeToken(id, roleId int, jwtConf config.Auth) (string, error) {
	// 过期时间
	expTime := time.Now().Add(time.Duration(jwtConf.AccessExpire) * time.Hour)
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Id:     uint(id),
		RoleId: roleId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	})
	signedString, err := tokenClaim.SignedString([]byte(jwtConf.AccessSecret))
	if err != nil {
		logrus.Errorf("生成jwt出错 : %+v", errors.WithStack(err))
	}
	return signedString, err
}

// ParseToken 解析 jwt 令牌
func ParseToken(token string, jwtConf config.Auth) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtConf.AccessSecret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	logrus.Errorf("解析jwt出错 : %+v", errors.WithStack(err))
	return nil, err
}

// GetToken 各种方法获取 token
func GetToken(c *gin.Context) (string, error) {
	if token := c.GetHeader(JwtName); token != "" {
		return token, nil
	}

	if token, _ := c.Cookie(JwtName); token != "" {
		return token, nil
	}
	return "", errors.New("没有找到" + JwtName)
}

// GetAndParseToken 对 GetToken 和 ParseToken 的封装
func GetAndParseToken(c *gin.Context, jwtConf config.Auth) (*Claims, error) {
	token, err := GetToken(c)
	if err != nil {
		return nil, err
	}
	return ParseToken(token, jwtConf)
}
