package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// If the message is "ping" reply with "Pong!"
	if m.Content == "NSFW" {
		s.ChannelMessageSend(m.ChannelID, "https://matias.ma/nsfw/")
	}

	if m.Content == "Zeta" {
		s.ChannelMessageSend(m.ChannelID, "https://scontent.fsjo8-1.fna.fbcdn.net/v/t39.30808-6/347231621_797580248245202_8437773414200328318_n.jpg?_nc_cat=107&cb=99be929b-59f725be&ccb=1-7&_nc_sid=09cbfe&_nc_ohc=W0ZkkmbTeQUAX_OmVCj&_nc_ht=scontent.fsjo8-1.fna&oh=00_AfDWDRVBqKunvDE4Lhi_J5dcryRTubp_odksayvj0xHBnQ&oe=64A950BB")
	}

	if m.Content == "Lorenzo" {
		s.ChannelMessageSend(m.ChannelID, "https://scontent.fsjo8-1.fna.fbcdn.net/v/t1.6435-9/132549785_1193010267780186_3316927142384384506_n.jpg?_nc_cat=100&cb=99be929b-59f725be&ccb=1-7&_nc_sid=09cbfe&_nc_ohc=9GMp8QgsdxMAX_jlaTi&_nc_ht=scontent.fsjo8-1.fna&oh=00_AfB4juDt75dPu-9WOd7ZyvlNYmIYSN9eOwx0oaD7y_WhJQ&oe=64CB2DF8")
	}

	if m.Content == "whoami" {
		name := m.Author.Username
		s.ChannelMessageSend(m.ChannelID, "You are "+name)
	}
}
