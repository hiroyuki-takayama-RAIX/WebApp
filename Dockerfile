# デプロイ用コンテナに含めるバイナリを作成するコンテナ
FROM golang:1.20.4-alpine as deploy-builder

WORKDIR /app

# COPY src/go.mod .
# COPY src/go.sum .
# RUN go mod download
# why `go mod download` causes error in "docker compose build"…?

COPY . .
RUN go build ./src

# デプロイ用のコンテナ
FROM debian:bullseye-slim as deploy

RUN apt-get update

COPY --from=deploy-builder /app .

CMD ["./WebApp"]