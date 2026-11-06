package convert

import (
	"github.com/xtls/xray-core/main/commands/base"
)

// CmdConvert do config convertion
var CmdConvert = &base.Command{
	UsageLine: "svc convert",
	Short:     "Convert configs",
	Long: `svc svc provides tools to convert config.
`,
	Commands: []*base.Command{
		cmdProtobuf,
		cmdJson,
	},
}
