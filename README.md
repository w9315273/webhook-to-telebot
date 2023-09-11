# webhook-to-telebot
# 使用说明
这是一个基于 Go 的 Webhook 服务, 它可以接收 HTTP 请求并将其转发到 Telegram bot

---

## 目录
1. [项目概述](#项目概述)
2. [运行容器](#运行容器)
3. [环境变量](#环境变量)
4. [请求格式](#请求格式)
5. [安全性建议](#安全性建议)
6. [常见问题与解决方法](#常见问题与解决方法)

---

## 项目概述
当该服务接收到一个合法的 HTTP 请求时, 它会解析请求主体中的消息部分, 并将其转发到指定的 Telegram bot 聊天中

---

## 运行容器
### Docker
```
docker run -d \
--name=webhook-to-telebot \
--restart=always \
-P 5000:5000 \
-e PORT=5000 \
-e TEXT_COUNT=10 \
-e WEBHOOK_PATH='/webhook'
-e BOT_TOKEN='YOUR_BOT_TOKEN'
-e CHAT_ID='YOUR_CHAT_ID'
-e AUTH_TOKEN='YOUR_AUTH_TOKEN'
yesmydark/webhook-to-telebot:latest
```
### 环境变量
`PORT` 服务监听的端口,默认为 `5000`

`TEXT_COUNT` 指定应提取的文本部分的数量, 默认为10条 `text1` `text2` `text3`......`text10`

`WEBHOOK_PATH` 服务的路由路径, 例如 `/webhook` 请求地址为 `localhost:5000/webhook` `example.com:5000/webhook`

`BOT_TOKEN` 你的 `Telegram bot api token`

`CHAT_ID` 你的 `Telegram chat ID` 消息将被发送到这里

`AUTH_TOKEN` 自定义用于验证 HTTP 请求的令牌

---

## 请求格式
请求应该使用 JSON 格式发送

内容类型 / HTTP 标头为 `application/json`

另一个标头为 `Authorization: 你设置的令牌`

HTTP 请求主体应包含 `text1` `text2` 等键 例如
```
{
  "text1": "Hello World",
  "text2": "消息2",
  "text3": "消息3"
}
```

每个键对应于要发送的消息的一部分

这些部分将被合并, 并作为一条消息发送到 Telegram bot

Telegram bot 将会向你发送以下内容
```
Hello World
消息2
消息3
```

---

## 安全性建议
不要将 `AUTH_TOKEN` `BOT_TOKEN` 或任何其他敏感数据公开或与他人共享

始终确保你的服务在防火墙之后运行, 并只允许受信任的来源访问

使用 HTTPS 来增强通信的安全性 比如使用 Nginx Caddy 等反向代理添加 HTTPS 支持

---

## 常见问题与解决方法
问题: 我收到了 `Request too large` 错误, 怎么办?

答: 确保你发送的请求主体不超过 `10KB` （目前设定 `10KB` ）

问题: 我收到了 `401` 响应，这是什么原因？

答: 确保你在发送请求时提供了正确的 `AUTH_TOKEN` 例如: `Authorization: password`

问题: 我如何知道我的 `Telegram chat ID` ？

答: 可以通过向 `@userinfobot` 发送消息来获取 `Telegram chat ID`
