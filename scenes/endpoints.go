package scenes

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	greenapi "github.com/green-api/telegram-api-client-golang"
	chatbot "github.com/green-api/telegram-chatbot-golang"
	"github.com/green-api/telegram-demo-chatbot-golang/util"
)

type EndpointsScene struct{}

func (s EndpointsScene) Start(bot *chatbot.Bot) {

	bot.IncomingMessageHandler(func(message *chatbot.Notification) {
		if !util.IsSessionExpired(message) {
			lang := message.GetStateData()["lang"].(string)
			text, _ := message.Text()
			senderName := ""
			if sd, ok := message.Body["senderData"].(map[string]interface{}); ok {
				if sn, ok := sd["senderName"].(string); ok {
					senderName = sn
				}
			}
			senderId, _ := message.Sender()
			botNumber := ""
			if id, ok := message.Body["instanceData"].(map[string]interface{}); ok {
				if wid, ok := id["wid"].(string); ok {
					botNumber = wid
				}
			}

			menuBtnText := util.GetString([]string{"menu_button", lang})
			stopBtnText := util.GetString([]string{"stop_button", lang})

			switch text {
			case "1":
				message.SendText(util.GetString([]string{"send_text_message", lang}) + util.GetString([]string{"links", lang, "send_text_documentation"}))
				return
			case "2":
				message.SendUrlFile(
					"https://storage.yandexcloud.net/sw-prod-03-test/ChatBot/corgi.pdf",
					"corgi.pdf",
					util.GetString([]string{"send_file_message", lang})+util.GetString([]string{"links", lang, "send_file_documentation"}))
				return
			case "3":
				message.SendUrlFile(
					"https://storage.yandexcloud.net/sw-prod-03-test/ChatBot/corgi.jpg",
					"corgi.jpg",
					util.GetString([]string{"send_image_message", lang})+util.GetString([]string{"links", lang, "send_file_documentation"}))
				return
			case "4":
				message.SendText(util.GetString([]string{"send_audio_message", lang}) + util.GetString([]string{"links", lang, "send_file_documentation"}))
				var fileLink = "https://storage.yandexcloud.net/sw-prod-03-test/ChatBot/Audio_bot_eng.mp3"
				if lang == "ru" {
					fileLink = "https://storage.yandexcloud.net/sw-prod-03-test/ChatBot/Audio_bot.mp3"
				}
				message.SendUrlFile(fileLink, "audio.mp3", "")
				return
			case "5":
				var fileLink = "https://storage.yandexcloud.net/sw-prod-03-test/ChatBot/Video_bot_eng.mp4"
				if lang == "ru" {
					fileLink = "https://storage.yandexcloud.net/sw-prod-03-test/ChatBot/Video_bot_ru.mp4"
				}
				message.SendUrlFile(fileLink, "video.mp4",
					util.GetString([]string{"send_video_message", lang})+util.GetString([]string{"links", lang, "send_file_documentation"}))
				return
			case "6":
				message.SendText(util.GetString([]string{"send_contact_message", lang}) + util.GetString([]string{"links", lang, "send_contact_documentation"}))
				phoneStrSender := strings.ReplaceAll(senderId, "@c.us", "")
				phoneIntSender, _ := strconv.Atoi(phoneStrSender)
				message.SendContact(greenapi.Contact{PhoneContact: phoneIntSender, FirstName: senderName})
				return
			case "7":
				message.SendText(util.GetString([]string{"send_location_message", lang}) + util.GetString([]string{"links", lang, "send_location_documentation"}))
				message.SendLocation("", "", 35.888171, 14.440230)
				return
			case "9":
				message.SendText(util.GetString([]string{"get_avatar_message", lang}) + util.GetString([]string{"links", lang, "get_avatar_documentation"}))
				resp, _ := message.Service().GetAvatar(senderId)
				var avatar map[string]interface{}
				_ = json.Unmarshal(resp.Body, &avatar)

				if avatarURL, ok := avatar["urlAvatar"].(string); ok && avatarURL != "" {
					message.SendUrlFile(
						avatarURL,
						"avatar.jpg",
						util.GetString([]string{"avatar_found", lang}))
				} else {
					message.SendText(util.GetString([]string{"avatar_not_found", lang}))
				}
				return
			case "10":
				message.SendText(util.GetString([]string{"send_link_message_preview", lang}) + util.GetString([]string{"links", lang, "send_link_documentation"}))
				message.SendText(util.GetString([]string{"send_link_message_no_preview", lang}) + util.GetString([]string{"links", lang, "send_link_documentation"}))
				return
			case "11":
				message.SendText(util.GetString([]string{"add_to_contact", lang}))
				botPhoneStr := strings.ReplaceAll(botNumber, "@c.us", "")
				botPhoneInt, _ := strconv.Atoi(botPhoneStr)
				message.SendContact(greenapi.Contact{PhoneContact: botPhoneInt, FirstName: util.GetString([]string{"bot_name", lang})})
				message.ActivateNextScene(CreateGroupScene{})
				return
			case "12":
				message.AnswerWithText(util.GetString([]string{"send_quoted_message", lang}) + util.GetString([]string{"links", lang, "send_quoted_message_documentation"}))
				return
			case "13":
				message.SendUrlFile("https://raw.githubusercontent.com/green-api/telegram-demo-chatbot-golang/refs/heads/master/assets/about_go.jpg", "logo.jpg",
					util.GetString([]string{"about_go_chatbot", lang})+
						util.GetString([]string{"link_to_docs", lang})+
						util.GetString([]string{"links", lang, "chatbot_documentation"})+
						util.GetString([]string{"link_to_source_code", lang})+
						util.GetString([]string{"links", lang, "chatbot_source_code"})+
						util.GetString([]string{"link_to_green_api", lang})+
						util.GetString([]string{"links", lang, "greenapi_website"})+
						util.GetString([]string{"link_to_console", lang})+
						util.GetString([]string{"links", lang, "greenapi_console"})+
						util.GetString([]string{"link_to_youtube", lang})+
						util.GetString([]string{"links", lang, "youtube_channel"}))
				return
			case "стоп", "Стоп", "stop", "Stop", "0", stopBtnText:
				message.SendText(util.GetString([]string{"stop_message", lang}) + "*" + senderName + "*!")
				message.ActivateNextScene(StartScene{})
				return
			case "menu", "меню", "Menu", "Меню", menuBtnText:
				menuContent := util.GetString([]string{"menu", lang})

				var welcomeFileURL string
				if lang == "en" || lang == "es" || lang == "he" {
					welcomeFileURL = "https://raw.githubusercontent.com/green-api/telegram-demo-chatbot-golang/refs/heads/master/assets/welcome_en.jpg"
				} else {
					welcomeFileURL = "https://raw.githubusercontent.com/green-api/telegram-demo-chatbot-golang/refs/heads/master/assets/welcome_ru.jpg"
				}
				message.SendUrlFile(welcomeFileURL, "welcome.jpg", menuContent)
				return
			case "":
			default:
				message.SendText(util.GetString([]string{"not_recognized_message", lang}))
			}
		} else {
			senderId, _ := message.Sender()
			log.Printf("Session expired or missing for user %s. Redirecting to StartScene.", strings.ReplaceAll(senderId, "@c.us", ""))
			message.ActivateNextScene(StartScene{})
			message.SendText(util.GetString([]string{"select_language"}))
		}
	})
}
