package dao

import (
	"github.com/golang/glog"
	"github.com/liberalman/im/service/idgen/conf"
)

type Dao struct {
	Etcd *Etcd
}

func NewDao() (dao *Dao, err error) {
	e, err := NewEtcd(conf.Conf.Etcd)
	if err != nil {
		glog.Error(err)
		return
	}
	dao = &Dao{
		Etcd: e,
	}
	return
}
