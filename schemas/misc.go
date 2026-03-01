package schemas

import (
	"encoding/json"
	"time"
)

type Screenshot struct {
	Id        int       `gorm:"column:id;primaryKey;autoIncrement"`
	UserId    string    `gorm:"column:user_id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	Hidden    bool      `gorm:"column:hidden;default:false"`
}

func (Screenshot) TableName() string {
	return "screenshots"
}

type Benchmark struct {
	Id         int             `gorm:"column:id;primaryKey;autoIncrement"`
	UserId     int             `gorm:"column:user_id"`
	Smoothness float64         `gorm:"column:smoothness"`
	Framerate  int             `gorm:"column:framerate"`
	Score      int64           `gorm:"column:score"`
	Grade      string          `gorm:"column:grade;default:N"`
	CreatedAt  time.Time       `gorm:"column:created_at;autoCreateTime"`
	Client     string          `gorm:"column:client"`
	Hardware   json.RawMessage `gorm:"column:hardware;type:jsonb"`
}

func (Benchmark) TableName() string {
	return "benchmarks"
}

type Log struct {
	Id      int       `gorm:"column:id;primaryKey;autoIncrement"`
	Level   string    `gorm:"column:level"`
	Type    string    `gorm:"column:type"`
	Message string    `gorm:"column:message"`
	Time    time.Time `gorm:"column:time;autoCreateTime"`
}

func (Log) TableName() string {
	return "logs"
}
