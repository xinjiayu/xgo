package main

import (
	"xgo/vars"

	"github.com/xinjiayu/log"
)

func init() {
	vars.Init()
}

func main() {
	log.Debug("这是一个测试")
	mlist := vars.Engine.GetDevList("yy15260524156_CWSA04")
	log.Debug(mlist)
}
