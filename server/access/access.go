package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/liberalman/im/codec"
	"github.com/liberalman/im/libnet"
	"github.com/liberalman/im/server/access/conf"
	"github.com/liberalman/im/server/access/rpc"
	"github.com/liberalman/im/server/access/server"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

func main() {
	var err error
	flag.Parse()
	if err = conf.Init(); err != nil {
		glog.Error("conf.Init() error: ", err)
		panic(err)
	}
	accessServer := server.New()
	protobuf := codec.Protobuf()
	if accessServer.Server, err = libnet.Serve(conf.Conf.Server.Proto, conf.Conf.Server.Addr, protobuf, 0); err != nil {
		glog.Error(err)
		panic(err)
	}
	go accessServer.SDHeart()
	rpcClient, err := rpc.NewRPCClient()
	if err != nil {
		glog.Error(err)
		panic(err)
	}
	accessServer.Loop(rpcClient)
}
