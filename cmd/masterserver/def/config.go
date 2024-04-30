package def

import (
	"fmt"
	"strconv"

	"github.com/rxjh-emu/server/cmd/masterserver/database"
	"github.com/rxjh-emu/server/share/conf"
	"github.com/rxjh-emu/server/share/directory"
	"github.com/rxjh-emu/server/share/log"
)

type Config struct {
	Port       int
	ServerName string
	LoginDB    *database.Database
	GameDB     map[int]*database.Database
}

// Attempts to read server configuration file
func (c *Config) Read() {
	log.Info("Reading configuration...")

	var location = directory.Root() + "/cfg/masterserver.ini"

	// parse configuration file...
	if err := conf.Open(location); err != nil {
		log.Fatal(err.Error())
		return
	}

	// read values from configuration...
	c.Port = conf.GetInt("network", "port", 9001)
	c.ServerName = conf.GetString("server", "name", "World")

	// login db
	c.LoginDB = &database.Database{}
	c.LoginDB.Ip = conf.GetString("login", "ip", "127.0.0.1")
	c.LoginDB.Port = conf.GetInt("login", "port", 3306)
	c.LoginDB.Name = conf.GetString("login", "database", "database")
	c.LoginDB.User = conf.GetString("login", "username", "root")
	c.LoginDB.Password = conf.GetString("login", "password", "")

	// load all game databases
	c.LoadGameDB()
}

// Attemps to read all [1..255] GameDatabase configurations
func (c *Config) LoadGameDB() {
	c.GameDB = make(map[int]*database.Database)

	for i := 1; i < 256; i++ {
		var section = fmt.Sprintf("game_%d", i)
		if conf.SectionExist(section) {
			c.GameDB[i] = &database.Database{
				Ip:       conf.GetString(section, "ip", "127.0.0.1"),
				Port:     conf.GetInt(section, "port", 3306),
				Name:     conf.GetString(section, "database", "database"),
				User:     conf.GetString(section, "username", "root"),
				Password: conf.GetString(section, "password", ""),
				Index:    i,
				DB:       nil,
				Config:   "",
			}

			c.GameDB[i].Config = c.GetDBConfig(c.GameDB[i])
		}
	}
}

// Returns GameDatabase configuration string
func (c *Config) GetDBConfig(db *database.Database) string {
	var str = db.User + ":" + db.Password
	str += "@tcp(" + db.Ip + ":" + strconv.Itoa(db.Port) + ")"
	str += "/" + db.Name + "?parseTime=true"
	return str
}
