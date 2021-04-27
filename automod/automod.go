package automod

import (
	"fmt"
	"strings"
	"time"

	"github.com/blackmesadev/black-mesa/automod/censor"
	"github.com/blackmesadev/black-mesa/automod/spam"
	"github.com/blackmesadev/black-mesa/config"
	"github.com/blackmesadev/black-mesa/logging"
	"github.com/blackmesadev/black-mesa/moderation"
	"github.com/blackmesadev/discordgo"
)

// Gets the closest level that the ideal level can match in the level -> interface map
func getClosestLevel(i []int64, targetLevel int64) int64 {
	var closest int64 = 0
	for _, level := range i {
		if level == targetLevel {
			return targetLevel
		}

		if level < targetLevel {
			closest = level
		} else {
			return closest // micro optimization; return early if the level is ever higher than the target
		}
	}

	return closest
}

func Process(s *discordgo.Session, m *discordgo.Message) {
	ok, reason, weight, _ := Check(s, m)
	if !ok {
		go RemoveMessage(s, m) // remove
		if strings.HasPrefix(reason, "Censor") { // log
			logging.LogMessageCensor(s, m, reason)
		} else {
			logging.LogMessageViolation(s, m, reason)
		}
		moderation.IssueStrike(s, m.GuildID, m.Author.ID, "AutoMod", weight, fmt.Sprintf("Violated AutoMod rules [%v]", reason), 0, m.ChannelID) // strike
		// and with that the moderation cycle is complete! :)
	}
}

// Return true if all is okay, return false if not.
// This function should be "silent" if a message is okay.
func Check(s *discordgo.Session, m *discordgo.Message) (bool, string, int, time.Time) {
	filterProcessingStart := time.Now()

	conf, err := config.GetConfig(m.GuildID)

	if conf == nil || err != nil {
		fmt.Println(conf, err)
		return true, "", 0, filterProcessingStart
	}

	automod := conf.Modules.Automod

	content := m.Content

	if len(automod.SpamLevels) == 0 && len(automod.SpamChannels) == 0 &&
		len(automod.CensorLevels) == 0 && len(automod.SpamChannels) == 0 {
		return true, "", 0, filterProcessingStart
	}

	censorChannel := automod.CensorChannels[m.ChannelID]

	// levels take priority
	userLevel := config.GetLevel(s, m.GuildID, m.Author.ID)

	i := 0
	automodCensorLevels := make([]int64, len(automod.CensorLevels))
	for k := range automod.CensorLevels {
		automodCensorLevels[i] = k
		i++
	}

	censorLevel := automod.CensorLevels[getClosestLevel(automodCensorLevels, userLevel)]

	// Level censors
	if censorLevel != nil {
		// Zalgo
		if censorLevel.FilterZalgo {
			ok := censor.ZalgoCheck(content)
			if !ok {
				return false, "Censor->Zalgo", 1, filterProcessingStart
			}
		}

		// Invites
		if censorLevel.FilterInvites {
			ok, invite := censor.InvitesWhitelistCheck(content, censorLevel.InvitesWhitelist)
			if !ok {
				return false, fmt.Sprintf("Censor->Invite (%v)", invite), 1, filterProcessingStart
			}
		} else if len(*censorLevel.InvitesBlacklist) != 0 {
			ok, invite := censor.InvitesBlacklistCheck(content, censorLevel.InvitesBlacklist)
			if !ok {
				return false, fmt.Sprintf("Censor->InvitesBlacklist (%v)", invite), 1, filterProcessingStart
			}
		}

		// Domains

		if censorLevel.FilterDomains {
			ok, domain := censor.DomainsWhitelistCheck(content, censorLevel.DomainWhitelist)
			if !ok {
				return false, fmt.Sprintf("Censor->Domain (%v)", domain), 1, filterProcessingStart
			}
		} else if len(*censorLevel.DomainBlacklist) != 0 {
			ok, domain := censor.DomainsBlacklistCheck(content, censorLevel.DomainBlacklist)
			if !ok {
				return false, fmt.Sprintf("Censor->DomainsBlacklist (%v)", domain), 1, filterProcessingStart
			}
		}

		// Strings / Substrings

		if censorLevel.FilterStrings {
			ok, str := censor.StringsCheck(content, censorLevel.BlockedStrings)
			if !ok {
				return false, fmt.Sprintf("Censor->BlockedString (%v)", str), 1, filterProcessingStart
			}
		}
	}

	// Channel censors
	if censorChannel != nil {
		// Zalgo
		if censorChannel.FilterZalgo {
			ok := censor.ZalgoCheck(content)
			if !ok {
				return false, "Censor->FilterZalgo", 1, filterProcessingStart
			}
		}

		// Invites
		if censorChannel.FilterInvites {
			ok, invite := censor.InvitesWhitelistCheck(content, censorChannel.InvitesWhitelist)
			if !ok {
				return false, fmt.Sprintf("Censor->Invite (%v)", invite), 1, filterProcessingStart
			}

		} else if len(*censorChannel.InvitesBlacklist) != 0 {
			ok, invite := censor.InvitesBlacklistCheck(content, censorChannel.InvitesBlacklist)
			if !ok {
				return false, fmt.Sprintf("Censor->InvitesBlacklist (%v)", invite), 1, filterProcessingStart
			}
		}

		// Domains

		if censorChannel.FilterDomains {
			ok, domain := censor.DomainsWhitelistCheck(content, censorChannel.DomainWhitelist)
			if !ok {
				return false, fmt.Sprintf("Censor->Domain (%v)", domain), 1, filterProcessingStart
			}
		} else if len(*censorChannel.DomainBlacklist) != 0 {
			ok, domain := censor.DomainsBlacklistCheck(content, censorChannel.DomainBlacklist)
			if !ok {
				return false, fmt.Sprintf("Censor->DomainsBlacklist (%v)", domain), 1, filterProcessingStart
			}
		}

		// Strings / Substrings

		if censorChannel.FilterStrings {
			ok, str := censor.StringsCheck(content, censorChannel.BlockedStrings)
			if !ok {
				return false, fmt.Sprintf("Censor->BlockedString (%v)", str), 1, filterProcessingStart
			}
		}
	}

	// Spam
	{ // messages
		ten, _ := time.ParseDuration("10s")
		limit := 5
		ok := spam.ProcessMaxMessages(m.Author.ID, m.GuildID, limit, ten, false)
		if !ok {
			return false, fmt.Sprintf("Spam->Messages (%v/%v)", limit, ten), 1, filterProcessingStart
		}
	}
	{ // newlines
		limit := 10
		ok, count := spam.ProcessMaxNewlines(m.Content, limit)
		if !ok {
			//                                                             1 strike per limit violation
			return false, fmt.Sprintf("Spam->NewLines (%v > %v)", count, limit), (count / limit), filterProcessingStart
		}
	}
	{ // mentions
		limit := 2
		ok, count := spam.ProcessMaxMentions(m, limit)
		if !ok {
			//                                                      1 strike for every mention over the limit
			return false, fmt.Sprintf("Spam->Mentions (%v > %v)", count, limit), (count - limit), filterProcessingStart
		}
		ok, count = spam.ProcessMaxRoleMentions(m, limit)
		if !ok {
            // see above
			return false, fmt.Sprintf("Spam->RoleMentions (%v > %v)", count, limit), (count - limit), filterProcessingStart
		}
	}
	{ // links
		limit := 2
		ok, count := spam.ProcessMaxLinks(m.Content, limit)
		if !ok {
			return false, fmt.Sprintf("Spam->Links (%v > %v)", count, limit), (count - limit), filterProcessingStart
		}
	}
	{ // uppercase
		limit := 50.0
		minLength := 20
		ok, percent := spam.ProcessMaxUppercase(m.Content, limit, minLength)
		if !ok {
			// flat rate because there's basically no calculation we can do here
			return false, fmt.Sprintf("Spam->Uppercase (%v%% > %v%%)", percent, limit), 1, filterProcessingStart
		}
	}
	{ // emoji
		limit := 10
		ok, count := spam.ProcessMaxEmojis(m, limit)
		if !ok {
			//                                                             1 strike per limit violation
			return false, fmt.Sprintf("Spam->Emojis (%v > %v)", count, limit), (count / limit), filterProcessingStart
		}
	}
	{ // attachments
		limit := 2
		ok, count := spam.ProcessMaxAttachments(m, limit)
		if !ok {
			//                                                       1 strike per attachment over the limit
			return false, fmt.Sprintf("Spam->Attachments (%v > %v)", count, limit), (count - limit), filterProcessingStart
		}
	}

	return true, "", 0, filterProcessingStart
}
