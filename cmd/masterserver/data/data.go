package data

import (
	"encoding/json"
	"io/ioutil"

	"github.com/rxjh-emu/server/share/directory"
	"github.com/rxjh-emu/server/share/log"
	"github.com/rxjh-emu/server/share/models"
)

type Loader struct {
	*InitialData
}

type InitialData struct {
	Jobs []models.InitialJob
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
