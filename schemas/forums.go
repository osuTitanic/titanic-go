package schemas

import "time"

type Forum struct {
	Id          int       `gorm:"column:id;primaryKey;autoIncrement"`
	ParentId    *int      `gorm:"column:parent_id"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description;default:"`
	TopicCount  int       `gorm:"column:topic_count;default:0;->"`
	PostCount   int       `gorm:"column:post_count;default:0;->"`
	AllowIcons  bool      `gorm:"column:allow_icons;default:true"`
	Hidden      bool      `gorm:"column:hidden;default:false"`
}

func (Forum) TableName() string {
	return "forums"
}

type ForumTopic struct {
	Id            int        `gorm:"column:id;primaryKey;autoIncrement"`
	ForumId       int        `gorm:"column:forum_id"`
	CreatorId     int        `gorm:"column:creator_id"`
	Title         string     `gorm:"column:title"`
	StatusText    *string    `gorm:"column:status_text"`
	CreatedAt     time.Time  `gorm:"column:created_at;autoCreateTime"`
	LastPostAt    time.Time  `gorm:"column:last_post_at;autoCreateTime"`
	LockedAt      *time.Time `gorm:"column:locked_at"`
	Views         int        `gorm:"column:views;default:0"`
	PostCount     int        `gorm:"column:post_count;default:0;->"`
	IconId        *int       `gorm:"column:icon"`
	CanChangeIcon bool       `gorm:"column:can_change_icon;default:true"`
	CanStar       bool       `gorm:"column:can_star;default:false"`
	Announcement  bool       `gorm:"column:announcement;default:false"`
	Hidden        bool       `gorm:"column:hidden;default:false"`
	Pinned        bool       `gorm:"column:pinned;default:false"`
}

func (ForumTopic) TableName() string {
	return "forum_topics"
}

type ForumPost struct {
	Id         int64     `gorm:"column:id;primaryKey;autoIncrement"`
	TopicId    int       `gorm:"column:topic_id"`
	ForumId    int       `gorm:"column:forum_id"`
	UserId     int       `gorm:"column:user_id"`
	IconId     *int      `gorm:"column:icon_id"`
	Content    string    `gorm:"column:content"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	EditTime   time.Time `gorm:"column:edit_time;autoCreateTime"`
	EditCount  int       `gorm:"column:edit_count;default:0"`
	EditLocked bool      `gorm:"column:edit_locked;default:false"`
	Hidden     bool      `gorm:"column:hidden;default:false"`
	Draft      bool      `gorm:"column:draft;default:false"`
	Deleted    bool      `gorm:"column:deleted;default:false"`
}

func (ForumPost) TableName() string {
	return "forum_posts"
}

type ForumIcon struct {
	Id       int    `gorm:"column:id;primaryKey;autoIncrement"`
	Name     string `gorm:"column:name"`
	Location string `gorm:"column:location"`
	Order    int    `gorm:"column:order;default:0"`
}

func (ForumIcon) TableName() string {
	return "forum_icons"
}

type ForumReport struct {
	PostId     int        `gorm:"column:post_id;primaryKey"`
	UserId     int        `gorm:"column:user_id;primaryKey"`
	CreatedAt  time.Time  `gorm:"column:created_at;autoCreateTime"`
	ResolvedAt *time.Time `gorm:"column:resolved_at"`
	Reason     string     `gorm:"column:reason"`
}

func (ForumReport) TableName() string {
	return "forum_reports"
}

type ForumStar struct {
	TopicId   int       `gorm:"column:topic_id;primaryKey"`
	UserId    int       `gorm:"column:user_id;primaryKey"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (ForumStar) TableName() string {
	return "forum_stars"
}

type ForumBookmark struct {
	UserId  int `gorm:"column:user_id;primaryKey"`
	TopicId int `gorm:"column:topic_id;primaryKey"`
}

func (ForumBookmark) TableName() string {
	return "forum_bookmarks"
}

type ForumSubscriber struct {
	UserId  int `gorm:"column:user_id;primaryKey"`
	TopicId int `gorm:"column:topic_id;primaryKey"`
}

func (ForumSubscriber) TableName() string {
	return "forum_subscribers"
}
