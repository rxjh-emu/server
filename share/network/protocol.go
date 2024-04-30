package network

type Protocol interface {
	Read([]byte) *PacketArgs
	Write([]byte) *[]byte
	GetHeaderSize() int
	SetUserIdx(int)
}
