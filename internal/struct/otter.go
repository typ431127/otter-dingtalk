package _struct

import "gorm.io/gorm"

type Pipeline struct {
	ID              int
	NAME            string
	PARAMETERS      string
	DestinationName string
	gorm.Model      `gorm:"table:pipeline"`
}

type Binlog struct {
	Type     string   `json:"@type"`
	Identity Identity `json:"identity"`
	Postion  Postion  `json:"postion"`
}
type SourceAddress struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}
type Identity struct {
	SlaveID       int           `json:"slaveId"`
	SourceAddress SourceAddress `json:"sourceAddress"`
}
type Postion struct {
	Gtid        string `json:"gtid"`
	Included    bool   `json:"included"`
	JournalName string `json:"journalName"`
	Position    int    `json:"position"`
	ServerID    int    `json:"serverId"`
	Timestamp   int64  `json:"timestamp"`
	TimeString  string
}
