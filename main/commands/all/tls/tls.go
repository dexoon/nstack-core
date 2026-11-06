package tls

import (
	"github.com/xtls/xray-core/main/commands/base"
)

// CmdTLS holds all tls sub commands
var CmdTLS = &base.Command{
	UsageLine: "svc tls",
	Short:     "TLS tools",
	Long: `svc svc provides tools for TLS.
`,
	Commands: []*base.Command{
		cmdCert,
		cmdPing,
		cmdHash,
		cmdECH,
	},
}
