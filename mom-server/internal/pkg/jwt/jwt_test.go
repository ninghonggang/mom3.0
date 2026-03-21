package jwt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"mom-server/internal/config"
)

func TestJWT_GenerateToken(t *testing.T) {
	cfg := &config.JWTConfig{
		Secret:            "test-secret-key",
		AccessTokenExpire:  time.Hour,
		RefreshTokenExpire: time.Hour * 24 * 7,
	}
	j := New(cfg)

	userID := int64(1)
	tenantID := int64(1)
	username := "testuser"
	roles := []string{"admin", "user"}

	accessToken, refreshToken, err := j.GenerateToken(userID, tenantID, username, roles)

	assert.NoError(t, err)
	assert.NotEmpty(t, accessToken)
	assert.NotEmpty(t, refreshToken)
	assert.NotEqual(t, accessToken, refreshToken)
}

func TestJWT_ParseToken(t *testing.T) {
	cfg := &config.JWTConfig{
		Secret:            "test-secret-key",
		AccessTokenExpire:  time.Hour,
		RefreshTokenExpire: time.Hour * 24 * 7,
	}
	j := New(cfg)

	userID := int64(1)
	tenantID := int64(1)
	username := "testuser"
	roles := []string{"admin", "user"}

	accessToken, _, err := j.GenerateToken(userID, tenantID, username, roles)
	assert.NoError(t, err)

	claims, err := j.ParseToken(accessToken)

	assert.NoError(t, err)
	assert.Equal(t, userID, claims.UserID)
	assert.Equal(t, tenantID, claims.TenantID)
	assert.Equal(t, username, claims.Username)
	assert.Equal(t, roles, claims.Roles)
	assert.Equal(t, "mom-server", claims.Issuer)
}

func TestJWT_ParseToken_Invalid(t *testing.T) {
	cfg := &config.JWTConfig{
		Secret:            "test-secret-key",
		AccessTokenExpire:  time.Hour,
		RefreshTokenExpire: time.Hour * 24 * 7,
	}
	j := New(cfg)

	// Test with invalid token - should return an error
	_, err := j.ParseToken("invalid-token")
	assert.Error(t, err)
	// Malformed token returns ErrTokenMalformed
	assert.True(t, err == ErrTokenMalformed || err == ErrTokenInvalid)
}

func TestJWT_ParseToken_WrongSecret(t *testing.T) {
	cfg1 := &config.JWTConfig{
		Secret:            "secret-key-1",
		AccessTokenExpire:  time.Hour,
		RefreshTokenExpire: time.Hour * 24 * 7,
	}
	cfg2 := &config.JWTConfig{
		Secret:            "secret-key-2",
		AccessTokenExpire:  time.Hour,
		RefreshTokenExpire: time.Hour * 24 * 7,
	}

	j1 := New(cfg1)
	j2 := New(cfg2)

	// Generate token with j1
	accessToken, _, err := j1.GenerateToken(1, 1, "testuser", []string{"admin"})
	assert.NoError(t, err)

	// Try to parse with j2 (different secret)
	_, err = j2.ParseToken(accessToken)
	assert.Error(t, err)
}

func TestJWT_RefreshToken(t *testing.T) {
	cfg := &config.JWTConfig{
		Secret:            "test-secret-key",
		AccessTokenExpire:  time.Hour,
		RefreshTokenExpire: time.Hour * 24 * 7,
	}
	j := New(cfg)

	userID := int64(1)
	tenantID := int64(1)
	username := "testuser"
	roles := []string{"admin"}

	_, refreshToken, err := j.GenerateToken(userID, tenantID, username, roles)
	assert.NoError(t, err)

	newAccessToken, newRefreshToken, err := j.RefreshToken(refreshToken)

	assert.NoError(t, err)
	assert.NotEmpty(t, newAccessToken)
	assert.NotEmpty(t, newRefreshToken)

	// Verify new access token claims
	claims, _ := j.ParseToken(newAccessToken)
	assert.Equal(t, userID, claims.UserID)
	assert.Equal(t, tenantID, claims.TenantID)
	assert.Equal(t, username, claims.Username)
}

func TestJWT_RefreshToken_Invalid(t *testing.T) {
	cfg := &config.JWTConfig{
		Secret:            "test-secret-key",
		AccessTokenExpire:  time.Hour,
		RefreshTokenExpire: time.Hour * 24 * 7,
	}
	j := New(cfg)

	_, _, err := j.RefreshToken("invalid-refresh-token")
	assert.Error(t, err)
}

func TestClaims_Fields(t *testing.T) {
	cfg := &config.JWTConfig{
		Secret:            "test-secret-key",
		AccessTokenExpire:  time.Hour,
		RefreshTokenExpire: time.Hour * 24 * 7,
	}
	j := New(cfg)

	userID := int64(123)
	tenantID := int64(456)
	username := "adminuser"
	roles := []string{"admin", "manager"}

	accessToken, _, _ := j.GenerateToken(userID, tenantID, username, roles)
	claims, _ := j.ParseToken(accessToken)

	assert.Equal(t, userID, claims.UserID)
	assert.Equal(t, tenantID, claims.TenantID)
	assert.Equal(t, username, claims.Username)
	assert.Contains(t, claims.Roles, "admin")
	assert.Contains(t, claims.Roles, "manager")
	assert.NotNil(t, claims.ExpiresAt)
	assert.NotNil(t, claims.IssuedAt)
	assert.NotNil(t, claims.NotBefore)
}
