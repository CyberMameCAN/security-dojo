package main

import (
	"fmt"
	"jwt-refresh-token-sample/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// ログインとアクセストークン、リフレッシュトークンの発行
	r.POST("/login", func(c *gin.Context) {
		// 実際にはここでユーザー認証を行う
		userID := uint(123)
		username := "testuser"

		// ユーザー認証が成功した場合、アクセストークンとリフレッシュトークンを生成
		accessToken, err := auth.GenerateToken(userID, username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "アクセストークンの生成に失敗しました"})
			return
		}

		// リフレッシュトークンは、ユーザーIDを元に生成
		// ここでは簡略化のため、ユーザーIDをそのまま使用していますが、実際にはセキュアな方法で生成する必要があります
		// 例えば、UUIDやランダムな文字列を使用することが推奨されます
		refreshToken, err := auth.GenerateRefreshToken(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "リフレッシュトークンの生成に失敗しました"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"access_token": accessToken, "refresh_token": refreshToken})
	})

	// アクセストークンのリフレッシュ（リフレッシュトークンローテーション）
	r.POST("/refresh", func(c *gin.Context) {
		oldRefreshToken := c.PostForm("refresh_token")
		if oldRefreshToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "リフレッシュトークンが必要です"})
			return
		}

		userID, newRefreshToken, ok, err := auth.ValidateRefreshToken(oldRefreshToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "リフレッシュトークンの検証に失敗しました"})
			return
		}
		fmt.Println("新しいリフレッシュトークン:", newRefreshToken)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "無効なリフレッシュトークンです"})
			return
		}
		fmt.Println("ユーザーID:", userID)

		// 新しいアクセストークンを発行
		username := "testuser" // 実際にはユーザーIDからユーザー情報を取得する
		newAccessToken, err := auth.GenerateToken(userID, username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "新しいアクセストークンの生成に失敗しました"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"access_token": newAccessToken, "refresh_token": newRefreshToken})
	})

	// JWT検証のエンドポイント（ミドルウェアで保護する例）
	authMiddleware := func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "認証が必要です"})
			return
		}
		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "無効なトークンです"})
			return
		}
		c.Set("claims", claims) // 検証済みのクレームをコンテキストに保存
		c.Next()
	}

	r.GET("/protected", authMiddleware, func(c *gin.Context) {
		claims, _ := c.Get("claims")
		c.JSON(http.StatusOK, gin.H{"message": "認証成功", "user": claims.(*auth.Claims)})
	})

	r.Run(":8080")
}
