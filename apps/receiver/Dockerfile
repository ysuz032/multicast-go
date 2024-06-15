# マルチステージビルドのビルドステージ
FROM golang:1.22.4 AS builder

# アプリケーションディレクトリの作成
WORKDIR /app

# ソースコードをコピー
COPY src/ /app

# Goアプリケーションをビルド
RUN go build -o receiver cmd/main.go

# 実行用ステージ
FROM golang:1.22.4

# 必要なファイルをコピー
COPY --from=builder /app/receiver /usr/local/bin/receiver

# 実行可能ファイルを設定
CMD ["/usr/local/bin/receiver"]

# デフォルトの環境変数
ENV MULTICAST_ADDR="224.0.0.1"
ENV MULTICAST_PORT="28000"