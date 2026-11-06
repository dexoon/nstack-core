package lua

import "github.com/xtls/xray-core/main/commands/base"

// CmdLua holds the Lua script commands.
var CmdLua = &base.Command{
	UsageLine: "svc lua",
	Short:     "Manage Lua scripts",
	Long:      "svc svc provides tools for local Lua scripts.",
	Commands: []*base.Command{
		cmdSign,
	},
}
