package network

type PacketArgs struct {
	Session      *Session
	UserIdx      int
	PacketLength int
	Length       int
	Type         int
	Data         []byte

	Reader *Reader
}
