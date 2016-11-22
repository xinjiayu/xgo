package vars

import (
	"xgo/config"
	"xgo/db"

	"github.com/xinjiayu/log"
	"github.com/xinjiayu/mylog"
)

var (
	Engine *db.Engine
)

func Init() {

	config.Load("conf/config.conf")
	mylog.LogInit()
	Engine = db.NewEngine()
	log.Debug("初始化完成")

}
