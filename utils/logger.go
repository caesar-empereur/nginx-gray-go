package utils

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func GetLogger() *logs.BeeLogger {

	log := logs.NewLogger()

	config := make(map[string]interface{})
	config["filename"] = beego.AppConfig.String("log_path")
	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("initLogger failed, marshal err:", err)
	}
	log.SetLogger(logs.AdapterFile, string(configStr))
	log.SetLevel(logs.LevelDebug)
	log.EnableFuncCallDepth(true)
	return log
}
