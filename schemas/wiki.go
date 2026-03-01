package schemas

import (
	"encoding/json"
	"time"
)

type WikiPage struct {
	Id          int       `gorm:"column:id;primaryKey;autoIncrement"`
	Name        string    `gorm:"column:name"`
	Path        string    `gorm:"column:path"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	LastUpdated time.Time `gorm:"column:last_updated;autoCreateTime"`
	CategoryId  int       `gorm:"column:category_id"`
}

func (WikiPage) TableName() string {
	return "wiki_pages"
}

type WikiCategory struct {
	Id           int             `gorm:"column:id;primaryKey;autoIncrement"`
	Name         string          `gorm:"column:name"`
	Translations json.RawMessage `gorm:"column:translations;type:jsonb;default:'{}'"`
	ParentId     *int            `gorm:"column:parent_id"`
	CreatedAt    time.Time       `gorm:"column:created_at;autoCreateTime"`
}

func (WikiCategory) TableName() string {
	return "wiki_categories"
}

type WikiContent struct {
	PageId      int       `gorm:"column:page_id;primaryKey"`
	Language    string    `gorm:"column:language;primaryKey"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	LastUpdated time.Time `gorm:"column:last_updated;autoCreateTime"`
	Title       string    `gorm:"column:title"`
	Content     string    `gorm:"column:content"`
	Search      string    `gorm:"column:search;type:tsvector;->"`
}

func (WikiContent) TableName() string {
	return "wiki_content"
}

type WikiOutlink struct {
	PageId    int       `gorm:"column:page_id;primaryKey"`
	TargetId  int       `gorm:"column:target_id;primaryKey"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (WikiOutlink) TableName() string {
	return "wiki_outlinks"
}
