package alert

import (
	"fmt"
	"github.com/blinkbean/dingtalk"
	"otter-dingtalk/internal/global"
	"time"
)

func DingAlert(channelname string, state string) {
	dm := dingtalk.DingMap()
	dm.Set(fmt.Sprintf("通道:%s状态异常", channelname), dingtalk.H2)
	dm.Set(fmt.Sprintf("状态:%s", state), dingtalk.RED)
	dm.Set(fmt.Sprintf("时间: %s", time.Now().Format("2006-01-02 15:04:05")), dingtalk.N)
	global.GL_DING.SendMarkDownMessageBySlice("otter同步异常", dm.Slice())
}
