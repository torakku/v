package cagliostro

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/bwmarrin/discordgo"
)

// cmdEmoji handles the emoji command.
func (c *Cagliostro) cmdEmoji(s *discordgo.Session, m *discordgo.MessageCreate, emoji string) error {
	var (
		err error
		f   *os.File
	)

	name := fmt.Sprintf("%s.png", emoji)
	file := path.Join(c.EmojiDir, name)

	f, err = os.Open(file)
	if err != nil {
		return errors.New("Unknown emoji")
	}
	defer f.Close()

	_, err = s.ChannelFileSend(m.ChannelID, name, f)

	return err
}
