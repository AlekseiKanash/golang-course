package bot

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/AlekseiKanash/golang-course/lesson_10/store/src/pgs"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func Try(token, app_token, channel_id string) { // Paste your channel id here
	client := slack.New(token, slack.OptionDebug(true), slack.OptionAppLevelToken(app_token))
	attachment := slack.Attachment{
		Pretext: "",
		Text:    "Hello from GoBot!",
	}

	channelId, timestamp, err := client.PostMessage(
		channel_id,
		slack.MsgOptionText("This is the main message", false),
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true),
	)

	if err != nil {
		log.Fatalf("%s\n", err)
	}

	log.Printf("Message successfully sent to Channel %s at %s\n", channelId, timestamp)

	socketClient := socketmode.New(
		client,
		socketmode.OptionDebug(true),
		// Option to set a custom logger
		socketmode.OptionLog(log.New(os.Stdout, "socketmode: ", log.Lshortfile|log.LstdFlags)),
	)

	// Create a context that can be used to cancel goroutine
	ctx, cancel := context.WithCancel(context.Background())
	// Make this cancel called properly in a real program , graceful shutdown etc
	defer cancel()

	go func(ctx context.Context, client *slack.Client, socketClient *socketmode.Client) {
		// Create a for loop that selects either the context cancellation or the events incomming
		for {
			select {
			// inscase context cancel is called exit the goroutine
			case <-ctx.Done():
				log.Println("Shutting down socketmode listener")
				return
			case event := <-socketClient.Events:
				// We have a new Events, let's type switch the event
				// Add more use cases here if you want to listen to other events.
				switch event.Type {

				case socketmode.EventTypeSlashCommand:
					// Just like before, type cast to the correct event type, this time a SlashEvent
					command, ok := event.Data.(slack.SlashCommand)
					if !ok {
						log.Printf("Could not type cast the message to a SlashCommand: %v\n", command)
						continue
					}
					// Dont forget to acknowledge the request
					socketClient.Ack(*event.Request)
					// handleSlashCommand will take care of the command
					err := handleSlashCommand(command, client)
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	}(ctx, client, socketClient)
	socketClient.Run()
}

// handleSlashCommand will take a slash command and route to the appropriate function
func handleSlashCommand(command slack.SlashCommand, client *slack.Client) error {
	// We need to switch depending on the command
	switch command.Command {
	case "/weather":
		// This was a hello command, so pass it along to the proper function
		return handleWeather(command, client)
	}
	return nil
}

func requestToWeb(api_url string) []byte {
	client := &http.Client{}

	req, err := http.NewRequest("GET", api_url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	return bodyBytes
}

// handleHelloCommand will take care of /hello submissions
func handleWeather(command slack.SlashCommand, client *slack.Client) error {
	// The Input is found in the text field so
	// Create the attachment and assigned based on the message
	attachment := slack.Attachment{}

	cmd := ""
	city := ""

	tokens := strings.Split(command.Text, " ")
	if len(tokens) >= 1 {
		cmd = tokens[0]
	}
	if len(tokens) == 2 {
		city = tokens[1]
	}

	switch cmd {
	case "fetch":
		{
			api_url := fmt.Sprintf("http://web_service/api/v0/fetch/%s", city)
			bodyBytes := requestToWeb(api_url)

			attachment.Text = fmt.Sprintf("Fetching weather for  %s", city)
			attachment.Pretext = fmt.Sprintf("Response %s", string(bodyBytes))
		}
	case "get":
		{
			api_url := fmt.Sprintf("http://web_service/api/v0/city/%s", city)
			bodyBytes := requestToWeb(api_url)
			weather := pgs.CityWeatherInfo{}
			json.Unmarshal(bodyBytes, &weather)
			attachment.Text = fmt.Sprintf("Temperature %.02f", weather.Temperature)
			attachment.Pretext = fmt.Sprintf("Weather for %s", weather.City)
		}
	case "list":
		{
			api_url := fmt.Sprintf("http://web_service/api/v0/list")
			bodyBytes := requestToWeb(api_url)

			var weathers []pgs.CityWeatherInfo
			json.Unmarshal(bodyBytes, &weathers)

			for _, weather := range weathers {
				fmt.Printf("%s\nCity: %s Temperature: %.02f", attachment.Text, weather.City, weather.Temperature)
				attachment.Text = fmt.Sprintf("%s\nCity: %s Temperature: %.02f", attachment.Text, weather.City, weather.Temperature)
			}
		}
	}

	attachment.Color = "#4af030"
	// Add Some default context like user who mentioned the bot
	attachment.Fields = []slack.AttachmentField{
		{
			Title: "Date",
			Value: time.Now().String(),
		}, {
			Title: "Initializer",
			Value: command.UserName,
		},
	}

	// Send the message to the channel
	// The Channel is available in the command.ChannelID
	_, _, err := client.PostMessage(command.ChannelID, slack.MsgOptionAttachments(attachment))
	if err != nil {
		return fmt.Errorf("failed to post message: %w", err)
	}
	return nil
}
