package def

import (
	"github.com/rxjh-emu/server/share/conf"
	"github.com/rxjh-emu/server/share/directory"
	"github.com/rxjh-emu/server/share/log"
)

type Config struct {
	PublicIp   string
	Port       int
	MaxUsers   int
	UseLocalIp bool

	ServerType int

	MasterIp   string
	MasterPort int

	ScriptDirectory string
}

// Attempts to read server configuration file
func (c *Config) Read() {
	log.Info("Reading configuration...")

	var location = directory.Root() + "/cfg/" + GetName() + ".ini"

	// parse configuration file...
	if err := conf.Open(location); err != nil {
		log.Fatal(err.Error())
		return
	}

	// read values from configuration...
	c.PublicIp = conf.GetString("network", "ip", "127.0.0.1")
	c.Port = conf.GetInt("network", "port", 16101)
	c.MaxUsers = conf.GetInt("network", "max_users", 100)
	c.UseLocalIp = conf.GetBool("network", "use_local_ip", false)

	c.ServerType = conf.GetInt("server", "server_type", 0)

	c.MasterIp = conf.GetString("master", "ip", "127.0.0.1")
	c.MasterPort = conf.GetInt("master", "port", 9001)

	c.ScriptDirectory = conf.GetString("script", "directory", "")
}
