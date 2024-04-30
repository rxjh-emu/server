package packet

import (
	"strings"

	"github.com/rxjh-emu/server/share/log"
	"github.com/rxjh-emu/server/share/models/account"
	"github.com/rxjh-emu/server/share/network"
	"github.com/rxjh-emu/server/share/rpc"
)

func Authorize(session *network.Session, reader *network.Reader) {
	log.Debug("Authorize")

	username := strings.Replace(reader.ReadString(31), "\x00", "", -1)
	password := strings.TrimSpace(reader.ReadString(31))
	reader.ReadInt16()
	reader.ReadInt16()
	channelIdx := reader.ReadInt16()
	ipAddress := strings.Replace(reader.ReadString(15), "\x00", "", -1)

	log.Debugf("Username: %s, Pass: %s, Channel: %d, IP: %s", username, password, channelIdx, ipAddress)

	var r = account.GetAccountRes{}
	err := g_RPCHandler.Call(rpc.GetAccount, account.GetAccountReq{UserId: username}, &r)

	log.Debugf("Result: %v", r)
	if err != nil {
		r.Status = account.ErrorMessage
	}

	// password verified
	packet := network.NewWriter(SM_AUTHORIZE)
	packet.WriteInt32(0)
	packet.WriteInt32(1)
	packet.WriteInt32(0)
	packet.WriteInt32(1)
	packet.WriteInt32(0)
	packet.WriteInt32(1)
	packet.WriteInt32(0)
	packet.WriteUint32(2288888516)
	packet.WriteByte(0)
	packet.WriteByte(0)
	packet.WriteByte(0)
	packet.WriteByte(0)
	packet.WriteByte(0)
	packet.WriteByte(0)
	packet.WriteByte(0)
	packet.WriteByte(0)
	packet.WriteInt32(36)

	session.Send(packet)

	if r.Status == account.Success {
		log.Infof("Account `%s` successfully logged in.", username)
		session.Data.AccountId = r.Id
		session.Data.LoggedIn = true
	} else if r.Status == account.Onlined {
		log.Infof("Account `%s` already logged in.", username)
		// todo disconnect old session
		session.Data.AccountId = r.Id
	} else {
		log.Infof("Account `%s` failed to log in.", username)
		session.Close()
	}
}
