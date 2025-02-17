package structs

type WebAccess struct {
	Admin  []string
	Editor []string
	Viewer []string
}

type Persistance struct {
	Roles            bool
	WhitelistedRoles []string `json:"whitelistedRoles" bson:"whitelistedRoles"` // slice of ids
	Nickname         bool
	Voice            bool
}

type ReactRoleEmote struct {
	Role string
}

type ReactRoleChannel struct {
	Emotes map[string]*ReactRoleEmote // emojiID : reactRoleEmote
}

type ReactRoles struct {
	Channel map[string]*ReactRoleChannel // channelID : reactRoleChannel
}

type Guild struct {
	ConfirmActions      bool              `json:"confirmActions" bson:"confirmActions"`
	RoleAliases         map[string]string `json:"roleAliases" bson:"roleAliases"`                 // name: roleid
	SelfAssignableRoles map[string]string `json:"selfAssignableRoles" bson:"selfAssignableRoles"` // name: roleid
	LockedRoles         []string          `json:"lockedRoles" bson:"lockedRoles"`                 // slice of ids
	Persistance         *Persistance      `json:"persistance" bson:"persistance"`
	AutoRole            []string          `json:"autoRole" bson:"autoRole"` // slice of ids
	ReactRoles          *ReactRoles       `json:"reactRoles" bson:"reactRoles"`
	UnsafePermissions   bool              `json:"unsafePermissions" bson:"unsafePermissions"`
	StaffLevel          int64             `json:"staffLevel" bson:"staffLevel"`
}

type Censor struct {
	FilterZalgo            bool     `json:"filterZalgo" bson:"filterZalgo"`
	FilterInvites          bool     `json:"filterInvites" bson:"filterInvites"`
	FilterDomains          bool     `json:"filterDomains" bson:"filterDomains"`
	FilterStrings          bool     `json:"filterStrings" bson:"filterStrings"`
	FilterIPs              bool     `json:"filterIPs" bson:"filterIPs"`
	FilterRegex            bool     `json:"filterRegex" bson:"filterRegex"`
	FilterEnglish          bool     `json:"filterEnglish" bson:"filterEnglish"`
	FilterObnoxiousUnicode bool     `json:"filterObnoxiousUnicode" bson:"filterObnoxiousUnicode"`
	FilterUntrustworthy    bool     `json:"filterUntrustworthy" bson:"filterUntrustworthy"`
	InvitesWhitelist       []string `json:"invitesWhitelist" bson:"invitesWhitelist"`   // slice of invitelinks/ids
	InvitesBlacklist       []string `json:"invitesBlacklist" bson:"invitesBlacklist"`   // slice of invitelinks/ids
	DomainWhitelist        []string `json:"domainWhitelist" bson:"domainWhitelist"`     // slice of domains
	DomainBlacklist        []string `json:"domainBlacklist" bson:"domainBlacklist"`     // slice of domains
	BlockedSubstrings      []string `json:"blockedSubstrings" bson:"blockedSubstrings"` // slice of substrings
	BlockedStrings         []string `json:"blockedStrings" bson:"blockedStrings"`       // slice of strings
	Regex                  string   `json:"regex" bson:"regex"`
}

type Spam struct {
	Punishment          string
	PunishmentDuration  int64   `json:"punishmentDuration" bson:"punishmentDuration"` // seconds
	Count               int64   `json:"count" bson:"count"`                           // amount per interval
	Interval            int64   `json:"interval" bson:"interval"`                     // seconds
	MaxMessages         int64   `json:"maxMessages" bson:"maxMessages"`
	MaxMentions         int64   `json:"maxMentions" bson:"maxMentions"`
	MaxLinks            int64   `json:"maxLinks" bson:"maxLinks"`
	MaxAttachments      int64   `json:"maxAttachments" bson:"maxAttachments"`
	MaxEmojis           int64   `json:"maxEmojis" bson:"maxEmojis"`
	MaxNewlines         int64   `json:"maxNewlines" bson:"maxNewlines"`
	MaxDuplicates       int64   `json:"maxDuplicates" bson:"maxDuplicates"`
	MaxCharacters       int64   `json:"maxCharacters" bson:"maxCharacters"`
	MaxUppercasePercent float64 `json:"maxUppercasePercent" bson:"maxUppercasePercent"`
	MinUppercaseLimit   int64   `json:"minUppercaseLimit" bson:"minUppercaseLimit"`
	Clean               bool    `json:"clean" bson:"clean"`
	CleanCount          int64   `json:"cleanCount" bson:"cleanCount"`
	CleanDuration       int64   `json:"cleanDuration" bson:"cleanDuration"`
}

type GuildOptions struct {
	MinimumAccountAge string `json:"minimumAccountAge" bson:"minimumAccountAge"`
}

type Automod struct {
	Enabled          bool               `json:"enabled" bson:"enabled"`
	GuildOptions     *GuildOptions      `json:"guildOptions" bson:"guildOptions"`
	CensorLevels     map[int64]*Censor  `json:"censorLevels" bson:"censorLevels"`
	CensorChannels   map[string]*Censor `json:"censorChannels" bson:"censorChannels"`
	SpamLevels       map[int64]*Spam    `json:"spamLevels" bson:"spamLevels"`
	SpamChannels     map[string]*Spam   `json:"spamChannels" bson:"spamChannels"`
	PublicHumilation bool               `json:"publicHumilation" bson:"publicHumilation"`
	StaffBypass      bool               `json:"staffBypass" bson:"staffBypass"`
}

type Logging struct {
	Enabled            bool     `json:"enabled" bson:"enabled"`
	ChannelID          string   `json:"channelID" bson:"channelID"`
	IncludeActions     []string `json:"includeActions" bson:"includeActions"` // list of actions
	ExcludeActions     []string `json:"excludeActions" bson:"excludeActions"` // list of actions
	Timestamps         bool     `json:"timestamps" bson:"timestamps"`
	Timezone           string   `json:"timezone" bson:"timezone"`
	IgnoredUsers       []string `json:"ignoredUsers" bson:"ignoredUsers"`             // slice of user ids
	IgnoredChannels    []string `json:"ignoredChannels" bson:"ignoredChannels"`       // slice of channel ids
	NewMemberThreshold int64    `json:"newMemberThreshold" bson:"newMemberThreshold"` // seconds
}

type Moderation struct {
	CensorSearches              bool                       `json:"censorSearches" bson:"censorSearches"`
	CensorStaffSearches         bool                       `json:"censorStaffSearches" bson:"censorStaffSearches"`
	ConfirmActionsMessage       bool                       `json:"confirmActionsMessage" bson:"confirmActionsMessage"`
	ConfirmActionsMessageExpiry int64                      `json:"confirmActionsMessageExpiry" bson:"confirmActionsMessageExpiry"`
	ConfirmActionsReaction      bool                       `json:"confirmActionsReaction" bson:"confirmActionsReaction"`
	DefaultStrikeDuration       string                     `json:"defaultStrikeDuration" bson:"defaultStrikeDuration"`
	DisplayNoPermission         bool                       `json:"displayNoPermission" bson:"displayNoPermission"`
	MuteRole                    string                     `json:"muteRole" bson:"muteRole"`
	ReasonEditLevel             int64                      `json:"reasonEditLevel" bson:"reasonEditLevel"`
	NotifyActions               bool                       `json:"notifyActions" bson:"notifyActions"`
	ShowModeratorOnNotify       bool                       `json:"showModeratorOnNotify" bson:"showModeratorOnNotify"`
	SilenceLevel                int64                      `json:"silenceLevel" bson:"silenceLevel"`
	StrikeEscalation            map[int64]StrikeEscalation `json:"strikeEscalation" bson:"strikeEscalation"`
	StrikeCushioning            int64                      `json:"strikeCushioning" bson:"strikeCushioning"`
}

type StrikeEscalation struct {
	Type     string `json:"type" bson:"type"`
	Duration string `json:"duration" bson:"duration"`
}

type AntiNuke struct {
	Enabled      bool                        `json:"enabled" bson:"enabled"`
	MemberRemove map[int64]AntiNukeThreshold `json:"memberRemove" bson:"memberRemove"` // index is level, value is what to do.
}

type AntiNukeThreshold struct {
	Max      int64  `json:"max" bson:"max"`
	Interval int64  `json:"interval" bson:"interval"`
	Type     string `json:"type" bson:"type"`
}

type VoteMute struct {
	Enabled         bool  `json:"enabled" bson:"enabled"`
	MaxDuration     int64 `json:"maxDuration" bson:"maxDuration"`
	UpvotesRequired int64 `json:"upvotesRequired" bson:"upvotesRequired"`
	ExpiresAfter    int64 `json:"expiresAfter" bson:"expiresAfter"`
}

type Voting struct {
	UpvoteEmoji   string    `json:"upvoteEmoji" bson:"upvoteEmoji"`
	UpvoteEmojiID string    `json:"upvoteEmojiId" bson:"upvoteEmojiId"`
	VoteMute      *VoteMute `json:"voteMute" bson:"voteMute"`
}

type Modules struct {
	Guild      *Guild      `json:"guild" bson:"guild"`
	Automod    *Automod    `json:"automod" bson:"automod"`
	Logging    *Logging    `json:"logging" bson:"logging"`
	Moderation *Moderation `json:"moderation" bson:"moderation"`
	AntiNuke   *AntiNuke   `json:"antiNuke" bson:"antiNuke"`
	Voting     *Voting     `json:"voting" bson:"voting"`
}

type Config struct {
	guildID     string
	Nickname    string           `json:"nickname" bson:"nickname"`
	WebAccess   *WebAccess       `json:"webAccess" bson:"webAccess"`
	Prefix      string           `json:"prefix" bson:"prefix"`
	Permissions map[string]int64 `json:"permissions" bson:"permissions"`
	Levels      map[string]int64 `json:"levels" bson:"levels"`
	Modules     *Modules         `json:"modules" bson:"modules"`
}
