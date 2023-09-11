FROM golang:1.21.1 AS builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

RUN cp -r /etc/ssl/certs /tmp/certs

FROM scratch

ENV PORT=5000
ENV WEBHOOK_PATH=/webhook
ENV BOT_TOKEN=YOUR_BOT_TOKEN
ENV CHAT_ID=YOUR_CHAT_ID
ENV AUTH_TOKEN=YOUR_AUTH_TOKEN

COPY --from=builder /app/main /app/main
COPY --from=builder /tmp/certs /etc/ssl/certs

EXPOSE 5000

CMD ["/app/main"]