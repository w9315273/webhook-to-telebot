使用说明
项目概述
此Go应用程序充当一个Webhook服务，可以接收HTTP请求并将其转发到Telegram bot。请求主体应包含要发送到Telegram的消息。

环境变量设置
为了使此应用程序工作，您需要设置以下环境变量：

BOT_TOKEN: 您的Telegram bot令牌。
PORT: 服务应该在其上运行的端口（默认为5000）。
WEBHOOK_PATH: 服务的路由路径（例如，/webhook）。
CHAT_ID: Telegram聊天的ID，应发送消息到这里。
AUTH_TOKEN: 验证请求的令牌。
TEXT_COUNT: 可选的环境变量，指定应提取多少条文本部分（默认为20）。
请求格式
请求应以JSON格式发送，内容类型为application/json。JSON应该包含text1，text2等键，每个键对应于要发送的消息的一部分。这些部分将被合并并作为单个消息发送到Telegram。

json
Copy code
{
    "text1": "Hello,",
    "text2": "World!"
}
该消息将发送到Telegram为：“Hello,\nWorld!”

请求认证
为了增强安全性，您应该在请求头中添加Authorization字段，其值应与AUTH_TOKEN环境变量相匹配。

启动服务
一旦你设置了所有的环境变量，你可以运行这个程序。默认情况下，它将在端口5000上运行，但你可以通过PORT环境变量更改这个设置。

每当有一个请求发送到WEBHOOK_PATH时，它都会被处理，并且消息会被发送到指定的Telegram聊天。

错误处理
以下是可能的HTTP响应状态码及其含义：

200 OK: 消息成功发送到Telegram。
401 Unauthorized: 请求的Authorization头不匹配AUTH_TOKEN。
415 Unsupported Media Type: 请求的内容类型不是application/json。
413 Request Entity Too Large: 请求体超过了10KB。
500 Internal Server Error: 在处理请求或发送消息到Telegram时发生了错误。
