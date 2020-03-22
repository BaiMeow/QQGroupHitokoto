package data

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/Tnze/CoolQ-Golang-SDK/cqp"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

type sqliteSource struct {
	name   string
	engine *xorm.Engine
}

func (a *sqliteSource) getName() string {
	return a.name
}

func (a *sqliteSource) getHitokoto() (string, error) {
	values := hitokoto{}
	has, err := a.engine.OrderBy("random()").Get(&values)
	if !has || err != nil {
		return "", err
	}
	return fmt.Sprintf("「%s」 - %s", values.Hitokoto, values.From), nil
}

func addSQLITESource(name, path string) error {
	if nameGetSource(name) != nil {
		return errors.New("Source " + name + " has existed")
	}
	engine, err := xorm.NewEngine("sqlite3", filepath.Join(cqp.GetAppDir(), path))
	if err != nil {
		return err
	}

	sourceList = append(sourceList, &sqliteSource{
		name:   name,
		engine: engine,
	})
	return nil
}
