package main

import "github.com/bwmarrin/discordgo"

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

