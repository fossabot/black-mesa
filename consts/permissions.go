package consts

const (
	CATEGORY_MODERATION = "moderation"
	CATEGORY_ADMIN      = "admin"
	CATEGORY_GUILD      = "guild"
	CATEGORY_ROLES      = "roles"
	CATEGORY_TRUSTED    = "trusted"
	CATEGORY_MUSIC      = "music"

	PERMISSION_BAN          = CATEGORY_MODERATION + ".ban"
	PERMISSION_KICK         = CATEGORY_MODERATION + ".kick"
	PERMISSION_MUTE         = CATEGORY_MODERATION + ".mute"
	PERMISSION_PURGE        = CATEGORY_MODERATION + ".purge"
	PERMISSION_PURGE_ALL    = CATEGORY_MODERATION + ".purgeall"
	PERMISSION_REMOVEACTION = CATEGORY_MODERATION + ".remove"
	PERMISSION_SEARCH       = CATEGORY_MODERATION + ".search"
	PERMISSION_SEARCHSELF   = CATEGORY_MODERATION + ".searchself"
	PERMISSION_SOFTBAN      = CATEGORY_MODERATION + ".softban"
	PERMISSION_STRIKE       = CATEGORY_MODERATION + ".strike"
	PERMISSION_UNBAN        = CATEGORY_MODERATION + ".unban"
	PERMISSION_UNMUTE       = CATEGORY_MODERATION + ".unmute"

	PERMISSION_CONFIGGET = CATEGORY_ADMIN + ".get"
	PERMISSION_CONFIGSET = CATEGORY_ADMIN + ".set"
	PERMISSION_MAKEMUTE  = CATEGORY_ADMIN + ".makemute"
	PERMISSION_SETUP     = CATEGORY_ADMIN + ".setup"

	PERMISSION_VIEWCMDLEVEL  = CATEGORY_GUILD + ".viewcommandlevel"
	PERMISSION_VIEWUSERLEVEL = CATEGORY_GUILD + ".viewuserlevel"

	PERMISSION_ROLEADD    = CATEGORY_ROLES + ".add"
	PERMISSION_ROLEREMOVE = CATEGORY_ROLES + ".remove"
	PERMISSION_ROLECREATE = CATEGORY_ROLES + ".create"

	PERMISSION_BANFILE = CATEGORY_TRUSTED + ".banfile"

	PERMISSION_PLAY   = CATEGORY_MUSIC + ".play"
	PERMISSION_STOP   = CATEGORY_MUSIC + ".stop"
	PERMISSION_SKIP   = CATEGORY_MUSIC + ".skip"
	PERMISSION_REMOVE = CATEGORY_MUSIC + ".remove"
	PERMISSION_DC     = CATEGORY_MUSIC + ".dc"
	PERMISSION_SEEK   = CATEGORY_MUSIC + ".seek"
	PERMISSION_VOLUME = CATEGORY_MUSIC + ".volume"
	PERMISSION_QUERY  = CATEGORY_MUSIC + ".query"
)
