package cmd

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/conf"
	"github.com/rwxrob/help"
	"github.com/rwxrob/vars"
)

var OBS = &Z.Cmd{
	Name: "obs",
	Commands: []*Z.Cmd{
		help.Cmd, vars.Cmd, conf.Cmd,
		Start, Scenes, Scene,
	},
}
