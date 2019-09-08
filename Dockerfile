FROM golang:1.13 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org go build -o app .

FROM scratch
COPY --from=builder /app/app /app
ENTRYPOINT ["/app"]