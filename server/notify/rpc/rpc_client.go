package rpc

import (
	"github.com/golang/glog"
	"github.com/liberalman/im/server/notify/rpc/client"
)

type RPCClient struct {
	Access *client.AccessServerRPCCli
}

func NewRPCClient() (c *RPCClient, err error) {
	access, err := client.NewAccessServerRPCCli()
	if err != nil {
		glog.Error(err)
		return
	}
	c = &RPCClient{
		Access: access,
	}
	return
}
