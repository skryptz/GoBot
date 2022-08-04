package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	Token string = "MTAwMjI0NDQzMDc2MzQwOTQ1OA.Gr_ZUm.loOkTxu15weoxYroob8M-ogRnwCIoh7REyL4DA"
)

// func init() {

// 	flag.StringVar(&Token, "t", "", "Bot Token")
// 	flag.Parse()
// }

func main() {
	fmt.Println("Discord Bot Written In Golang!")

	session, err := discordgo.New(Token)
	checkNilErr(err)

	// Intents
	session.Identify.Intents |= discordgo.IntentMessageContent
	session.Identify.Intents |= discordgo.IntentsAll
	session.Identify.Intents |= discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates

	// Callbacks
	session.AddHandler(messageCreate)
	session.AddHandler(guildCreate)

	err = session.Open()
	checkNilErr(err)

	// var messageOne *discordgo.MessageCreate
	// messageOne.ChannelID = "977059104617021462"

	// session.ChannelMessageSend(messageOne.ChannelID, "I'm Alive!")

	defer session.Close()

	fmt.Println("Bot running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

}

func checkNilErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		fmt.Println("Bot wrote something?")
		return
	}

	if m.Content == "ping" {
		fmt.Println("recieved ping")
		s.ChannelMessageSend(m.ChannelID, "pong")
	}

	if m.Content == "pong" {
		fmt.Println("recieved pong")
		s.ChannelMessageSend(m.ChannelID, "ping")
	}
}

func guildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {

	if event.Guild.Unavailable {
		return
	}

	for _, channel := range event.Guild.Channels {
		if channel.ID == event.Guild.ID {
			_, _ = s.ChannelMessageSend(channel.ID, "Airhorn is ready! Type !airhorn while in a voice channel to play a sound.")
			return
		}
	}
}
