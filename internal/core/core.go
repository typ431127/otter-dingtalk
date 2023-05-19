package core

import (
	"github.com/blinkbean/dingtalk"
	"github.com/go-zookeeper/zk"
	"otter-dingtalk/internal/global"
	"time"
)

func zkConn() *zk.Conn {
	conn, _, err := zk.Connect([]string{global.ZKADDR}, time.Second*5)
	if err != nil {
		global.GL_LOG.Fatalln(err)
	}
	return conn
}
func Init() {
	global.GL_LOG = Zap()
	global.GL_ZK = zkConn()
	global.GL_DB = Gorm()
	global.GL_DING = dingtalk.InitDingTalkWithSecret(global.DINGTOKEN, global.DINGSECRTE)
}
