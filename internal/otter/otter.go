package otter

import (
	"gorm.io/gorm"
	"otter-dingtalk/internal/alert"
	"otter-dingtalk/internal/global"
	"otter-dingtalk/pkg/ext"
	"strconv"
	"strings"
	"time"
)

type Channel struct {
	ID         int
	NAME       string
	gorm.Model `gorm:"table:channel"`
}

func getChannelName(channelid int) Channel {
	var channelName = Channel{ID: channelid}
	global.GL_DB.Raw("SELECT ID,NAME FROM `channel` WHERE ID = ?", channelid).Scan(&channelName)
	return channelName
}
func getAllChannel() []string {
	list, _, err := global.GL_ZK.Children("/otter/channel")
	if err != nil {
		global.GL_LOG.Error(err)
	}
	return list
}
func CheckAlarms() {
	var arrayChannel []string
	for {
		if ext.InString("all", global.CHANNEL) {
			arrayChannel = getAllChannel()
		} else {
			arrayChannel = global.CHANNEL
		}
		for _, channelid := range arrayChannel {
			bytestate, _, err := global.GL_ZK.Get("/otter/channel/" + channelid)
			if err != nil {
				global.GL_LOG.Error(err)
			}
			state := strings.Replace(string(bytestate), "\"", "", 2)
			if state != "START" && len(state) != 0 {
				id, _ := strconv.Atoi(channelid)
				channel := getChannelName(id)
				if len(channel.NAME) != 0 {
					global.GL_LOG.Infof("ChannelId:%s State:%s Name:%s", channelid, state, channel.NAME)
					alert.DingAlert(channel.NAME, state)
				}
			} else {
				global.GL_LOG.Infof("ChannelId:%s State:%s", channelid, state)
			}

		}
		time.Sleep(global.REFRESHTIME)
	}
}
