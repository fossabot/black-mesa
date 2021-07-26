package discord

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/blackmesadev/black-mesa/automod"
	"github.com/blackmesadev/black-mesa/config"
	"github.com/blackmesadev/discordgo"
)

func (bot *Bot) OnMessageCreate(s *discordgo.Session, mc *discordgo.MessageCreate) {

	var err error

	// Ignore all messages created by the Bot account itself
	if mc.Author.ID == s.State.User.ID {
		return
	}

	automod.Process(s, mc.Message)

	// Create Context struct that we can put various infos into
	ctx := &discordgo.Context{
		Content: strings.TrimSpace(mc.Content),
	}

	// Fetch the channel for this Message
	var c *discordgo.Channel
	c, err = s.State.Channel(mc.ChannelID)
	if err != nil {
		// Try fetchin  via REST API
		c, err = s.Channel(mc.ChannelID)
		if err != nil {
			log.Printf("unable to fetch Channel for Message, %s", err)
		} else {
			// Attempt to add this channel into our State
			err = s.State.ChannelAdd(c)
			if err != nil {
				log.Printf("error updating State with Channel, %s", err)
			}
		}
	}
	// Add Channel info into Context (if we successfully got the channel)
	if c != nil {
		if c.Type == discordgo.ChannelTypeDM {
			ctx.IsPrivate, ctx.IsDirected = true, true
		}
	}

	// Detect @name or @nick mentions
	if !ctx.IsDirected {

		// Detect if Bot was @mentioned
		for _, v := range mc.Mentions {

			if v.ID == s.State.User.ID {

				ctx.IsDirected, ctx.HasMention = true, true

				reg := regexp.MustCompile(fmt.Sprintf("<@!?(%s)>", s.State.User.ID))

				// Was the @mention the first part of the string?
				mentionSearch := reg.FindStringIndex(ctx.Content)
				if len(mentionSearch) > 0 {
					if mentionSearch[0] == 0 {
						ctx.HasMentionFirst = true
					}
				}

				// strip bot mention tags from content string
				ctx.Content = reg.ReplaceAllString(ctx.Content, "")

				break
			}
		}
	}

	prefix := config.GetPrefix(mc.GuildID)

	if !ctx.IsDirected && len(prefix) > 0 {
		if strings.HasPrefix(ctx.Content, prefix) {
			ctx.IsDirected, ctx.HasPrefix, ctx.HasMentionFirst = true, true, true
			ctx.Content = strings.TrimPrefix(ctx.Content, prefix)
		}
	}

	if !ctx.IsDirected {
		return
	}

	r, params, args := bot.Router.Match(ctx.Content)
	if r != nil {
		ctx.Fields = params
		r.Run(s, mc.Message, ctx, args)
		return
	}

	if bot.Router.Default != nil && (ctx.HasMentionFirst) {
		bot.Router.Default.Run(s, mc.Message, ctx, make([]string, 0))
	}
}
