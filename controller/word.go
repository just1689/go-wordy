package controller

import (
	"fmt"
	"github.com/just1689/go-wordy/mem"
	"github.com/just1689/go-wordy/model"
	"github.com/just1689/go-wordy/util"
	log "github.com/sirupsen/logrus"
)

func Exists(value string) (bool, error) {
	_, err := mem.GetUniqueById(model.WordTablename, value)
	if err == nil {
		return true, err
	}
	return false, err
}

func LoadAll() {
	ch := util.ReadFile(model.FileName)
	for item := range *ch {
		word := model.Word{ID: item}
		err := mem.Push(word)
		if err != nil {
			log.Errorln(fmt.Sprintf("There was a problem decoding the post message: %s", err))
		}
	}

}
