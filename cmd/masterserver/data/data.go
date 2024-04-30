package data

import (
	"encoding/json"
	"github.com/rxjh-emu/server/share/directory"
	"github.com/rxjh-emu/server/share/log"
	"io/ioutil"
)

type Loader struct {
	*InitialData
}

type InitialData struct {
	Jobs []struct {
		ID       int `json:"id"`
		location struct {
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
}

// Initializes DataLoader
func (dl *Loader) Init() {
	log.Info("Loading data...")

	dl.InitialData = &InitialData{}
	dl.load("initial_data.json", dl.InitialData)
}

func (dl *Loader) load(filename string, data interface{}) {
	var s, err = ioutil.ReadFile(directory.Root() + "/data/" + filename)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = json.Unmarshal([]byte(s), data)
	//err = yaml.Unmarshal(s, data)
	if err != nil {
		log.Fatal(err.Error())
	}
}
