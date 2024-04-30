package def

import (
	"github.com/jmoiron/sqlx"
	"github.com/rxjh-emu/server/cmd/masterserver/data"
	"github.com/rxjh-emu/server/cmd/masterserver/database"
	"github.com/rxjh-emu/server/cmd/masterserver/server"
	"github.com/rxjh-emu/server/share/rpc"
)

var ServerConfig = &Config{}
var ServerSettings = &Settings{}
var RPCHandler = &rpc.Server{}
var LoginDatabase = &sqlx.DB{}
var ServerManager = &server.ServerManager{}
var DatabaseManager = &database.DatabaseManager{}
var DataLoader = &data.Loader{}
