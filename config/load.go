// load
package config

import (
	"github.com/xinjiayu/log"

	"gopkg.in/ini-1"
)

var (
	Params = make(map[string]map[string]string)
)

func Load(filePath string) {
	cfg, err := ini.Load(filePath)
	if err != nil {
		log.Errorf("没有找到配置文件 $s ,或配置文件没有与文件同级目录.$s", filePath, err)
		return
	}
	names := cfg.SectionStrings()
	if len(names) > 0 {
		for _, v := range names {
			section, err := cfg.GetSection(v)
			if err != nil {
				return
			}
			maps := section.KeysHash()
			Params[v] = maps
		}
	}

}
