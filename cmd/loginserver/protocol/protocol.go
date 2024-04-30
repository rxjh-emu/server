package protocol

import (
	"github.com/rxjh-emu/server/share/log"
	"github.com/rxjh-emu/server/share/network"
	"github.com/rxjh-emu/server/share/util"
)

type LoginProtocol struct {
	MaxBufferSize int
	HeaderSize    int
	UserIdx       int
}

func NewLoginProtocol() *LoginProtocol {
	return &LoginProtocol{
		MaxBufferSize: 4096,
		HeaderSize:    4,
	}
}

func (p *LoginProtocol) SetUserIdx(idx int) {
	p.UserIdx = idx
}

func (p *LoginProtocol) Read(buffer []byte) *network.PacketArgs {
	// log.Debugf("Read: %v", util.FormatHex(buffer))

	pr := network.NewReader(buffer)
	// Extract opcode (2 bytes, little-endian)
	opcode := pr.ReadInt16()

	// Extract data length (2 bytes, little-endian)
	dataLength := pr.ReadInt16()

	// Extract data from buffer
	data := pr.ReadBytes(int(dataLength))

	// Create new packet reader
	reader := network.NewReader(data)

	// Create packet event argument
	return &network.PacketArgs{
		Session: nil,
		Type:    int(opcode),
		Length:  int(dataLength),
		Data:    data,
		Reader:  reader,
	}
}

func (p *LoginProtocol) Write(buffer []byte) *[]byte {
	log.Debugf("Write: %v", util.FormatHex(buffer))
	return &buffer
}

func (p *LoginProtocol) GetHeaderSize() int {
	return p.HeaderSize
}
