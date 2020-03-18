package main

import (
	"path/filepath"

	"github.com/MscBaiMeow/QQGroupHitokoto/config"
	"github.com/MscBaiMeow/QQGroupHitokoto/data"
	"github.com/MscBaiMeow/QQGroupHitokoto/users"
	"github.com/Tnze/CoolQ-Golang-SDK/cqp"
)

//go:generate cqcfg -c .
// cqp: 名称: Hitokoto
// cqp: 版本: 1.0.0:0
// cqp: 作者: BaiMeow
// cqp: 简介: 群一言
func main() { /*此处应当留空*/ }

func init() {
	cqp.AppID = "cn.miaoscraft.Hitokoto" // TODO: 修改为这个插件的ID
	cqp.GroupMsg = onGroupMsg
	cqp.Enable = onEnable
}

func onEnable() int32 {
	config.Load(filepath.Join(cqp.GetAppDir(), "conf.json"))
	return 0
}

func onGroupMsg(subType, msgID int32, fromGroup, fromQQ int64, fromAnonymous, msg string, font int32) int32 {
	if msg != "一言" {
		return 0
	}
	kind := users.GetType(fromGroup)
	hitokoto, err := data.GetHitokoto(kind)
	if err != nil {
		cqp.AddLog(cqp.Error, "一言", err.Error())
	}
	cqp.SendGroupMsg(fromGroup, hitokoto)
	return 0
}
