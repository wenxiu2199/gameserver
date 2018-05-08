package internal

import (
	"server/manager"

	"github.com/name5566/leaf/chanrpc"
)

var sessionMgr *sessionManager
var serverMgr *manager.ServerManager

func Init(servers map[manager.ServerType]*chanrpc.Server) {
	sessionMgr = &sessionManager{}
	sessionMgr.init()

	serverMgr = &manager.ServerManager{}
	serverMgr.Init(servers)
}