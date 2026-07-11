FROM golang:1.26.5-alpine3.24 AS builder
WORKDIR /app
COPY . .
RUN go run main.go
WORKDIR /app/output

FROM nginx:1.31.2
COPY --from=builder /app/output /usr/share/nginx/html
EXPOSE 80
