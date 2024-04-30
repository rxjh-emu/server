package packet

import (
	"strings"

	"github.com/rxjh-emu/server/share/log"
	"github.com/rxjh-emu/server/share/network"
)

func CharacterList(session *network.Session, reader *network.Reader) {

	// todo
	log.Debug("CharacterList")

	packet := network.NewWriter(SM_CHARACTER_LIST)
	packet.WriteByte(-1)
	session.Send(packet)
}

func CheckName(session *network.Session, reader *network.Reader) {
	name := strings.Replace(reader.ReadString(15), "\x00", "", -1)
	log.Debugf("CheckName name: %s", name)

	// TODO: check character name from masterserver rpc

	packet := network.NewWriter(SM_CHECK_NAME)
	packet.WriteInt32(1)
	packet.AppendString(name, 15)

	session.Send(packet)
}
