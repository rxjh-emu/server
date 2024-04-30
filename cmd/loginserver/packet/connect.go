package packet

import (
	"strings"

	"github.com/rxjh-emu/server/share/log"
	"github.com/rxjh-emu/server/share/models/account"
	"github.com/rxjh-emu/server/share/network"
	"github.com/rxjh-emu/server/share/rpc"
)

func Login(session *network.Session, reader *network.Reader) {
	// skip 1 byte
	reader.ReadByte()
	username := strings.TrimSpace(reader.ReadString(int(reader.ReadInt16())))
	password := strings.TrimSpace(reader.ReadString(int(reader.ReadInt16())))

	var r = account.AuthResponse{Status: account.None}
	err := g_RPCHandler.Call(rpc.AuthCheck, account.AuthRequest{UserId: username, Password: password}, &r)

	message := "N/A"
	// if server is down...
	if err != nil {
		r.Status = account.ErrorMessage
		message = "Server is not available, please try again later."
	}

	var packet = network.NewWriter(SM_LOGIN)
	switch r.Status {
	case account.None:
		packet.WriteInt16(account.ErrorMessage)
		packet.WriteString("Error while login.")
		packet.WriteByte(0)
	case account.Success:
		session.Data.Username = username

		packet.WriteInt32(0)
		// packet.WriteByte(32)
		// packet.WriteByte(32)
		packet.WriteString(username)
		packet.WriteString("e265ed793f8ad01f5338e1ec7bc0623534bc4c406f518de82877fa0eb3058b2a")
		packet.WriteString("601471cfb6b7747948be85e83885295c23a2751000c6558501f88932d0824e58")
		packet.WriteString(username)
	case account.AccountNotFound:
		packet.WriteInt16(9)
		packet.WriteInt16(3)
	case account.WrongPassword:
		packet.WriteInt16(10)
		packet.WriteInt16(3)
	case account.Onlined:
	case account.ErrorMessage:
		packet.WriteInt16(23)
		packet.WriteString(message)
		packet.WriteByte(0)
	}

	session.Send(packet)

	if r.Status == account.Success {
		log.Infof("Account `%s` successfully logged in.", username)
		session.Data.AccountId = r.Id
		session.Data.LoggedIn = true
		// event.Trigger(event.PlayerLogin, session, username, true)
	} else if r.Status == account.Onlined {
		session.Data.AccountId = r.Id
		log.Infof("User `%s` double login attempt.", username)
	} else {
		log.Infof("User `%s` failed to log in.", username)
		// event.Trigger(event.PlayerLogin, session, username, false)
	}
}
