package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/k0kubun/pp"
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	DiscordToken string
}

func main() {
	c := config{}
	if err := envconfig.Process("roomkeeper", &c); err != nil {
		log.Fatalf("failed to parse config: %s", err)
	}
	pp.Println(c)
	dg, err := discordgo.New("Bot " + c.DiscordToken)
	if err != nil {
		log.Fatalf("failed to init discord: %s", err)
	}

	dg.AddHandler(onMessage)

	err = dg.Open()
	if err != nil {
		log.Fatalf("failed to open connection: %s", err)
	}
	defer dg.Close()

	log.Println("Start roomkeeper")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func onMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore myself
	if m.Author.ID == s.State.User.ID {
		return
	}

	s.ChannelTyping(m.ChannelID)
	return
}
