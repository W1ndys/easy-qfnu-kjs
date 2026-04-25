FROM golang:1.25-alpine AS builder
WORKDIR /app

ENV GOPROXY=https://goproxy.cn,direct
ENV GOSUMDB=sum.golang.google.cn

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /easy-qfnu-kjs .

FROM alpine:3.20
WORKDIR /app

RUN adduser -D app \
    && mkdir -p /app/data /app/logs \
    && chown -R app:app /app

COPY --from=builder /easy-qfnu-kjs /app/easy-qfnu-kjs

USER app

ENV GIN_MODE=release
ENV PORT=8080

EXPOSE 8080

ENTRYPOINT ["/app/easy-qfnu-kjs"]
