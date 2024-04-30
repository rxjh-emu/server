package def

import "github.com/rxjh-emu/server/share/models/server"

type Settings struct {
	server.Settings
	ServerId  int
	ChannelId int
}
