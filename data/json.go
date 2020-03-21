package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
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
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	var body []hitokoto
	err = json.Unmarshal(file, &body)
	if err != nil {
		return err
	}
	var jsonSource = jsonSource{
		name: name,
		body: body,
	}
	sourceList = append(sourceList, &jsonSource)
	return nil
}
