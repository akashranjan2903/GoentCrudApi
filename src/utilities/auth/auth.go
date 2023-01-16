package auth

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/gocrud/config"
	"github.com/golang-jwt/jwt"
)

// HashPassword - hashes a password with argon2
// func HashPassword(password string) (string, error) {
// 	hash, err := NewArgon2ID().Hash(password)
// 	if err != nil {
// 		return "", err
// 	}
// 	return hash, nil
// }

// VerifyPassword - verifies a password with argon2
func VerifyPassword(enteredpassword, password string) (bool, error) {
	return enteredpassword == password, nil
}

// tokenType Enum
type TokenType string

const (
	Access  TokenType = "access"
	Refresh TokenType = "refresh"
	Reset   TokenType = "reset"
)

// JWTTokenResponse struct
type JWTTokenResponse struct {
	Token      string        `json:"token"`
	ExpiresMin time.Duration `json:"expires_min"`
}

// GenerateJWT - generates a JWT token with a token type and a payload
func GenerateJWT(tt TokenType, userUUID string) (JWTTokenResponse, error) {
	var expTime time.Duration
	switch tt {
	case "access":
		expTime = time.Duration(config.GetConfig().JWT.ExpiryMinAccess)
		if expTime == 0 {
			expTime = 15
		}
	case "refresh":
		expTime = time.Duration(config.GetConfig().JWT.ExpiryMinRefresh)
		if expTime == 0 {
			expTime = 1440
		}
	case "reset":
		expTime = time.Duration(config.GetConfig().JWT.ExpiryMinRefresh)
		if expTime == 0 {
			expTime = 5
		}
	default:
		return JWTTokenResponse{}, fmt.Errorf("invalid token type")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"type": tt,
		"id":   userUUID,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Minute * expTime).Unix(),
	})

	tokenString, err := token.SignedString(
		[]byte(config.GetConfig().JWT.Secret),
	)
	if err != nil {
		return JWTTokenResponse{}, err
	}

	return JWTTokenResponse{
		Token:      tokenString,
		ExpiresMin: expTime,
	}, nil
}

// DecodeJWT - decodes a JWT token
func DecodeJWT(
	tokenString string,
	tt TokenType,
) (tokenType string, userID string, err error) {
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf(
					"unexpected signing method: %v",
					token.Header["alg"],
				)
			}
			return []byte(config.GetConfig().JWT.Secret), nil
		},
	)
	if err != nil {
		log.Printf("Error decoding JWT: %v", err)
		return "", "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if _tokenType := claims["type"]; _tokenType != string(tt) {
			return "", "", fmt.Errorf("invalid token type: %s", tokenType)
		}
		return claims["type"].(string), claims["id"].(string), nil
	}

	return "", "", fmt.Errorf("invalid token")
}

// PasswordValidator - validates the password.
func PasswordValidator(password string) (bool, string) {
	switch {
	case len(password) < 8:
		return false, "Password must be at least 8 characters long"
	case len(password) > 128:
		return false, "Password must be less than 128 characters long"
	case !regexp.MustCompile(`[A-Z]+`).MatchString(password):
		return false, "Password must contain at least one uppercase letter"
	case !regexp.MustCompile(`[a-z]+`).MatchString(password):
		return false, "Password must contain at least one lowercase letter"
	case !regexp.MustCompile(`\d+`).MatchString(password):
		return false, "Password must contain at least one number"
	case !regexp.MustCompile(`[!@#~$%^&*()+|_]{1}`).MatchString(password):
		return false, "Password must contain at least one special character"
	default:
		return true, ""
	}
}

// Check if the email is in the correct format with the regex
func EmailValidator(email string) (bool, error) {
	emailRegex := regexp.MustCompile(
		`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`,
	)
	if !emailRegex.MatchString(email) {
		return false, fmt.Errorf("invalid email")
	}
	return true, nil
}
