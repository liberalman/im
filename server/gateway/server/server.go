package server

import (
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"github.com/libebralman/im/common/ecode"
	"github.com/libebralman/im/conf_discovery/etcd"
	"github.com/libebralman/im/libnet"
	"github.com/libebralman/im/protocol/external"
	"github.com/libebralman/im/server/gateway/client"
)

type Server struct {
	Server        *libnet.Server
	Master        *etcd.Master //etcd服务器主机节点
	MsgServerList []*etcd.Member
}

func New() (s *Server) {
	s = &Server{}
	return
}

func (s *Server) sessionLoop(client *client.Client) {
	for {
		reqData, err := client.Session.Receive() // 阻塞，等待接收客户端发来的数据
		if err != nil {
			glog.Error(err)
		}
		if reqData != nil {
			baseCMD := &external.Base{}
			if err = proto.Unmarshal(reqData, baseCMD); err != nil { // 将接收到的数据，按照protobuf协议反序列化，放到baseCMD中
				if err = client.Session.Send(&external.Error{
					Cmd:     external.ErrServerCMD,
					ErrCode: ecode.ServerErr.Uint32(),
					ErrStr:  ecode.ServerErr.String(),
				}); err != nil {
					glog.Error(err)
				}
				continue
			}
			if err = client.Parse(baseCMD.Cmd, reqData); err != nil { // 解析
				glog.Error(err)
				continue
			}
		}
	}
}

func (s *Server) Loop() {
	glog.Info("loop")
	for {
		session, err := s.Server.Accept() // 阻塞，等待客户端连接；当获取到新的客户端连接上来的时候，创建session，使用协程处理!
		if err != nil {
			glog.Error(err)
		}
		go s.sessionLoop(client.New(session))
	}
}
