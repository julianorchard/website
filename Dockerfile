FROM golang:1.26.5-alpine3.24 AS builder

WORKDIR /app
COPY . .
RUN go run main.go

FROM nginx:1.31.2
COPY --from=builder /app/output /usr/share/nginx/html

LABEL org.opencontainers.image.authors="Julian Orchard <jorchard@pm.me>"
LABEL org.opencontainers.image.source=https://github.com/julianorchard/website

EXPOSE 80
