package registry

import gptbot "github.com/green-api/telegram-chatgpt-go"

var gptHelperInstance *gptbot.TelegramGptBot

func RegisterGptHelper(instance *gptbot.TelegramGptBot) {
	gptHelperInstance = instance
}

func GetGptHelper() *gptbot.TelegramGptBot {
	return gptHelperInstance
}
