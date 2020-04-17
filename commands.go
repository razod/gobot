package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

/* COMMAND TEMPLATE
func nameCMD(s *discordgo.Session, m *discordgo.MessageCreate) {
	// code goes here
}
*/

func pingCMD(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, ":ping_pong: Pong!")
}

func pongCMD(s *discordgo.Session, m *discordgo.MessageCreate) {
	embed := &discordgo.MessageEmbed{
		Description: ":ping_pong: Ping",
		Color:       0x00ff00,
		Author:      &discordgo.MessageEmbedAuthor{},
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}

func airhornCMD(s *discordgo.Session, m *discordgo.MessageCreate) {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		return
	}
	g, err := s.State.Guild(c.GuildID)
	if err != nil {
		return
	}
	for _, vs := range g.VoiceStates {
		if vs.UserID == m.Author.ID {
			err = playSound(s, g.ID, vs.ChannelID)
			if err != nil {
				fmt.Println("Error playing sound:", err)
			}

			return
		}
	}
}
