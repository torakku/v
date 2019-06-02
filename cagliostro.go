package cagliostro

import (
	"errors"
	"strings"
	"sync"

	"github.com/bwmarrin/discordgo"
)

// Cagliostro is the main struct that wraps all the data required by
// the bot.
type Cagliostro struct {
	Token    string
	Prefix   string
	EmojiDir string

	Logger Logger

	eventsMutex    sync.Mutex
	currentEvents  []*CachedEvent
	upcomingEvents []*CachedEvent

	session *discordgo.Session
}

// logger returns c.Logger when available, falling back to a no-op logger.
//
// This function never returns nil.
func (c *Cagliostro) logger() Logger {
	if c.Logger == nil {
		return quietLoggerSingleton
	}

	return c.Logger
}

// OnMessageCreate is the event handler for discordgo library.
func (c *Cagliostro) OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if c.Prefix == "" {
		return
	}
	if !strings.HasPrefix(m.Content, c.Prefix) {
		return
	}

	var (
		err error

		head string
		tail string

		log = c.logger()
	)

	parts := strings.SplitN(m.Content, " ", 2)
	if len(parts) == 0 {
		return
	}

	head = strings.TrimPrefix(parts[0], c.Prefix)
	if len(parts) > 1 {
		tail = parts[1]
	}

	log.Printf("head=%#v tail=%#v", head, tail)

	switch head {
	case "events":
		err = c.cmdEvents(s, m)
	case "emo":
		err = c.cmdEmoji(s, m, strings.Split(tail, " ")[0]) // This way we can emoji and also include more text after. Not included within other commands because usernames can have spaces
	case "police":
		err = c.cmdPolice(s, m, tail)
	case "kill":
		err = c.cmdKill(s, m, tail)
	case "pasta":
		// err = c.cmdPasta(s, m)
	}

	if err != nil {
		_, _ = s.ChannelMessageSend(m.ChannelID, err.Error())
	}
}

// Open starts a connection with Discord.
func (c *Cagliostro) Open() error {
	if c.session != nil {
		return errors.New("session already opened")
	}

	var (
		err error
		s   *discordgo.Session

		log = c.logger()
	)

	s, err = discordgo.New("Bot " + c.Token)
	if err != nil {
		return err
	}

	s.AddHandler(c.OnMessageCreate)

	log.Print("opening session")

	err = s.Open()
	if err != nil {
		return err
	}

	log.Print("session opened successfully")

	c.session = s

	return nil
}

// Close terminates the connection with Discord.
func (c *Cagliostro) Close() error {
	if c.session == nil {
		return errors.New("no session")
	}

	log := c.logger()
	log.Print("closing session")

	err := c.session.Close()

	return err
}
