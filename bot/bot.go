package bot

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

var BOtID string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	BOtID = user.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BOtID {
		return
	}

	if m.Content == "spidy" {
		_, err := s.ChannelMessageSend(m.ChannelID, "lol Spidy Noob! 🤣🤣")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

}
