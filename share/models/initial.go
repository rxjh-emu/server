package models

type InitialJob struct {
	ID    int `json:"id"`
	Stats struct {
		Hp  int `json:"hp"`
		Mp  int `json:"mp"`
		Rp  int `json:"rp"`
		Spr int `json:"spr"`
		Str int `json:"str"`
		Stm int `json:"stm"`
		Dex int `json:"dex"`
	}
	Location struct {
		Map int     `json:"map"`
		X   float64 `json:"x"`
		Y   float64 `json:"y"`
		Z   float64 `json:"z"`
	}
	Inventory []struct {
		Item     int `json:"item"`
		Quantity int `json:"quantity"`
	}
}
