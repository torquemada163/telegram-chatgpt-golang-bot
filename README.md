[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/torquemada163/telegram-chatgpt-golang-bot.svg)](https://github.com/torquemada163/telegram-chatgpt-golang-bot)
[![GitHub license](https://img.shields.io/github/license/torquemada163/telegram-chatgpt-golang-bot.svg)](https://github.com/torquemada163/telegram-chatgpt-golang-bot/blob/main/LICENSE)
[![Github all releases](https://img.shields.io/github/downloads/torquemada163/telegram-chatgpt-golang-bot/total.svg)](https://GitHub.com/torquemada163/telegram-chatgpt-golang-bot/releases/)
[![GitHub release](https://img.shields.io/github/release/torquemada163/telegram-chatgpt-golang-bot.svg)](https://GitHub.com/torquemada163/telegram-chatgpt-golang-bot/releases/)
[![GitHub forks](https://badgen.net/github/forks/torquemada163/telegram-chatgpt-golang-bot/)](https://GitHub.com/torquemada163/telegram-chatgpt-golang-bot/network/)
[![GitHub issues](https://img.shields.io/github/issues/torquemada163/telegram-chatgpt-golang-bot)](https://GitHub.com/torquemada163/telegram-chatgpt-golang-bot/issues/)

# Simply Telegram bot on Golang and ChatGPT

## Disclaimer
This is a simple, test program written for educational purposes. If something goes wrong - the author does not bear any responsibility.

This Telegram bot works on the principle of `long polling`, for real use it is better to switch to using a `webhook`.

## Install
Just a few simple steps:
1. Clone repository
2. Setup `c` variables (your ChatGPT token) in `app.go` with your value:
```golang
c := gogpt.NewClient("YOUR_CHATGPT_TOKEN")
```
3. Setup your Telegram bot token in `app.go` with your value:
```golang
bot, err := tgbotapi.NewBotAPI("YOUT_TELEGRAM_BOT_TOKEN_from_BotFather")
```
4. Setup external Golang packages (run this command in cloned project folder)
```
go get github.com/sashabaranov/go-gpt3
go get github.com/go-telegram-bot-api/telegram-bot-api
```
5. Then run with "go run" or compile the binary

## Usage
To send a ChatGPT request, send a command to the bot starting with **/cg** (can be changed in the aap.go file) followed by your question with a space. For example:

`/cg What is Golang?`

Remember that the request takes some time, so the answer will not come immediately. Typically within 10 seconds.
If for some reason the request to ChatGPT cannot be completed, the bot will send back the phrase:

`ChatGPT API error`

Happy using! ;-)
