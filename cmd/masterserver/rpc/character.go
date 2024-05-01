package rpc

import (
	"github.com/jmoiron/sqlx"
	"github.com/rxjh-emu/server/share/log"
	"github.com/rxjh-emu/server/share/models"
	"github.com/rxjh-emu/server/share/models/character"
	"github.com/rxjh-emu/server/share/models/world"
	"github.com/rxjh-emu/server/share/rpc"
	"github.com/samber/lo"
)

// Load characters RPC call
func LoadCharacters(c *rpc.Client, r *character.ListReq, s *character.ListRes) error {
	var db = g_DatabaseManager.Get(r.Server)
	var res = character.ListRes{List: make([]character.Character, 0)}

	if db == nil {
		*s = res
		return nil
	}

	// get characters from database
	// var characters = make([]character.Character, 0)
	// db.Select(&characters, "SELECT * FROM characters WHERE account = ? ORDER BY slot ASC", r.Account)
	// res.Characters = characters

	var rows, err = db.Queryx(
		"SELECT "+
			"id, name, slot, class, level, ki, spr, str, stm, dex, fame, "+
			"morals, hp, mp, rp, gender, hair, hair_color, face, voice "+
			"FROM characters "+
			"WHERE account = ?", r.Account)

	if err != nil {
		log.Error("[DATABASE]", err)
		*s = res
		return nil
	}

	// iterate over each row
	for rows.Next() {
		var c = character.Character{}
		var err = rows.StructScan(&c)

		if err == nil {
			c.Position = loadPosition(db, c.Id)
			res.List = append(res.List, c)
		} else {
			log.Error("[DATABASE]", err)
			*s = res
			return nil
		}
	}

	*s = res
	return nil
}

// local function to load character position
func loadPosition(db *sqlx.DB, id int) world.Position {
	var position = world.Position{}
	var err = db.Get(&position,
		"SELECT map, x, y, z "+
			"FROM characters "+
			"WHERE id = ?", id)

	if err != nil {
		log.Error("[DATABASE]", err)
	}

	return position
}

// CheckCharacterName RPC Call
func CheckCharacterName(c *rpc.Client, r *character.CheckNameReq, s *character.CheckNameRes) error {
	var db = g_DatabaseManager.Get(r.Server)
	var res = character.CheckNameRes{}

	var name = 0
	db.Get(&name, "SELECT id FROM characters WHERE name = ? LIMIT 1", r.Name)
	if name > 0 {
		res.Result = character.NameInUse
		*s = res
		return nil
	}

	res.Result = character.NameCanUse
	*s = res
	return nil
}

// CreateCharacter RPC Call
func CreateCharacter(c *rpc.Client, r *character.CreateCharacterReq, s *character.CreateCharacterRes) error {
	var db = g_DatabaseManager.Get(r.Server)
	var res = character.CreateCharacterRes{}

	// get initial data
	init, _ := lo.Find(g_DataLoader.Jobs, func(j models.InitialJob) bool {
		return j.ID == r.Class
	})

	// get free slot from characters table by function
	var slot = 0
	db.Get(&slot, "SELECT getCharacterFreeSlot(?)", r.AccountId)

	log.Debugf("CreateCharacter slot: %d", slot)

	// TODO: create character
	var sql = "INSERT INTO characters ("
	sql += "account, name, slot, class, level, ki, spr, str, stm, dex, fame, "
	sql += "morals, hp, mp, rp, gender, hair, hair_color, face, voice, map, x, y, z"
	sql += ") VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	result := db.MustExec(sql,
		r.AccountId,
		r.Name,
		slot,
		r.Class,
		1,
		0,
		init.Stats.Spr,
		init.Stats.Str,
		init.Stats.Stm,
		init.Stats.Dex,
		0,
		0,
		init.Stats.Hp,
		init.Stats.Mp,
		init.Stats.Rp,
		r.Gender,
		r.Hair,
		r.HairColor,
		r.Face,
		r.Voice,
		init.Location.Map,
		init.Location.X,
		init.Location.Y,
		init.Location.Z,
	)
	lastID, _ := result.LastInsertId()
	log.Debugf("CreateCharacter result: %d", lastID)

	// TODO: create equipment
	// TODO: create inventory
	// TODO: create ability

	if lastID > 0 {
		res.Result = character.CreateCharacterSuccess
	} else {
		res.Result = character.CreateCharacterFail
	}

	*s = res
	return nil
}
