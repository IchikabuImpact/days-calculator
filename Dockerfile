# ベースイメージ
FROM golang:1.24-alpine

# 必要なパッケージをインストール
RUN apk add --no-cache gcc musl-dev

# 作業ディレクトリ設定
WORKDIR /app

# ソースコードと環境ファイルをコピー
COPY . .

# 必要な依存関係をインストール
RUN go mod tidy

# ビルド
RUN go build -o days_calculator ./app/days_calculator.go

# ポート公開 (デフォルトの8080)
EXPOSE 8089

# アプリケーション実行
CMD ["/app/days_calculator"]

