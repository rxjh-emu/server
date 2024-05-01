package packet

import (
	"strings"

	"github.com/rxjh-emu/server/share/models/character"
	"github.com/rxjh-emu/server/share/network"
	"github.com/rxjh-emu/server/share/rpc"
)

func EnterLobby(session *network.Session, reader *network.Reader) {
	packet := network.NewWriter(SM_ENTER_LOBBY)
	packet.WriteInt32(631327)
	packet.WriteInt16(-1)
	packet.WriteInt16(0)

	session.Send(packet)
}

func CharacterList(session *network.Session, reader *network.Reader) {
	// TODO: load characters from masterserver

	packet := network.NewWriter(SM_CHARACTER_LIST)
	packet.WriteByte(-1)
	session.Send(packet)
}

func CheckName(session *network.Session, reader *network.Reader) {
	name := strings.Replace(reader.ReadString(15), "\x00", "", -1)

	// check character name from masterserver rpc
	var r = character.CheckNameRes{}
	g_RPCHandler.Call(rpc.CheckCharacterName, character.CheckNameReq{Server: byte(g_ServerSettings.ServerId), Name: name}, &r)
	var result = r.Result

	packet := network.NewWriter(SM_CHECK_NAME)
	packet.WriteInt32(result)
	packet.AppendString(name, 15)

	session.Send(packet)
}

func CreateCharacter(session *network.Session, reader *network.Reader) {
	name := strings.Replace(reader.ReadString(16), "\x00", "", -1)
	class := reader.ReadByte()
	hair := reader.ReadByte()        // Hair
	haircolor := reader.ReadUint16() // Hair Color
	face := reader.ReadByte()        // face
	voice := reader.ReadByte()       // Voice
	gender := reader.ReadByte()

	// create character to masterserver rpc
	var r = character.CreateCharacterRes{}
	g_RPCHandler.Call(rpc.CreateCharacter, character.CreateCharacterReq{
		Server:    byte(g_ServerSettings.ServerId),
		AccountId: int(session.Data.AccountId),
		Name:      name,
		Class:     int(class),
		Hair:      int(hair),
		HairColor: int(haircolor),
		Face:      int(face),
		Voice:     int(voice),
		Gender:    int(gender),
	}, &r)

	packet := network.NewWriter(SM_CHARACTER_CREATE)
	packet.WriteInt32(r.Result)

	session.Send(packet)
}

func DeleteCharacter(session *network.Session, reader *network.Reader) {
}

func SelectCharacter(session *network.Session, reader *network.Reader) {
}

func SignOut(session *network.Session, reader *network.Reader) {
}
