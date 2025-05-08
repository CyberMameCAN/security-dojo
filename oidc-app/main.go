package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

var (
	clientID     = os.Getenv("CLIENT_ID")
	clientSecret = os.Getenv("CLIENT_SECRET")
	redirectURL  = os.Getenv("REDIRECT_URL")
	providerURL  = os.Getenv("PROVIDER_URL")
)

func main() {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, providerURL)
	if err != nil {
		log.Fatalf("failed to get provider: %v", err)
	}

	oidcConfig := &oidc.Config{
		ClientID: clientID,
	}

	verifier := provider.Verifier(oidcConfig)

	oauth2Config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		url := oauth2Config.AuthCodeURL("state-example")
		fmt.Println("Redirecting to:", url)
		c.Redirect(http.StatusFound, url)
	})

	router.GET("/callback", func(c *gin.Context) {
		if c.Query("state") != "state-example" {
			c.String(http.StatusBadRequest, "invalid state")
			return
		}

		token, err := oauth2Config.Exchange(ctx, c.Query("code"))
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("token exchange error: %v", err))
			return
		}

		rawIDToken, ok := token.Extra("id_token").(string)
		if !ok {
			c.String(http.StatusInternalServerError, "missing id_token")
			return
		}

		idToken, err := verifier.Verify(ctx, rawIDToken)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to verify ID Token: %v", err))
			return
		}

		var claims map[string]interface{}
		if err := idToken.Claims(&claims); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to parse claims: %v", err))
			return
		}

		c.JSON(http.StatusOK, claims)
	})

	router.Run(":8080")
}
