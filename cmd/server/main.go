package main

import (
	"fmt"

	common "github.com/KoruptTinker/korupt-monitor/cmd/common/init"
)

func main() {
	engine, config := common.InitServer()

	engine.Run(fmt.Sprintf("0.0.0.0:%d", config.Server.Port))
}
