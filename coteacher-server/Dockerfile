# ビルドステージ
FROM golang:1.21 AS builder

# プロジェクトのルートディレクトリに作業ディレクトリを設定
WORKDIR /app

# プロトコル生成ファイルをコピー
COPY ./proto-gen ./proto-gen

# go.mod と go.sum をコピー
COPY go.mod go.sum ./
# 依存関係をダウンロード
RUN go mod download

# ソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -tags timetzdata -o coteacher-backend ./cmd/server

# 実行ステージ
FROM ubuntu:22.04

# 必要なパッケージをインストール
RUN apt-get update && apt-get install -y --no-install-recommends tzdata ca-certificates && rm -rf /var/lib/apt/lists/*

# タイムゾーン設定
ENV TZ Asia/Tokyo

# ワーキングディレクトリを設定
WORKDIR /app

# ビルドステージからバイナリをコピー
COPY --from=builder /app/coteacher-backend .

# コンテナ起動時にバイナリを実行
ENTRYPOINT ["/app/coteacher-backend"]
