package rpc

import (
	"strings"

	"github.com/rxjh-emu/server/share/models/account"
	"github.com/rxjh-emu/server/share/rpc"
)

// AuthCheck RPC Call
func AuthCheck(c *rpc.Client, r *account.AuthRequest, s *account.AuthResponse) error {
	var res = account.AuthResponse{Status: account.None}
	var passHash string

	var err = g_LoginDatabase.Get(&passHash,
		"SELECT password FROM accounts WHERE username = ?", r.UserId)

	if err != nil {
		res.Status = account.AccountNotFound
		*s = res
		return nil
	}

	if strings.ToLower(passHash) != r.Password {
		res.Status = account.WrongPassword
	} else {
		g_LoginDatabase.Get(&res,
			"SELECT id FROM accounts WHERE username = ?", r.UserId)

		// check double login
		if g_ServerManager.IsOnline(res.Id) {
			res.Status = account.Onlined
			*s = res
			return nil
		}

		var count = account.CharCount{}

		// fetch char list on all of servers
		for _, value := range g_DatabaseManager.DBList {
			value.DB.Get(&count.Count,
				"SELECT COUNT(id) FROM characters WHERE account = ?",
				res.Id,
			)

			if count.Count > 0 {
				count.Server = byte(value.Index)
			}

			res.CharList = append(res.CharList, count)
		}

		res.Status = account.Success
	}

	*s = res
	return nil
}

func GetAccount(c *rpc.Client, r *account.GetAccountReq, s *account.GetAccountRes) error {
	var res = account.GetAccountRes{Status: account.None}
	var err = g_LoginDatabase.Get(&res, "SELECT id FROM accounts WHERE username = ?", r.UserId)
	if err != nil {
		s.Status = account.ErrorMessage
	} else {
		s.Id = res.Id
		s.Status = account.Success
	}

	// *s = res
	return nil
}
