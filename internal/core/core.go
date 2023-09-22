package core

import (
	"fmt"
	"github.com/blinkbean/dingtalk"
	"github.com/go-zookeeper/zk"
	"otter-dingtalk/internal/global"
	"otter-dingtalk/internal/otter"
	"otter-dingtalk/pkg/ext"
	"time"
)

func zkConn() *zk.Conn {
	conn, _, err := zk.Connect([]string{global.ZKADDR}, time.Second*60)
	if err != nil {
		global.GL_DING.SendTextMessage(fmt.Sprintf("ZK连接异常!\n环境:%s\n地址:%s\n", global.ENV, global.ZKADDR))
		global.GL_LOG.Fatalln(err)
	}
	if _, _, err := conn.Get("/"); err != nil {
		global.GL_DING.SendTextMessage(fmt.Sprintf("ZK连接异常!\n环境:%s\n地址:%s\n", global.ENV, global.ZKADDR))
		global.GL_LOG.Fatalln(err)
	}
	return conn
}
func Init() {
	global.GL_LOG = Zap()
	global.GL_DING = dingtalk.InitDingTalkWithSecret(global.DINGTOKEN, global.DINGSECRTE)
	global.GL_ZK = zkConn()
	global.GL_DB = Gorm()
	global.GL_HOSTNAME = ext.Hostname()
	global.GL_IP = ext.InterfaceAddrs()
	otter.Login()
}
