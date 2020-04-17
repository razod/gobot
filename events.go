package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if strings.HasPrefix(m.Content, Prefix+"ping") {
		pingCMD(s, m)
	} else if strings.HasPrefix(m.Content, Prefix+"pong") {
		pongCMD(s, m)
	}
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateStatus(0, "g!ping")
}
