# ベースイメージを指定
FROM golang:1.23-alpine

# 作業ディレクトリを設定
WORKDIR /app

# Goモジュールを有効化
ENV GO111MODULE=on

# Goモジュールファイルをコピー
COPY go.mod .
COPY go.sum .

# 依存関係をダウンロード
RUN go mod download

# ソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o main .

# ポートを指定（アプリケーションで使用するポート）
EXPOSE 8080

# アプリケーションを実行
CMD ["./main"]
