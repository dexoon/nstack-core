package api

import (
	"github.com/xtls/xray-core/main/commands/base"
)

// CmdAPI calls an API in an svc process
var CmdAPI = &base.Command{
	UsageLine: "svc api",
	Short:     "Call an API in an svc process",
	Long: `svc svc provides tools to manipulate svc via its API.
`,
	Commands: []*base.Command{
		cmdRestartLogger,
		cmdGetStats,
		cmdQueryStats,
		cmdSysStats,
		cmdBalancerInfo,
		cmdBalancerOverride,
		cmdAddInbounds,
		cmdAddOutbounds,
		cmdRemoveInbounds,
		cmdRemoveOutbounds,
		cmdListInbounds,
		cmdListOutbounds,
		cmdAddInboundUsers,
		cmdRemoveInboundUsers,
		cmdInboundUser,
		cmdInboundUserCount,
		cmdAddRules,
		cmdRemoveRules,
		cmdListRules,
		cmdSourceIpBlock,
		cmdOnlineStats,
		cmdOnlineStatsIpList,
		cmdGetAllOnlineUsers,
	},
}
