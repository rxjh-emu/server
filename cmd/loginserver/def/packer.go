package def

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/DarthPestilane/easytcp"
	"github.com/rxjh-emu/server/share/log"
	"github.com/rxjh-emu/server/share/util"
	"github.com/spf13/cast"
)

// LoginPacker treats packet as:
//
// opcode(2)|dataSize(2)|data(n)
//
// | segment     | type   | size    | remark                |
// | ----------- | ------ | ------- | --------------------- |
// | `opcode`	 | uint16 | 2       | opcode                |
// | `dataSize`  | uint16 | 2       | length of data        |
// | `data`      | []byte | dynamic |                       |
type LoginPacker struct{}

func (p *LoginPacker) bytesOrder() binary.ByteOrder {
	return binary.LittleEndian
}

func (p *LoginPacker) Pack(msg *easytcp.Message) ([]byte, error) {
	// format: id(2)|dataSize(2)|data(n)

	opcode := cast.ToUint16(msg.ID())
	buffer := make([]byte, 4+len(msg.Data()))
	p.bytesOrder().PutUint16(buffer[:2], uint16(opcode))           // write opcode
	p.bytesOrder().PutUint16(buffer[2:4], uint16(len(msg.Data()))) // write data size
	copy(buffer[4:4+len(msg.Data())], msg.Data())                  // write data

	log.Debugf("Pack: %v", util.FormatHex(buffer))

	return buffer, nil
}

func (p *LoginPacker) Unpack(reader io.Reader) (*easytcp.Message, error) {
	// format: id(2)|dataSize(2)|data(n)

	headerBuff := make([]byte, 2+2)
	if _, err := io.ReadFull(reader, headerBuff); err != nil {
		if err == io.EOF {
			return nil, err
		}
		return nil, fmt.Errorf("read header err: %s", err)
	}
	opcode := p.bytesOrder().Uint16(headerBuff[:2])        // read opcode
	dataSize := int(p.bytesOrder().Uint16(headerBuff[2:])) // read data size

	bodyBuff := make([]byte, dataSize)
	if _, err := io.ReadFull(reader, bodyBuff); err != nil {
		return nil, fmt.Errorf("read body err: %s", err)
	}
	data := bodyBuff // read body

	log.Debugf("Unpack: opcode[%d], %v", opcode, util.FormatHex(data))

	msg := easytcp.NewMessage(opcode, data)
	msg.Set("fullSize", dataSize+2+2)
	return msg, nil
}
