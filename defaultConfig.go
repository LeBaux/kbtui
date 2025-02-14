package main

var defaultConfig = `
[basics]
download_path = "/tmp/"
colorless = false
unicode_emojis = true

# The prefix before evaluating a command
cmd_prefix = "/"

[formatting]
# BASH-like PS1 variable equivalent
output_format = "┌──[$USER@$DEVICE] [$ID] [$DATE - $TIME]\n└╼ $MSG"
output_stream_format = "┌──[$USER@$DEVICE] [$ID] [$DATE - $TIME]\n└╼ $MSG"
output_mention_format = "┌──[$USER@$DEVICE] [$ID] [$DATE - $TIME]\n└╼ $MSG"
pm_format = "PM from $USER@$DEVICE: $MSG"

# 02 = Day, Jan = Month, 06 = Year
date_format = "02Jan06"

# 15 = hours, 04 = minutes, 05 = seconds
time_format = "15:04"


[colors]
	 [colors.channels]
		  [colors.channels.basic]
		  foreground = "normal"
		  [colors.channels.header]
		  foreground = "magenta"
		  bold = true
		  [colors.channels.unread]
		  foreground = "green"
		  italic = true

	 [colors.message]
		  [colors.message.body]
		  foreground = "normal"
		  [colors.message.header]
		  foreground = "grey"
		  [colors.message.mention]
		  foreground = "green"
		  italic = true
		  bold = true
		  [colors.message.id]
		  foreground = "yellow"
		  [colors.message.time]
		  foreground = "magenta"
		  [colors.message.sender_default]
		  foreground = "cyan"
		  bold = true
		  [colors.message.sender_device]
		  foreground = "cyan"
		  [colors.message.attachment]
		  foreground = "red"
		  [colors.message.link_url]
		  foreground = "yellow"
		  [colors.message.link_keybase]
		  foreground = "yellow"
		  [colors.message.reaction]
		  foreground = "magenta"
		  bold = true
		  [colors.message.code]
		  foreground = "cyan"
		  background = "grey"
	 [colors.feed]
		  [colors.feed.basic]
		  foreground = "grey"
		  [colors.feed.error]
		  foreground = "red"
`
