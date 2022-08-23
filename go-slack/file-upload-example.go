package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArray := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"test.txt", "test2.txt"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArray,
			File:     fileArr[i],
		}

		file, err := api.UploadFile(params)

		fmt.Printf("your token %v\n", api)

		if err != nil {
			fmt.Printf("We found an error: %s\n", err)
			return
		}

		fmt.Printf("Your file : %s was upload\n", file.Name)
	}
}

/* NOTES:
- Create an app in slack > api.slack.com/apps
- Called it "file-bot"
- enable socket mode, give it a name
- give it permission, go to OAuth & Permissions, add a new OAuth Scope and select:
 -- channels: read
 -- chat:write
 -- file:read and write
 -- im:read and write
 -- remote files read and share
- Install on your workspace
- copy your bot user OAuth Token...set as an environment var
- get your channel ID... set as an environment var
- add the bot to your channel
- execute: go run file-upload-example.go
*/
