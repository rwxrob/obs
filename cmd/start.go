package cmd

import Z "github.com/rwxrob/bonzai/z"

var Start = &Z.Cmd{
	Name:    `start`,
	Summary: `start a connection to an OBS Websocket server`,
	Call: func(x *Z.Cmd, _ ...string) error {
		println(`hello`)
		return nil
	},
}
