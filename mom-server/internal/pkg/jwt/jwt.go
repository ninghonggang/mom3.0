package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"mom-server/internal/config"
)

var (
	ErrTokenExpired     = errors.New("token已过期")
	ErrTokenNotValidYet = errors.New("token未生效")
	ErrTokenInvalid     = errors.New("token无效")
	ErrTokenMalformed   = errors.New("token格式错误")
)

// Claims JWT Claims
type Claims struct {
	UserID   int64    `json:"user_id"`
	TenantID int64    `json:"tenant_id"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
	jwt.RegisteredClaims
}

// JWT JWT工具
type JWT struct {
	secret             []byte
	accessTokenExpire  time.Duration
	refreshTokenExpire time.Duration
}

// New 创建JWT实例
func New(cfg *config.JWTConfig) *JWT {
	return &JWT{
		secret:             []byte(cfg.Secret),
		accessTokenExpire: cfg.AccessTokenExpire,
		refreshTokenExpire: cfg.RefreshTokenExpire,
	}
}

// GenerateToken 生成Token
func (j *JWT) GenerateToken(userID, tenantID int64, username string, roles []string) (string, string, error) {
	// Access Token
	accessClaims := Claims{
		UserID:   userID,
		TenantID: tenantID,
		Username: username,
		Roles:    roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.accessTokenExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "mom-server",
		},
	}
	accessTokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenStr, err := accessTokenStruct.SignedString(j.secret)
	if err != nil {
		return "", "", err
	}

	// Refresh Token
	refreshClaims := Claims{
		UserID:   userID,
		TenantID: tenantID,
		Username: username,
		Roles:    roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.refreshTokenExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "mom-server",
		},
	}
	refreshTokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenStr, err := refreshTokenStruct.SignedString(j.secret)
	if err != nil {
		return "", "", err
	}

	return accessTokenStr, refreshTokenStr, nil
}

// ParseToken 解析Token
func (j *JWT) ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.secret, nil
	})

	if err != nil {
		return nil, handleJWTError(err)
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}

func handleJWTError(err error) error {
	if errors.Is(err, jwt.ErrTokenExpired) {
		return ErrTokenExpired
	}
	if errors.Is(err, jwt.ErrTokenNotValidYet) {
		return ErrTokenNotValidYet
	}
	if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
		return ErrTokenInvalid
	}
	if errors.Is(err, jwt.ErrTokenMalformed) {
		return ErrTokenMalformed
	}
	return ErrTokenInvalid
}

// RefreshToken 刷新Token
func (j *JWT) RefreshToken(refreshTokenStr string) (string, string, error) {
	claims, err := j.ParseToken(refreshTokenStr)
	if err != nil {
		return "", "", err
	}

	return j.GenerateToken(claims.UserID, claims.TenantID, claims.Username, claims.Roles)
}
