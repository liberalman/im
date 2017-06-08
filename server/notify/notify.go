package main

import (
	"flag"

	"github.com/golang/glog"
	"github.com/liberalman/im/server/notify/conf"
	"github.com/liberalman/im/server/notify/conf_discovery"
	"github.com/liberalman/im/server/notify/rpc"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		glog.Error("conf.Init() error: ", err)
		panic(err)
	}
	go conf_discovery.ConfDiscoveryProc()
	rpcClient, err := rpc.NewRPCClient()
	if err != nil {
		glog.Error(err)
		panic(err)
	}
	rpc.RPCServerInit(rpcClient)
}
