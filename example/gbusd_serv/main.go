package main

import (
	"fmt"
	"github.com/evenzhang/gbusd/gbusd"
	"github.com/evenzhang/gbusd/gbusd/res"
	"github.com/evenzhang/gbusd/gserv/application"
)

func main() {
	App := application.NewApplication(application.ProjConstInfo{
		Name:            res.RES_PROJ_NAME,
		Desc:            res.RES_PROJ_DESC,
		Version:         res.RES_PROJ_VERSION,
		DefaultConfPath: res.GetDefaultConf("serv")})

	daemon := gbusd.NewDaemonServer(NewEventHandler())
	daemon.ParserFlag()

	opts := App.EnableAndParserFlagCmd(res.RES_PROJ_USR_USAGE)
	if err := App.LoadAndMergeConfigFile(opts); err != nil {
		fmt.Println(err)
	}

	App.Init(daemon)
	App.Run()
}
