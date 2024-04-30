package account

type AuthRequest struct {
	UserId   string
	Password string
}

type AuthResponse struct {
	Id       int32
	AuthKey  uint32
	Status   byte
	CharList []CharCount
}

type GetAccountReq struct {
	UserId string
}

type GetAccountRes struct {
	Id     int32
	Status byte
}

type VerifyReq struct {
	AuthKey   uint32
	UserIdx   uint16
	ServerId  byte
	ChannelId byte
	IP        string
	DBIdx     int32
}

type VerifyRes struct {
	Verified bool
}

type AuthCheckReq struct {
	Id       int32
	Password string
}

type AuthCheckRes struct {
	Result bool
}

type CharCount struct {
	Server byte
	Count  byte
}

type OnlineReq struct {
	Account int32
	Kick    bool
}

type OnlineRes struct {
	Result bool
}
