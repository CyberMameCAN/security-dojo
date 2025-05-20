package auth

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
	"time"
)

// リフレッシュトークンの長さ
const refreshTokenLength = 32

// リフレッシュトークンの有効期限（例：7日間）
const refreshTokenExpiration = time.Hour * 24 * 7

type RefreshTokenData struct {
	UserID    uint
	ExpiresAt time.Time
}

// リフレッシュトークンを保存するmap（実際にはDBなどを使用する）
var refreshTokens = make(map[string]RefreshTokenData)
var refreshTokenMutex sync.Mutex

// リフレッシュトークンの生成
func GenerateRefreshToken(userID uint) (string, error) {
	// ランダムなバイト列を生成
	b := make([]byte, refreshTokenLength)
	// crypto/randを使用してセキュアなランダムバイトを生成
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	// バイト列をBase64エンコードしてトークンを生成
	// Base64エンコードを使用してURLセーフなトークンを生成
	// 生成されたトークンは、リフレッシュトークンとして使用される
	refreshToken := base64.URLEncoding.EncodeToString(b)
	// 有効期限を設定
	expiresAt := time.Now().Add(refreshTokenExpiration)

	// リフレッシュトークンとユーザーIDを保存
	// refreshTokensマップにリフレッシュトークンとユーザーIDを保存
	// リフレッシュトークンをキーとして、ユーザーIDと有効期限を保存
	// refreshTokenMutexを使用して排他制御を行う
	refreshTokenMutex.Lock()
	refreshTokens[refreshToken] = RefreshTokenData{UserID: userID, ExpiresAt: expiresAt}
	refreshTokenMutex.Unlock()

	return refreshToken, nil
}

// リフレッシュトークンの検証とユーザーIDの取得
func ValidateRefreshToken(oldRefreshToken string) (uint, string, bool, error) {
	var userID uint

	refreshTokenMutex.Lock()
	data, ok := refreshTokens[oldRefreshToken]
	if ok && time.Now().Before(data.ExpiresAt) {
		delete(refreshTokens, oldRefreshToken) // 古いリフレッシュトークンを無効化
		userID = data.UserID
		refreshTokenMutex.Unlock()

		// 新しいリフレッシュトークンを生成
		newToken, err := GenerateRefreshToken(userID)
		if err != nil {
			return 0, "", false, err
		}
		return data.UserID, newToken, true, nil
	}
	if ok {
		delete(refreshTokens, oldRefreshToken) // 有効期限切れのため削除
	}
	refreshTokenMutex.Unlock()
	return 0, "", false, nil
}
