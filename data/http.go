package data

import (
	"encoding/json"
	"net/http"
)

type httpSource struct {
	url  string
	name string
}

func (a *httpSource) getHitokoto() (string, error) {
	resp, err := http.Get(a.url)
	if resp.StatusCode != 302 || err != nil {
		return "", err
	}
	var hitokoto hitokoto
	err = json.NewDecoder(resp.Body).Decode(&hitokoto)
	if err != nil {
		return "", err
	}
	return hitokoto.hitokoto + hitokoto.from, nil
}

func (a *httpSource) getName() string {
	return a.name
}
