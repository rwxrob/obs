package cmd

import (
	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/requests/scenes"
	Z "github.com/rwxrob/bonzai/z"
)

var Scene = &Z.Cmd{
	Name:    `scene`,
	Summary: `change, create, delete scenes`,
	UseConf: true,
	Call: func(x *Z.Cmd, args ...string) error {
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

		params := &scenes.SetCurrentProgramSceneParams{args[0]}
		_, err = client.Scenes.SetCurrentProgramScene(params)
		return err
	},
}
