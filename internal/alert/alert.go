package alert

import (
	"fmt"
	"github.com/blinkbean/dingtalk"
	"otter-dingtalk/internal/global"
	_struct "otter-dingtalk/internal/struct"
	"time"
)

type MarkType string

const (
	ALERT MarkType = "alert"
)

type PostAlert struct {
	Module    string   `json:"module"  binding:"required"`
	Content   string   `json:"content" binding:"required"`
	Labels    []string `json:"labels" binding:"required"`
	Level     string   `json:"level" binding:"required"`
	Server    string   `json:"server" binding:"required"`
	Hostname  string   `json:"hostname" binding:"required"`
	Env       string   `json:"env" binding:"required"`
	Timestamp int32    `json:"@timestamp"`
}

func DingAlert(channelname, state, delaytime, synctime string, binlog _struct.Binlog, pipeline _struct.Pipeline, log string) {
	dm := dingtalk.DingMap()
	dm.Set(fmt.Sprintf("通道: %s状态异常", channelname), dingtalk.H2)
	dm.Set(fmt.Sprintf("状态: %s", state), dingtalk.RED)
	dm.Set(fmt.Sprintf("环境: %s", global.ENV), dingtalk.GREEN)
	dm.Set(fmt.Sprintf("延迟: %ss", delaytime), dingtalk.N)
	dm.Set(fmt.Sprintf("源库: %s:%d", binlog.Identity.SourceAddress.Address, binlog.Identity.SourceAddress.Port), dingtalk.N)
	dm.Set(fmt.Sprintf("位点: %s | %d", binlog.Postion.JournalName, binlog.Postion.Position), dingtalk.N)
	dm.Set(fmt.Sprintf("位点时间: %s", binlog.Postion.TimeString), dingtalk.N)
	dm.Set(fmt.Sprintf("最后同步: %s", synctime), dingtalk.N)
	dm.Set(fmt.Sprintf("报警时间: %s", time.Now().Format("2006-01-02 15:04:05")), dingtalk.N)
	dm.Set(fmt.Sprintf("错误日志:\n```\n%s\n```\n", log), dingtalk.N)
	dm.Set(fmt.Sprintf("[日志传送门](%s/log_record_tab.htm?pipelineId=%d)", global.OTTER_URL, pipeline.ID), dingtalk.N)
	dm.Set(fmt.Sprintf("自动解挂任务已开启"), dingtalk.BLUE)
	global.GL_DING.SendMarkDownMessageBySlice("otter同步异常", dm.Slice())
}
func DingAlertDelay(channelname, state, delaytime, synctime string, binlog _struct.Binlog, log string) {
	dm := dingtalk.DingMap()
	dm.Set(fmt.Sprintf("通道: %s延迟过大", channelname), dingtalk.H2)
	dm.Set(fmt.Sprintf("状态: %s", state), dingtalk.RED)
	dm.Set(fmt.Sprintf("时间: %s", time.Now().Format("2006-01-02 15:04:05")), dingtalk.N)
	dm.Set(fmt.Sprintf("环境: %s", global.ENV), dingtalk.GREEN)
	dm.Set(fmt.Sprintf("延迟: %ss", delaytime), dingtalk.N)
	dm.Set(fmt.Sprintf("源库: %s:%d", binlog.Identity.SourceAddress.Address, binlog.Identity.SourceAddress.Port), dingtalk.N)
	dm.Set(fmt.Sprintf("位点: %s | %d", binlog.Postion.JournalName, binlog.Postion.Position), dingtalk.N)
	dm.Set(fmt.Sprintf("位点时间: %s", binlog.Postion.TimeString), dingtalk.N)
	dm.Set(fmt.Sprintf("最后同步: %s", synctime), dingtalk.N)
	global.GL_DING.SendMarkDownMessageBySlice("otter延迟过大", dm.Slice())
}

func DingAlertResolve(channelname, delaytime, synctime, state string, binlog _struct.Binlog, atype int) {
	dm := dingtalk.DingMap()
	if atype == 1 {
		dm.Set(fmt.Sprintf("通道: %s 延迟过大恢复正常", channelname), dingtalk.H2)
	} else {
		dm.Set(fmt.Sprintf("通道: %s恢复正常", channelname), dingtalk.H2)
	}
	dm.Set(fmt.Sprintf("环境: %s", global.ENV), dingtalk.GREEN)
	dm.Set(fmt.Sprintf("状态: %s", state), dingtalk.GREEN)
	dm.Set(fmt.Sprintf("延迟: %ss", delaytime), dingtalk.N)
	dm.Set(fmt.Sprintf("源库: %s:%d", binlog.Identity.SourceAddress.Address, binlog.Identity.SourceAddress.Port), dingtalk.N)
	dm.Set(fmt.Sprintf("位点: %s | %d", binlog.Postion.JournalName, binlog.Postion.Position), dingtalk.N)
	dm.Set(fmt.Sprintf("位点时间: %s", binlog.Postion.TimeString), dingtalk.N)
	dm.Set(fmt.Sprintf("最后同步: %s", synctime), dingtalk.N)
	dm.Set(fmt.Sprintf("恢复时间: %s", time.Now().Format("2006-01-02 15:04:05")), dingtalk.N)
	global.GL_DING.SendMarkDownMessageBySlice("otter恢复正常", dm.Slice())
}
