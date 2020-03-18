package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/MscBaiMeow/QQGroupHitokoto/data"
)

type source struct {
	name       string
	sourceType string
	source     string
}

type config struct {
	sources []source
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
	for _, v := range conf.sources {
		if err := data.AddSource(v.name, v.sourceType, v.source); err == nil {
			return err
		}
	}
	return nil
}
