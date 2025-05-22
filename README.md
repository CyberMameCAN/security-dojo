# 情報処理安全確保支援士を学習中

IPAの情報処理安全確保支援士を勉強するために手を動かして理解を試みる記録

## このリポジトリについて

情報処理安全確保支援士（登録セキスペ）の試験範囲や実務で必要となるセキュリティ技術を、実際に手を動かして学ぶためのサンプル集です。

## ディレクトリ構成

- `IDToken/` : PythonでIDトークン（JWT）のデコードや検証の仕組みを学ぶサンプル
- `PKCE/` : PythonでPKCE（Proof Key for Code Exchange）の仕組みを確認するサンプル
- `Solt/` : ソルト・ペッパーを使ったハッシュ化のサンプル（Python）
- `jwt-refresh-token-sample/` : Go言語でJWTアクセストークン・リフレッシュトークンの発行・認証サンプル
- `oidc-app/` : Go言語でOpenID Connect認証を体験するサンプル（Google認証対応、Docker構成）

## 使用技術

- Python 3.x
- Go 1.23+
- Docker / docker-compose
- Gin（GoのWebフレームワーク）
- dotenv（Pythonの環境変数管理）

## 参考リンク

- [IPA 情報処理安全確保支援士試験](https://www.ipa.go.jp/shiken/kubun/r4sc.html)
- [RFC7636 (PKCE)](https://datatracker.ietf.org/doc/html/rfc7636)
- [OpenID Connect](https://openid.net/connect/)

## 今後の予定・メモ

- サンプルの追加や解説の充実を予定
- IssueやPR歓迎

## ライセンス

このリポジトリは [MIT License](./LICENSE) のもとで公開されています。