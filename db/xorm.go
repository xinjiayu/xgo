package db

import (
	"strconv"
	"xgo/config"
	"xgo/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/xinjiayu/log"
)

//Engine 定义Engine struct
type Engine struct {
	Xorm *xorm.Engine
}

//NewEngine 初始化Engine
func NewEngine() *Engine {
	return &Engine{Xorm: newXorm()}
}

func newXorm() *xorm.Engine {
	var e *xorm.Engine
	e, err := xorm.NewEngine("mysql", config.Params["mysql"]["user"]+":"+config.Params["mysql"]["pwd"]+"@tcp("+config.Params["mysql"]["host"]+")/"+config.Params["mysql"]["db"]+"?charset=utf8&loc=Asia%2FShanghai")

	if err != nil {
		log.Error("新建引擎：", err)
	}

	err = e.Ping()
	if err != nil {
		log.Error("数据库连接：", err)
		// os.Exit(1)
	}

	maxopenconns, _ := strconv.Atoi(config.Params["mysql"]["maxopenconns"])
	e.SetMaxOpenConns(maxopenconns) //设置最大打开连接数
	maxidleconns, _ := strconv.Atoi(config.Params["mysql"]["maxidleconns"])
	e.SetMaxIdleConns(maxidleconns) //设置连接池的空闲数大小

	e.ShowSQL(config.Params["logLevel"]["ShowSQL"] == "true")

	return e
}

// GetDevNow 获取指定设备的位置数据
func (e *Engine) GetDevList(dnum string) []model.MapDevicegps {
	var slice []model.MapDevicegps
	e.Xorm.Where("num = ?", dnum).Find(&slice)
	return slice
}
