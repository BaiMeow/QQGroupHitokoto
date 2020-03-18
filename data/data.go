package data

import (
	"errors"
	"math/rand"

	"github.com/google/uuid"
)

type hitokoto struct {
	id         int32
	hitokoto   string
	kind       string
	from       string
	fromWho    string
	creator    string
	creatorUID int32
	reviewer   int32
	uuid       uuid.UUID
	createdAt  int64
}

type source interface {
	getName() string
	getHitokoto() (string, error)
}

var sourceList []source

func addHTTPSource(name, url string) error {
	if nameGetSource(name) != nil {
		return errors.New("Source" + name + "has existed")
	}
	var httpSource = httpSource{
		url: url,
	}
	sourceList = append(sourceList, &httpSource)
	return nil
}

//ValidateType 验证数据源是否存在
func ValidateType(kind string) bool {
	return nameGetSource(kind) != nil
}

//AddSource 添加源
func AddSource(name, sourceType, source string) error {
	switch sourceType {
	case "http":
		addHTTPSource(name, source)
	case "json":
	case "sqlite":
	default:
		return errors.New("No such sourceType")
	}
	return nil
}

//GetHitokoto 获取一言
func GetHitokoto(kind string) (string, error) {
	hitokoto, err := nameGetSource(kind).getHitokoto()
	if err != nil {
		return "", err
	}
	return hitokoto, nil
}

func nameGetSource(name string) source {
	for _, v := range sourceList {
		if v.getName() == name {
			return v
		}
	}
	return nil
}

//IsSourceExist 判断源是否存在
func IsSourceExist(name string) bool {
	for _, v := range sourceList {
		if v.getName() == name {
			return true
		}
	}
	return false
}

//RandSourceName 返回随机一个源的名称
func RandSourceName() string {
	return sourceList[rand.Intn(len(sourceList))].getName()
}
