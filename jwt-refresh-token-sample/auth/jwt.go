package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 秘密鍵
var jwtSecret = []byte("your-secret-key") // 実際には環境変数などで管理する

// JWTのペイロードの構造体
type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// JWTの生成
func GenerateToken(userID uint, username string) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24) // トークンの有効期限

	// JWTのペイロードを作成
	// ユーザーIDとユーザー名をペイロードに含める
	claims := &Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// JWTを生成
	// HS256アルゴリズムを使用して署名
	// claimsをペイロードとして使用
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 署名を生成
	// jwtSecretを使用して署名
	// 署名されたトークンを生成
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// JWTの検証
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}
	return claims, nil
}
