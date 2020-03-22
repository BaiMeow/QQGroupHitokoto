package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"path/filepath"

	"github.com/Tnze/CoolQ-Golang-SDK/cqp"
)

type jsonSource struct {
	name string
	body []hitokoto
}

func (a *jsonSource) getHitokoto() (string, error) {
	if len(a.body) == 0 {
		return "", errors.New("No existed hitokoto")
	}
	hitokoto := a.body[rand.Intn(len(a.body))]
	return fmt.Sprintf("「%s」 - %s", hitokoto.Hitokoto, hitokoto.From), nil
}

func (a *jsonSource) getName() string {
	return a.name
}

func addJSONSource(name, path string) error {
	if nameGetSource(name) != nil {
		return errors.New("Source " + name + " has existed")
	}
	file, err := ioutil.ReadFile(filepath.Join(cqp.GetAppDir(), path))
	if err != nil {
		return err
	}
	var body []hitokoto
	err = json.Unmarshal(file, &body)
	if err != nil {
		return err
	}

	sourceList = append(sourceList, &jsonSource{
		name: name,
		body: body,
	})
	return nil
}
