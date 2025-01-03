FROM golang:1.23-alpine AS build

ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app
COPY . .
RUN go mod download;go build -o /app/hello

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/hello .

CMD ["./hello"]
