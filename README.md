# Golang Telegram Bot Project Template

Telegram Template to get you started with your telegram bot projects

For the golang bot api library docs click <b>[Here](https://pkg.go.dev/github.com/go-telegram-bot-api/telegram-bot-api?utm_source=godoc) </b>

# Prerequisites

Make sure you you have <a href="https://golang.org/dl/">Golang</a> installed on your machine<br/>
Telegram bot Token from [BotFather](https://core.telegram.org/bots#6-botfather)
**NB** You **must** set bot token in **config.json** <br />

# Installation(Local)

**NB** You **must** set bot token in **config.json** <br />

1.Clone this repository using <code>git clone https://github.com/mashaole/telegramtemplate.git</code><br/>
2.Run `go mod tidy` & `go mod vendor` to install required packages and add a vendor to your repo
3.Run <code>go install</code><br/>
4.run <code>go run main.go</code><br/>

# Deploy (Google App Engine)

**NB** You **must** set bot token in **config.json** <br />

1.Clone this repository using <code>git clone https://github.com/mashaole/telegramtemplate.git</code> on <b> gcp project terminal </b><code>Cloud Shell</code><br/>
2.Run `go mod tidy` & `go mod vendor` to install required packages and add a vendor to your repo<br/>
3.Run <code>go install</code><br/>
4.Run <code>go build</code><br/>
5.Run <code>gcloud app deploy</code><br/>
6.Open terminal and perform curl to Telegram messages to your bot <code>curl --data "url=(Custom Url)/(handled route)" https://api.telegram.org/bot(botToken)/SetWebhook</code> <br/>
`E.g curl --data "url=https://mywbesite.com//bsjhdbsakjbdjks" https://api.telegram.org/bot1234567890:kdjakjsdlksajdkl-sldlakslk-aklsjhdkjha/SetWebhook`<br/>

# Usage

<b>[telegram.go](https://github.com/Celbux/telegram-template/blob/MashReview/src/template-service/telegram.go) </b>

- Telegram is initiated using <code>InitTelegram()</code> and theres a implentation for <code>local environemnt</code> and <code>server environement</code>(uses web hooks)<br/>
- The <code>deleteChatHistory()</code> funtion is to clear screen when there is a new update.<br/>
- The <code>TelegramHandler()</code> function handles updates from inline or simple keyboard.<br/>
- The <code>SendSimpleMessage()</code> function handles sending simple message.<br/>
- The <code>SendInlineMessage()</code> function handles sending Inline message.<br/>
- The <code>EditMessage()</code> function handles editing message.<br/>
- The <code>CreateSimpleKeyBoard()</code> returns Simlple Keyboard.<br/>
- The <code>CreateInlineKeyBoard()</code> returns Inline Keyboard.<br/>
