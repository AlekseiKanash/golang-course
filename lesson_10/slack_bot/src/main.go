package main

import (
	"fmt"
	"os"

	"github.com/AlekseiKanash/golang-course/lesson_10/slack_bot/src/bot"
)

func main() {
	//godotenv.Load("/Users/akanash/.slack/.env")

	token := os.Getenv("SLACK_AUTH_TOKEN")
	channel_id := os.Getenv("SLACK_CHANNEL_ID")
	app_token := os.Getenv("WEB_SOCKET_TOKEN")

	fmt.Println(token)
	fmt.Println(channel_id)
	fmt.Println(app_token)
	bot.Try(token, app_token, channel_id)
}
