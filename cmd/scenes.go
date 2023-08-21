package cmd

import (
	"fmt"

	"github.com/andreykaipov/goobs"
	Z "github.com/rwxrob/bonzai/z"
)

var Scenes = &Z.Cmd{
	Name:     `scenes`,
	Summary:  `list all the scenes`,
	UseConf:  true,
	UseVars:  true,
	ConfVars: true,
	Call: func(x *Z.Cmd, _ ...string) error {
		// TODO change to current server setting
		addr, err := x.Caller.C(`servers.default.address`)
		if err != nil {
			return err
		}
		port := "4455"
		if p, err := x.Caller.C(`servers.default.port`); err != nil {
			port = p
		}
		pass, _ := x.Caller.C(`servers.default.password`)

		client, err := goobs.New(addr+`:`+port, goobs.WithPassword(pass))
		if err != nil {
			return err
		}
		defer client.Disconnect()

		// TODO assert compatible client version
		//version, err := client.General.GetVersion()
		//if err != nil {
		//log.Fatal(err)
		//}
		//fmt.Printf("OBS Studio version: %s\n", version.ObsVersion)
		//fmt.Printf("Websocket server version: %s\n", version.ObsWebSocketVersion)
		resp, _ := client.Scenes.GetSceneList()
		for _, v := range resp.Scenes {
			fmt.Printf("%s\n", v.SceneName)
		}
		return nil
	},
}
