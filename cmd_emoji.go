package cagliostro

import (
	"errors"
	"fmt"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/arbovm/levenshtein"
	"github.com/bwmarrin/discordgo"
)

const (
	emojiSuggestions            = 8
	emojiSuggestionsMaxDistance = 3
)

// cmdEmoji handles the emoji command.
func (c *Cagliostro) cmdEmoji(s *discordgo.Session, m *discordgo.MessageCreate, emoji string) error {
	if emoji == "" { // Case of !emo
		return errors.New("Please specify an emoji")
	}

	var (
		err error
		f   *os.File
	)

	name := ""
	if strings.HasSuffix(emoji, "gif") { // Checks for gif as the final part
		name = fmt.Sprintf("%s.gif", emoji)
	} else { // If we don't get an emote specifying gif, we match to png instead
		name = fmt.Sprintf("%s.png", emoji) // Tries to match up by concatanating passed in emoji name to .png
	}

	file := path.Join(c.EmojiDir, name)

	f, err = os.Open(file)
	if err == nil {
		_, err = s.ChannelFileSend(m.ChannelID, name, f)

		f.Close()

		return err
	}

	var (
		entries      []string
		names        []string
		similarNames []string
	)

	f, err = os.Open(c.EmojiDir) // Open Emoji Directory
	if err != nil {
		return err
	}
	defer f.Close()

	entries, err = f.Readdirnames(-1) // Read names from our directory
	if err != nil {
		return err
	}

	names = make([]string, 0, len(entries))
	for _, entry := range entries { // Range over our entries
		if !strings.HasSuffix(entry, ".png") && !strings.HasSuffix(entry, ".gif") {
			continue
		}

		name := ""
		if strings.HasSuffix(entry, ".png") { // Check for which one we have, a .png or .gif, and Trim appropriate
			name = strings.TrimSuffix(entry, ".png")
		} else {
			name = strings.TrimSuffix(entry, ".gif")
		}
		name = fmt.Sprintf("%s", name)

		names = append(names, name)
	}

	if len(names) > 0 {
		sort.StringSlice(names).Sort()

		namesByDistance := [emojiSuggestionsMaxDistance + 1][]string{}
		for i := 0; i <= emojiSuggestionsMaxDistance; i++ {
			namesByDistance[i] = make([]string, 0)
		}

		for i := 0; i < len(names); i++ {
			d := levenshtein.Distance(emoji, names[i])
			if d <= emojiSuggestionsMaxDistance {
				nameCode := fmt.Sprintf("`%s`", names[i])
				namesByDistance[d] = append(namesByDistance[d], nameCode)
				similarNames = append(similarNames, nameCode)
			}
		}

		similarNames = make([]string, 0, emojiSuggestions)
		for i := 0; i <= emojiSuggestionsMaxDistance; i++ {
			if len(similarNames) > 0 {
				break
			}

			for j := 0; j < len(namesByDistance[i]) && len(similarNames) < emojiSuggestions; j++ {
				similarNames = append(similarNames, namesByDistance[i][j])
			}
		}
	}

	if len(similarNames) == 0 {
		_, err = s.ChannelMessageSend(m.ChannelID, "Unknown emoji.")

		return err
	}

	response := fmt.Sprintf("Did you mean: %s", strings.Join(similarNames, ", "))

	_, err = s.ChannelMessageSend(m.ChannelID, response)

	return err
}
