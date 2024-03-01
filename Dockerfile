# First stage: build environment
FROM golang:1.21.5 AS builder

WORKDIR /srv/grpc
COPY /coteacher-server ./coteacher-server
COPY /proto-gen ./proto-gen
RUN cd ./coteacher-server && go mod download

ARG VERS="3.11.4"
ARG ARCH="linux-x86_64"

RUN cd ./coteacher-server && CGO_ENABLED=0 GOOS=linux \
    go build -a -installsuffix cgo \
    -o /go/bin/server \
    github.com/KinjiKawaguchi/Coteacher/coteacher-server/cmd/server

# Final stage: runtime environment
# CA証明書が必要なので、alpineをベースにする
FROM alpine:latest

# 必要なパッケージをインストール
RUN apk --no-cache add ca-certificates

# ビルドステージからバイナリと.envファイルをコピー
COPY --from=builder /go/bin/server /server
COPY --from=builder /srv/grpc/coteacher-server/.env /

# アプリケーションの起動
ENTRYPOINT ["/server"]