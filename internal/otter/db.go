package otter

import (
	"github.com/tidwall/gjson"
	"gorm.io/gorm"
	"otter-dingtalk/internal/global"
	_struct "otter-dingtalk/internal/struct"
	"regexp"
	"time"
)

type Channel struct {
	ID         int
	NAME       string
	gorm.Model `gorm:"table:channel"`
}
type Logrecord struct {
	CHANNEL_ID  int
	PIPELINE_ID int
	TITLE       string
	MESSAGE     string
	GMT_CREATE  string
	gorm.Model  `gorm:"table:log_record"`
}
type DelayStat struct {
	ID           int
	DELAY_TIME   int32
	PIPELINE_ID  int
	GMT_CREATE   time.Time
	GMT_MODIFIED time.Time
	gorm.Model   `gorm:"table:log_record"`
}

func getChannelName(channelId int) Channel {
	var channelName = Channel{ID: channelId}
	global.GL_DB.Raw("SELECT ID,NAME FROM `channel` WHERE ID = ?", channelId).Scan(&channelName)
	return channelName
}
func getPipeline(pipelineid int) _struct.Pipeline {
	var pipeline = _struct.Pipeline{ID: pipelineid}
	global.GL_DB.Raw("SELECT * FROM `pipeline` WHERE `CHANNEL_ID` = ? LIMIT 1", pipelineid).Scan(&pipeline)
	pipeline.DestinationName = gjson.Get(pipeline.PARAMETERS, "destinationName").String()
	return pipeline
}
func getDelaystat(channelId int) DelayStat {
	var result = DelayStat{PIPELINE_ID: channelId}
	global.GL_DB.Raw("SELECT * FROM `delay_stat` WHERE PIPELINE_ID = ? ORDER BY `delay_stat`.`ID` DESC LIMIT 1", channelId).Scan(&result)

	return result

}
func getChannelLog(channelId int) string {
	var log = ""
	//var re = regexp.MustCompile(`(?m)PreparedStatementCallback.*`)
	var logrecord = Logrecord{CHANNEL_ID: channelId}
	global.GL_DB.Raw("SELECT CHANNEL_ID,PIPELINE_ID,TITLE,MESSAGE,GMT_CREATE FROM `log_record` WHERE `CHANNEL_ID` = ? ORDER BY GMT_CREATE desc LIMIT 1;", channelId).Scan(&logrecord)
	//for _, match := range re.FindAllString(logrecord.MESSAGE, -1) {
	//	log += match
	//}
	for _, re := range []*regexp.Regexp{regexp.MustCompile(`(?m)PreparedStatementCallback.*`), regexp.MustCompile(`(?m)MySQLSyntaxErrorException.*`), regexp.MustCompile(`(?m)TransformException.*`)} {
		for _, match := range re.FindAllString(logrecord.MESSAGE, -1) {
			log += match
		}
	}
	return log
}
