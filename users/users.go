package users

import (
	"errors"

	"github.com/MscBaiMeow/QQGroupHitokoto/data"
)

//ErrUnkownHitokoto 无效的一言类型
var ErrUnkownHitokoto = errors.New("无效一言类型")

//type没得用，不想写Type，那就kind吧

//SetType 设置获取的一言类型
func SetType(group int64, kind string) error {
	if kind == "every" {
		delete(selectedKind, group)
		return nil
	}
	if !data.ValidateType(kind) {
		return ErrUnkownHitokoto
	}
	selectedKind[group] = kind
	return nil
}

//GetType 获取设置或者说是随机抽出来的的一言类型
func GetType(group int64) string {
	kind := selectedKind[group]
	if kind == "" {
		return data.RandSourceName()
	}
	return kind
}

var selectedKind = make(map[int64]string)
