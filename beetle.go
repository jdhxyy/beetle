// Copyright 2022-2022 The jdh99 Authors. All rights reserved.
// 命令行解析器
// Authors: jdh99 <jdh821@163.com>

package beetle

import (
	"github.com/jdhxyy/lagan"
	"strings"
)

const (
	tag           = "beetle"
	cmdLineLenMax = 128
)

// CmdType 命令类型
type CmdType uint8

const (
	// TypeCmd 命令字
	TypeCmd CmdType = 0
	// TypeShortOption 短选项
	TypeShortOption CmdType = 1
	// TypeLongOption 长选项
	TypeLongOption CmdType = 2
	// TypeParam 参数.子命令也是参数
	TypeParam CmdType = 3
)

// CmdItem 元素
type CmdItem struct {
	Type  CmdType
	Value string
}

var gObservers map[string]CmdCallback
var gItems []CmdItem

// CmdCallback 命令回调函数.返回的是应答字符串,如果为空表示不需要回复
type CmdCallback func() string

func init() {
	gObservers = make(map[string]CmdCallback)
}

// Register 注册命令字对应的回调函数
func Register(cmd string, callback CmdCallback) {
	gObservers[cmd] = callback
}

// GetItemNum 获取条目数
// 用户可在回调函数中调用本函数获取命令行条目
func GetItemNum() int {
	return len(gItems)
}

// GetItem 获取制定条目
// 用户可在回调函数中调用本函数获取序号对应的命令行条目,注意序号index从0开始
// 如果不存在条目,则返回nil
func GetItem(index int) *CmdItem {
	if len(gItems) < index {
		return nil
	}
	return &gItems[index]
}

// Parse 解析命令行.返回的是应答字符串,如果为空表示不需要回复
func Parse(cmdLine string) string {
	lagan.Debug(tag, "parse cmd line:%s", cmdLine)

	if isCmdLineValid(cmdLine) == false {
		lagan.Error(tag, "cmd line is invalid:%s", cmdLine)
		return ""
	}

	gItems = []CmdItem{}
	params := strings.Split(cmdLine, " ")
	items := make([]CmdItem, len(params))
	for i := 0; i < len(params); i++ {
		if i == 0 {
			items[i].Type = TypeCmd
			items[i].Value = params[i]
			continue
		}

		if strings.HasPrefix(params[i], "--") {
			if len(params[i]) <= 2 {
				lagan.Error(tag, "params is too short")
				return ""
			}
			items[i].Type = TypeLongOption
			items[i].Value = params[i][2:]
			continue
		}

		if strings.HasPrefix(params[i], "-") {
			if len(params[i]) <= 1 {
				lagan.Error(tag, "params is too short")
				return ""
			}
			items[i].Type = TypeShortOption
			items[i].Value = params[i][1:]
			continue
		}

		items[i].Type = TypeParam
		items[i].Value = params[i]
	}

	gItems = items
	callback, ok := gObservers[gItems[0].Value]
	if ok == false {
		lagan.Error(tag, "can not find callback")
		return ""
	}
	return callback()
}

func isCmdLineValid(cmdLine string) bool {
	cmdLen := len(cmdLine)
	if cmdLen == 0 || cmdLen > cmdLineLenMax {
		return false
	}

	// 第一个以及最后一个字符为空格则是格式错误
	if cmdLine[0] == ' ' || cmdLine[cmdLen-1] == ' ' {
		return false
	}

	// 字符只允许可显示和可输入字符
	// 不允许连续两个空格
	last := uint8(0)
	for i := 0; i < cmdLen; i++ {
		if cmdLine[i] < 0x20 || cmdLine[i] > 0x7E {
			return false
		}
		if cmdLine[i] == ' ' && last == ' ' {
			return false
		}
		last = cmdLine[i]
	}
	return true
}
