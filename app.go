package main

import (
	"context"
	"log"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	gogpt "github.com/sashabaranov/go-gpt3"
)

func sendChatGPT(c *gogpt.Client, sendText string) string {
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:            gogpt.GPT3TextDavinci003,
		MaxTokens:        2048,
		Prompt:           sendText,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
	}

	resp, errChat := c.CreateCompletion(ctx, req)
	if errChat != nil {
		return "ChatGPT API error"
	} else {
		return resp.Choices[0].Text
	}
}

func main() {
	// Chat GPT initialization
	c := gogpt.NewClient("YOUR_CHATGPT_TOKEN")

	// Telegram initialization
	bot, err := tgbotapi.NewBotAPI("YOUT_TELEGRAM_BOT_TOKEN_from_BotFather")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true // set to false for suppress logs in stdout
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Start Telegram long polling update
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)

	//Check message in updates
	for update := range updates {
		if update.Message == nil {
			continue
		}

		// If message present and not start from '/cg ' - ignore message
		if !strings.HasPrefix(update.Message.Text, "/cg ") {
			continue
		}

		// Cut text prefix '/cg '
		cutText, _ := strings.CutPrefix(update.Message.Text, "/cg ")

		// Send request to ChatGPT
		update.Message.Text = sendChatGPT(c, cutText)

		// Send message to Telegram
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		_, err = bot.Send(msg)
		if err != nil {
			log.Println("Error:", err)
		}
	}
}
