# First stage: build environment
FROM golang:1.21.5 AS builder

WORKDIR /srv/grpc
COPY . .

# ARGでビルド時の引数を定義
ARG DSN
ARG DB_USER_NAME
ARG DB_PASSWORD
ARG DB_ADDR
ARG DB_NAME
ARG FRONTEND_URL
ARG PORT

# .env ファイルを生成
RUN echo "DSN=$DSN" > .env \
    && echo "DB_USER_NAME=$DB_USER_NAME" >> .env \
    && echo "DB_PASSWORD=$DB_PASSWORD" >> .env \
    && echo "DB_ADDR=$DB_ADDR" >> .env \
    && echo "DB_NAME=$DB_NAME" >> .env \
    && echo "FRONTEND_URL=$FRONTEND_URL" >> .env \
    && echo "PORT=$PORT" >> .env

RUN go mod download

ARG VERS="3.11.4"
ARG ARCH="linux-x86_64"

RUN CGO_ENABLED=0 GOOS=linux \
    go build -a -installsuffix cgo \
    -o /go/bin/server \
    github.com/KinjiKawaguchi/Coteacher/coteacher-server/cmd/server

# Final stage: runtime environment
# CA証明書が必要なので、alpineをベースにする
# Final stage: runtime environment
FROM alpine:latest

# 必要なパッケージをインストール
RUN apk --no-cache add ca-certificates

# アプリケーション実行用のディレクトリを設定
WORKDIR /app

# ビルドステージからバイナリファイルをコピー
COPY --from=builder /go/bin/server /app/server

# .envファイルを/appディレクトリにコピー
COPY --from=builder /srv/grpc/.env /app/.env

# アプリケーションの起動
ENTRYPOINT ["/app/server"]
