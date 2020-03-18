package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/MscBaiMeow/QQGroupHitokoto/data"
)

type source struct {
	Name       string `json:"Name"`
	SourceType string `json:"SourceType"`
	Source     string `json:"Source"`
}

type config struct {
	Sources []source `json:"Sources"`
}

var conf config

//Load 读取配置文件
func Load(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &conf)
	if err != nil {
		return err
	}
	for _, v := range conf.Sources {
		if err := data.AddSource(v.Name, v.SourceType, v.Source); err == nil {
			return err
		}
	}
	return nil
}
