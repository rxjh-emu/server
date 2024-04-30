package protocol

import (
	"github.com/rxjh-emu/server/share/network"
)

type GameProtocol struct {
	MaxBufferSize int
	HeaderSize    int
	UserIdx       int
}

func NewGameProtocol() *GameProtocol {
	return &GameProtocol{
		MaxBufferSize: 4096,
		HeaderSize:    12,
	}
}

func (p *GameProtocol) SetUserIdx(idx int) {
	p.UserIdx = idx
}

func (p *GameProtocol) Read(buffer []byte) *network.PacketArgs {
	buffer = p.removePacketScope(buffer)

	pr := network.NewReader(buffer)
	packetLength := pr.ReadInt16()
	pr.ReadInt16() // cryptKey
	userIdx := pr.ReadInt16()
	opcode := pr.ReadInt16()
	dataLength := pr.ReadInt16()

	// Extract data from buffer
	data := pr.ReadBytes(int(dataLength))
	// Create new packet reader
	reader := network.NewReader(data)

	return &network.PacketArgs{
		Session:      nil,
		UserIdx:      int(userIdx),
		PacketLength: int(packetLength),
		Type:         int(opcode),
		Length:       int(dataLength),
		Data:         data,
		Reader:       reader,
	}
}

func (p *GameProtocol) Write(buffer []byte) *[]byte {
	// เพิ่มขนาด buffer 8 byte
	pw := network.NewEmptyWriter()
	pw.WriteInt16(int16(len(buffer) + 4))
	// pw.WriteInt16(0)
	pw.WriteInt32(0) // user idx
	pw.WriteBytes(buffer)

	_buffer := p.addPacketScope(pw.RawBytes())
	return &_buffer
}

func (p *GameProtocol) GetHeaderSize() int {
	return p.HeaderSize + 4 // + 4 from scope bytes
}

func (p *GameProtocol) removePacketScope(data []byte) []byte {
	if data[0] == 0xAA && data[1] == 0x55 && data[len(data)-2] == 0x55 && data[len(data)-1] == 0xAA {
		return data[2 : len(data)-2]
	}
	return data
}

func (p *GameProtocol) addPacketScope(data []byte) []byte {
	return append(append([]byte{0xAA, 0x55}, data...), 0x55, 0xAA)
}
