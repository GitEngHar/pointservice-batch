# 1) ビルド用ステージ
FROM golang:1.23 AS builder
ARG JOB_PATH=./cmd/healthCheck/main.go
WORKDIR /app

# Go モジュールのキャッシュを利用
COPY go.mod go.sum ./
RUN go mod download

# ソースをコピーしてバイナリをコンパイル
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o main ${JOB_PATH}

# 2) ランタイム用ステージ
FROM public.ecr.aws/lambda/go:1

# ビルドしたバイナリを Lambda のデフォルトパスに配置
COPY --from=builder /app/main /var/task/

EXPOSE 1324

# デフォルトのエントリーポイントは CMD ["main"] なので省略可
# CMD ["main"]