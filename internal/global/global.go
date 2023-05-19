package global

import (
	"github.com/blinkbean/dingtalk"
	"github.com/go-zookeeper/zk"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var (
	GL_DB         *gorm.DB
	GL_ZK         *zk.Conn
	GL_LOG        *zap.SugaredLogger
	GL_DING       *dingtalk.DingTalk
	GL_MYSQL_HOST string
	GL_MYSQL_DB   string
	GL_MYSQL_USER string
	GL_MYSQL_PASS string
	ZKADDR        string
	CHANNEL       []string
	DINGTOKEN     string
	DINGSECRTE    string
	REFRESHTIME   time.Duration
)
