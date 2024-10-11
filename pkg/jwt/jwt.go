package jwt

import (
	"errors"
	"fisco-cert-app/setting"
	"time"

	"go.uber.org/zap"

	"github.com/dgrijalva/jwt-go"
)

var (
	tokenSecret = []byte("冬日暖暖")
)

type MyClaims struct {
	UserID   uint64 `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成access JwtToken
func GenToken(userID uint64, username string) (string, error) {
	c := MyClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(
				time.Duration(setting.Conf.AuthConfig.TokenExpire) * time.Hour).Unix(),
			Issuer: "userInner",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(tokenSecret)
}

func GenBothToken(userID uint64, username string) (aToken, RToken string, err error) {
	c := MyClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(
				time.Duration(setting.Conf.AuthConfig.TokenExpire) * time.Hour).Unix(),
			Issuer: "userInner",
		},
	}

	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(tokenSecret)
	if err != nil {
		return
	}

	RToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(
			time.Duration(time.Hour * 24 * 7)).Unix(),
		Issuer: "userInner",
	}).SignedString(tokenSecret)

	return
}

// ParseToken 解析Access Token
func ParseToken(tokenStr string) (*MyClaims, error) {
	mc := &MyClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, mc, func(token *jwt.Token) (interface{}, error) {
		return tokenSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")
}

// ParseRefreshToken 解析Refresh Token
func ParseRefreshToken(AToken, RToken string) (newAToken, newRToken string, err error) {
	// 判断Refresh Token 是否有效
	_, err = jwt.Parse(RToken, func(token *jwt.Token) (interface{}, error) {
		return tokenSecret, nil
	})
	if err != nil {
		return
	}
	// 判断Access Token 是否有效
	mc := &MyClaims{}
	_, err = jwt.ParseWithClaims(AToken, mc, func(token *jwt.Token) (interface{}, error) {
		return tokenSecret, nil
	})
	if err != nil {
		zap.L().Error("jwt.ParseRefreshToken failed", zap.Error(err))
	}
	v, _ := err.(*jwt.ValidationError)

	// Refresh Token有效，再次签发Access Token
	if err == nil || v.Errors == jwt.ValidationErrorExpired {
		return GenBothToken(mc.UserID, mc.Username)
	}
	return
}
