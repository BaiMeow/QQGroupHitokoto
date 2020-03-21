package data

import (
	"errors"
	"math/rand"
)

type hitokoto struct {
	//ID         int32     `json:"id"`
	Hitokoto string `json:"hitokoto"`
	//Kind       string    `json:"type"`
	From string `json:"from"`
	//FromWho    string    `json:"from_who"`
	//Creator    string    `json:"creator"`
	//CreatorUID int32     `json:"creator_uid"`
	//Reviewer   int32     `json:"reviewer"`
	//UUID       uuid.UUID `json:"uuid"`
	//CreatedAt  string    `json:"created_at"`
}

type source interface {
	getName() string
	getHitokoto() (string, error)
}

var sourceList []source

//ValidateType 验证数据源是否存在
func ValidateType(kind string) bool {
	return nameGetSource(kind) != nil
}

//AddSource 添加源
func AddSource(name, sourceType, source string) error {
	switch sourceType {
	case "HTTP":
		if err := addHTTPSource(name, source); err != nil {
			return err
		}
	case "JSON":
		if err := addJSONSource(name, source); err != nil {
			return err
		}
	case "SQLITE":
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
