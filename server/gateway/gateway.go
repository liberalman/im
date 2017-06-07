package main

import (
	"flag"

	"github.com/golang/glog"
	"github.com/liberalman/im/codec"
	"github.com/liberalman/im/libnet"
	"github.com/liberalman/im/server/gateway/conf"
	"github.com/liberalman/im/server/gateway/job"
	"github.com/liberalman/im/server/gateway/server"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

func main() {
	var err error
	flag.Parse()
	if err = conf.Init(); err != nil { // 载入配置文件gateway.toml进行初始化
		glog.Error("conf.Init() error: ", err)
		panic(err)
	}
	gwServer := server.New()                                                                                         // gateway的server对象，然后下面的libnetServe负责创建server，监听端口
	protobuf := codec.Protobuf()                                                                                     // protobuf协议解析
	if gwServer.Server, err = libnet.Serve(conf.Conf.Server.Proto, conf.Conf.Server.Addr, protobuf, 0); err != nil { //创建server，监听端口
		glog.Error(err)
		panic(err)
	}
	go job.ConfDiscoveryProc() //启一个协程监控etcd服务，有配置更新的消息推送则更新本地配置
	gwServer.Loop()            //服务循环
}
