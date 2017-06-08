package conf_discovery

import (
	"time"

	"github.com/golang/glog"
	"github.com/liberalman/im/conf_discovery/etcd"
	"github.com/liberalman/im/server/notify/conf"
)

var (
	AccessServerList map[string]*etcd.Member
)

func loadAccessServerProc(master *etcd.Master) {
	for {
		// glog.Info("loadAccessServerProc")
		AccessServerList = master.Members()
		time.Sleep(time.Second * 5)
	}
}

func ConfDiscoveryProc() {
	glog.Info("ConfDiscoveryProc")
	master, err := etcd.NewMaster(conf.Conf.Etcd)
	if err != nil {
		glog.Error("Error: cannot connect to etcd:", err)
		panic(err)
	}
	go loadAccessServerProc(master)
	master.WatchWorkers()
}
