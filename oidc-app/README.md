# OpenID Connect

Golang 1.23+（Alpineベース）で**OpenID Connect**のハンズオン環境を構築するためのセットアップ手順です。

## 構成概要

- OpenID Provider：Google。[Google Cloud Console](https://console.cloud.google.com/apis/credentials)から**Google OAuth2 クライアント**を作る。
- クライアント：Golang + Ginで実装
- 認証後、ユーザー情報を表示(JSON)
- docker-compose + Dockerfile で構築

## まずはじめに

下記手順でgo.sum作って、コンテナを作成・起動する。

go.modに以下を記入する。

    go 1.23

    require (
        github.com/coreos/go-oidc/v3 v3.7.0
        github.com/gin-gonic/gin v1.10.0
        golang.org/x/oauth2 v0.15.0
    )

### go.sumを作る

    go mod tidy

### .env

.envを作成し、

    mv env.sample .env

以下のパラメータを**Google Cloud Console**で作った情報と合わせる。

- CLIENT_ID
- CLIENT_SECRET
- EDIRECT_URL


### コンテナ作成と起動

    docker compose up --build

## HTTPアクセス

以下にアクセスし、Google認証を試す。

    http://localhost:8080

JSONが返ってくる。**"email_verified":true**になってれば多分大丈夫。