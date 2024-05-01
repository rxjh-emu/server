package packet

import (
	"encoding/hex"
	"strings"

	"github.com/rxjh-emu/server/share/models/character"
	"github.com/rxjh-emu/server/share/network"
	"github.com/rxjh-emu/server/share/rpc"
	"github.com/rxjh-emu/server/share/util"
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
	var r = character.ListRes{}
	g_RPCHandler.Call(rpc.LoadCharacters, character.ListReq{Server: byte(g_ServerSettings.ServerId), Account: int(session.Data.AccountId)}, &r)

	session.Data.CharacterList = r.List

	packet := network.NewWriter(SM_CHARACTER_LIST)

	if len(r.List) > 0 {
		for i := 0; i < len(r.List); i++ {
			c := r.List[i]
			// 3889 bytes per character
			packet.WriteByte(c.Slot)
			packet.AppendString(c.Name, 15)
			packet.WriteByte(0)
			packet.WriteInt32(0)
			packet.AppendString("", 15) // may be guild name
			packet.WriteByte(0)
			packet.WriteInt32(0)
			packet.WriteInt16(c.Level)
			packet.WriteByte(0) // ? faction
			packet.WriteByte(c.Fame)
			packet.WriteByte(c.Class)
			packet.WriteByte(c.Hair)
			packet.WriteBytes(util.ConvertUint16ToBytes(uint16(c.HairColor)))
			packet.WriteByte(c.Face)
			packet.WriteByte(c.Voice)
			packet.WriteBytes(make([]byte, 8))
			packet.WriteByte(2) // ?
			packet.WriteByte(c.Gender)
			packet.WriteFloat(c.Position.X)
			packet.WriteFloat(c.Position.Z)
			packet.WriteFloat(c.Position.Y)
			packet.WriteInt32(c.Position.Map)
			packet.WriteInt32(17000433)

			temp, _ := hex.DecodeString("0000000000000000FFFFFFFFFFFFFFFF0100000000000000FFFFFFFFFFFFFFFF0200000000000000FFFFFFFFFFFFFFFF")
			packet.WriteBytes(temp)
			packet.WriteInt32(0)
			packet.WriteInt16(c.MaxHp)
			packet.WriteInt16(c.MaxMp)
			packet.WriteInt32(c.MaxRp)
			packet.WriteInt64(c.Exp) // may be max exp
			packet.WriteInt16(c.CurrentHp)
			packet.WriteInt16(c.CurrentMp)
			packet.WriteInt32(c.CurrentRp)
			packet.WriteInt64(c.Exp) // may be current exp
			packet.WriteInt32(0)
			packet.WriteBytes(make([]byte, 16))

			// TODO: Equipment Items packet

			packet.WriteBytes(make([]byte, 3704))
			session.Send(packet)
		}
	} else {
		packet.WriteByte(-1)

		session.Send(packet)
	}

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
	hair := reader.ReadByte()
	haircolor := reader.ReadUint16()
	face := reader.ReadByte()
	voice := reader.ReadByte()
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
