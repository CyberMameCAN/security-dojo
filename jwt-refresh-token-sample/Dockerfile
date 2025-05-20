FROM golang:1.23-alpine

WORKDIR /app

# アプリケーションのコピー
COPY . .

RUN apk upgrade --update && \
    apk add --no-cache alpine-sdk build-base
RUN go install github.com/air-verse/air@latest

# 依存関係のインストール
COPY go.mod go.sum ./
RUN go mod download

# アプリケーションのビルド
RUN go build -o main .

# ポート公開
EXPOSE 8080

# アプリケーション実行
# CMD ["./main"]
CMD ["air", "-c", ".air.toml"]