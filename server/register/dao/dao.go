package dao

import (
	// "github.com/golang/glog"
	"github.com/liberalman/im/common/dao/xredis"
	"github.com/liberalman/im/server/register/conf"
)

type Dao struct {
	redis *xredis.Pool
	Mysql *Mysql
}

func NewDao() (dao *Dao) {
	mysql := NewMysql()
	dao = &Dao{
		redis: xredis.NewPool(conf.Conf.Redis.Redis),
		Mysql: mysql,
	}
	return
}
