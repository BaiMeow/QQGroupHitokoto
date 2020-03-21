package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/MscBaiMeow/QQGroupHitokoto/cq"
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
	var sum int
	for _, v := range conf.Sources {
		if err := data.AddSource(v.Name, v.SourceType, v.Source); err != nil {
			cq.Info(fmt.Sprintf("已添加%d个一言源", sum))
			return err
		}
		sum = sum + 1
	}
	cq.Info(fmt.Sprintf("已添加%d个一言源", sum))
	return nil
}
