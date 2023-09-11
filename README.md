# webhook-to-telebot

# 使用说明

这是一个基于Go的Webhook服务, 它可以接收HTTP请求并将其转发到Telegram bot

---

## 目录

1. [项目概述](#项目概述)
2. [运行容器](#运行容器)
3. [环境变量](#环境变量)
4. [请求格式](#请求格式)

---

## 项目概述

当该服务接收到一个合法的HTTP请求时, 它会解析请求主体中的消息部分, 并将其转发到指定的Telegram聊天中

---

## 运行容器

### Docker

```bash
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
```bash
'PORT': 服务监听的端口,默认为5000
'TEXT_COUNT': 指定应提取的文本部分的数量, 默认为10条（text1,text2,text3...text10）
'WEBHOOK_PATH': 服务的路由路径, 例如/webhook （localhost:5000/webhook）（https://example.com/webhook）
'BOT_TOKEN': 你的Telegram bot api token
'CHAT_ID': 你的Telegram chat ID, 消息将被发送到这里
'AUTH_TOKEN': 自定义用于验证HTTP请求的令牌
```
---

## 请求格式
