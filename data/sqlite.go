package data

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/Tnze/CoolQ-Golang-SDK/cqp"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type sqliteSource struct {
	name string
	db   *gorm.DB
}

func (a *sqliteSource) getName() string {
	return a.name
}

func (a *sqliteSource) getHitokoto() (string, error) {
	var values hitokoto
	err := a.db.Order("random()").First(&values).Error
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("「%s」 - %s", values.Hitokoto, values.From), nil
}

func addSQLITESource(name, path string) error {
	if nameGetSource(name) != nil {
		return errors.New("Source " + name + " has existed")
	}
	db, err := gorm.Open("sqlite3", filepath.Join(cqp.GetAppDir(), path))
	db.SingularTable(true)
	if err != nil {
		return err
	}

	sourceList = append(sourceList, &sqliteSource{
		name: name,
		db:   db,
	})
	return nil
}
