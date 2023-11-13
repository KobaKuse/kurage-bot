package main

import (
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

const (
	URL_PATTERN     string = `(?i)https://(twitter\.com|x\.com)`
	FIX_TWITTER_URL string = "https://fxtwitter.com"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}
	token := os.Getenv("BOT_TOKEN")
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println(err)
	}
	discord.AddHandler(onReady)
	discord.AddHandler(onEvents)

	err = discord.Open()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Listening...")

	stopBot := make(chan os.Signal, 1)

	signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	<-stopBot
	err = discord.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func onReady(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateWatchStatus(0, "克巴巴")
}

func onEvents(s *discordgo.Session, m interface{}) {
	if mCreate, isMessageCreate := m.(*discordgo.MessageCreate); isMessageCreate {
		handleMessage(s, mCreate.Message)
	} else if mUpdate, isMessageUpdate := m.(*discordgo.MessageUpdate); isMessageUpdate {
		handleMessage(s, mUpdate.Message)
	}
}

func handleMessage(s *discordgo.Session, m *discordgo.Message) {
	if m != nil && m.Author != nil {
		if m.Author.ID != s.State.User.ID && regexp.MustCompile(URL_PATTERN).MatchString(m.Content) {
			convertedMsg := convertUrl(m.Content)
			s.ChannelMessageSend(m.ChannelID, convertedMsg+" 來自"+m.Author.Mention())
			s.ChannelMessageDelete(m.ChannelID, m.ID)
		}
	}
}

func convertUrl(msg string) string {
	return regexp.MustCompile(URL_PATTERN).ReplaceAllString(msg, FIX_TWITTER_URL)
}
