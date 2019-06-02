package cagliostro

import (
	"errors"
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const (
	killSayings = 3
	pastas      = 0
)

// BUGS
// TODO: Currently check can be circumvented by simply using <@ and garbage. Probably not a big deal

// This function sets the Keisatsu on a Hanin
func (c *Cagliostro) cmdPolice(s *discordgo.Session, m *discordgo.MessageCreate, person string) error {
	if !strings.HasPrefix(person, "<@") { // Case of !emo
		return errors.New("Please specify a criminal starting with @")
	}

	_, err := s.ChannelMessageSend(m.ChannelID, person+" COME WITH ME SIR <:CatPolice:432948329651896320> <:KannaPolice:407770382616100865> ") // Assumes those emojis exist in that server or this bot's on Nitro
	_ = c.cmdEmoji(s, m, "policegif")                                                                                                          // hard Coded in assuming policegif.gif exists

	return err
}

// This function literally murders someone
func (c *Cagliostro) cmdKill(s *discordgo.Session, m *discordgo.MessageCreate, person string) error {
	if !strings.HasPrefix(person, "<@") { // Case of !emo
		return errors.New("Please specify victim starting with @")
	}

	// Generate random number based on how many sayings we have, and pick which one to use based on that
	rando := rand.Intn(killSayings) // generate a random number from killSayings
	err := errors.New("Our Case statement failed somehow we got issues bud")
	switch rando {
	case 0:
		_, err = s.ChannelMessageSend(m.ChannelID, "Omae wa mou, shindeiru "+person+" <:mewgun:572168026506526760>")
	case 1:
		_, err = s.ChannelMessageSend(m.ChannelID, "Shinei "+person+" <:CatKnife:440039640590712832>")
	case 2:
		_, err = s.ChannelMessageSend(m.ChannelID, "It is time to pay for your crimes "+person)
		_ = c.cmdEmoji(s, m, "truckgif") // Assumes truckgif.gif exists
	}

	return err
}

// This function returns a random pasta (WIP)
func (c *Cagliostro) cmdPasta(s *discordgo.Session, m *discordgo.MessageCreate, person string) error {

	return nil
}
