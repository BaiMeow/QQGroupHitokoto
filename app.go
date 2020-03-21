package main

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/MscBaiMeow/QQGroupHitokoto/config"
	"github.com/MscBaiMeow/QQGroupHitokoto/cq"
	"github.com/MscBaiMeow/QQGroupHitokoto/data"
	"github.com/MscBaiMeow/QQGroupHitokoto/users"
	"github.com/Tnze/CoolQ-Golang-SDK/cqp"
)

//go:generate cqcfg -c .
// cqp: 名称: 一言Hitokoto
// cqp: 版本: 1.0.0:0
// cqp: 作者: BaiMeow
// cqp: 简介: 一言
func main() { /*此处应当留空*/ }

func init() {
	cqp.AppID = "cn.miaoscraft.Hitokoto" // TODO: 修改为这个插件的ID
	cqp.GroupMsg = onGroupMsg
	cqp.Enable = onEnable
}

func onEnable() int32 {
	err := config.Load(filepath.Join(cqp.GetAppDir(), "conf.json"))
	if err != nil {
		cq.Error(err)
	}
	return 0
}

func onGroupMsg(subType, msgID int32, fromGroup, fromQQ int64, fromAnonymous, message string, font int32) int32 {
	//判断是否接收处理
	msg := []rune(message)
	if len(msg) < 2 || string(msg[:2]) != "一言" {
		cq.Info(len(msg), msg[:2])
		return 0
	}
	//普通一言
	if len(msg) == 2 {
		kind := users.GetType(fromGroup)
		hitokoto, err := data.GetHitokoto(kind)
		if err != nil {
			cq.Error(err)
		}
		cqp.SendGroupMsg(fromGroup, hitokoto)
		return 1
	}
	//设置一言类型
	if len(msg) > 3 {
		var kind string
		fmt.Sscanf(message, "一言 %s", &kind)
		err := users.SetType(fromGroup, kind)
		//错误处理
		if err != nil {
			if errors.Is(err, users.ErrUnkownHitokoto) {
				cqp.SendGroupMsg(fromGroup, "无效一言类型")
				return 1
			}
			cq.Error(err)
		}
		cqp.SendGroupMsg(fromGroup, "成功设置一言类型为"+kind)
		return 1
	}
	return 0
}
