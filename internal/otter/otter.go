package otter

import (
	"encoding/json"
	"fmt"
	"otter-dingtalk/internal/alert"
	"otter-dingtalk/internal/global"
	_struct "otter-dingtalk/internal/struct"
	"otter-dingtalk/pkg/ext"
	"strconv"
	"strings"
	"time"
)

var alertData map[string]int

func getBinlog(pipelineinfo _struct.Pipeline) _struct.Binlog {
	var binlog _struct.Binlog
	path := fmt.Sprintf("/otter/canal/destinations/%s/%s/cursor", pipelineinfo.DestinationName, strconv.Itoa(pipelineinfo.ID))
	global.GL_LOG.Info(path)
	jsonData, _, err := global.GL_ZK.Get(path)
	if err != nil {
		global.GL_LOG.Error(err)
		return binlog
	}
	err = json.Unmarshal(jsonData, &binlog)
	if err != nil {
		global.GL_LOG.Error(err)
		return binlog
	}
	//millisecondsTimestamp := int64(binlog.Postion.Timestamp)
	seconds := binlog.Postion.Timestamp / 1000
	timestamp := time.Unix(seconds, 0)
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		global.GL_LOG.Errorf("加载上海时间失败!%d", binlog.Postion.Position)
	}
	beijingTime := timestamp.In(loc)
	binlog.Postion.TimeString = beijingTime.Format("2006-01-02 15:04:05")
	return binlog
}
func getAllChannel() []string {
	list, _, err := global.GL_ZK.Children("/otter/channel")
	if err != nil {
		global.GL_LOG.Error(err)
	}
	return list
}
func MonitoringPatrol() {
	InboxTk := time.NewTicker(global.REFRESHTIME)
	CheckAlarms()
	go func(T *time.Ticker) {
		for {
			<-T.C
			CheckAlarms()
		}
	}(InboxTk)
	defer InboxTk.Stop()
	select {}
}
func CheckAlarms() {
	var arrayChannel []string
	if ext.InString("all", global.CHANNEL) {
		arrayChannel = getAllChannel()
	} else {
		arrayChannel = global.CHANNEL
	}
	for _, channelid := range arrayChannel {
		bytestate, _, err := global.GL_ZK.Get("/otter/channel/" + channelid)
		if err != nil {
			global.GL_LOG.Error(err)
			continue
		}
		state := strings.Replace(string(bytestate), "\"", "", 2)
		channelid, _ := strconv.Atoi(channelid)
		channel := getChannelName(channelid)
		pipelinedata := getPipeline(channelid)
		binlogdata := getBinlog(pipelinedata)
		delaystat := getDelaystat(pipelinedata.ID)
		delaytime := fmt.Sprintf("%.2f", float32(delaystat.DELAY_TIME)/float32(1000))
		synctime := delaystat.GMT_CREATE.Format("2006-01-02 15:04:05")
		// 延迟过大检查
		if state == "START" && len(state) != 0 && state != "STOP" {
			if len(channel.NAME) != 0 && delaystat.DELAY_TIME > 60000 {
				log := getChannelLog(channelid)
				global.GL_LOG.Infof("延迟过大:%d", delaystat.DELAY_TIME)
				alert.DingAlertDelay(channel.NAME, state, delaytime, synctime, binlogdata, log)
				alertData[channel.NAME] = 1
			} else {
				if atype, ok := alertData[channel.NAME]; ok {
					delete(alertData, channel.NAME)
					alert.DingAlertResolve(channel.NAME, delaytime, synctime, state, binlogdata, atype)
				}
			}
		}
		// 挂起检查
		if state != "START" && len(state) != 0 && state != "STOP" {
			if len(channel.NAME) != 0 {
				log := getChannelLog(channelid)
				global.GL_LOG.Infof("ChannelId:%d State:%s Name:%s Log:%s", channelid, state, channel.NAME, log)
				alert.DingAlert(channel.NAME, state, delaytime, synctime, binlogdata, pipelinedata, log)
				alertData[channel.NAME+"_pause"]++
				go func() {
					for i := 0; i < 3; i++ {
						startchannel(strconv.Itoa(channelid))
						time.Sleep(5 * time.Second)
					}
				}()
			}
		} else {
			if atype, ok := alertData[channel.NAME+"_pause"]; ok {
				// 删除map数据发送恢复正常信息
				delete(alertData, channel.NAME+"_pause")
				alert.DingAlertResolve(channel.NAME, delaytime, synctime, state, binlogdata, atype)
			}
			global.GL_LOG.Infof("ChannelId:%d State:%s Delay:%ss", channelid, state, delaytime)
		}
	}
}

func init() {
	alertData = make(map[string]int)
}
