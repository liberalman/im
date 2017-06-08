package global

import (
	"github.com/liberalman/im/libnet"
)

type SessionMap map[int64]*libnet.Session

var GSessions SessionMap

func init() {
	GSessions = make(SessionMap)
}
