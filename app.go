package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	gogpt "github.com/sashabaranov/go-gpt3"
	"github.com/spf13/viper"
)

type Config struct {
	TelegramToken string `mapstructure:"tgToken"`
	GptToken      string `mapstructure:"gptToken"`
}

func LoadConfig(path string) (c Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)
	return
}

func sendChatGPT(c *gogpt.Client, sendText string) string {
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:            gogpt.GPT3TextDavinci003,
		MaxTokens:        2048,
		Prompt:           sendText,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
	}

	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return "ChatGPT API error"
	} else {
		return resp.Choices[0].Text
	}
}

func main() {
	// Reading config.yaml
	config, err := LoadConfig(".")

	if err != nil {
		panic(fmt.Errorf("fatal error with config.yaml: %w", err))
	}

	// Chat GPT initialization
	chatGPT := gogpt.NewClient(config.GptToken)

	// Telegram initialization
	bot, err := tgbotapi.NewBotAPI(config.TelegramToken)
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

		// Bot is typing...
		action := tgbotapi.NewChatAction(update.Message.Chat.ID, tgbotapi.ChatTyping)
		_, err := bot.Send(action)
		if err != nil {
			log.Println("Error:", err)
		}

		// Send request to ChatGPT
		update.Message.Text = sendChatGPT(chatGPT, cutText)

		// Send message to Telegram
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		_, err = bot.Send(msg)
		if err != nil {
			log.Println("Error:", err)
		}
	}
}
