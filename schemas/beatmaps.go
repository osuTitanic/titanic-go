package schemas

import (
	"time"

	"github.com/osuTitanic/common-go/constants"
)

type Beatmapset struct {
	Id                 int                     `gorm:"column:id;primaryKey;autoIncrement"`
	Title              *string                 `gorm:"column:title"`
	TitleUnicode       *string                 `gorm:"column:title_unicode"`
	Artist             *string                 `gorm:"column:artist"`
	ArtistUnicode      *string                 `gorm:"column:artist_unicode"`
	Source             *string                 `gorm:"column:source"`
	SourceUnicode      *string                 `gorm:"column:source_unicode"`
	Creator            *string                 `gorm:"column:creator"`
	DisplayTitle       *string                 `gorm:"column:display_title"`
	Description        *string                 `gorm:"column:description"`
	Tags               *string                 `gorm:"column:tags;default:"`
	Status             int                     `gorm:"column:submission_status;default:3"`
	HasVideo           bool                    `gorm:"column:has_video;default:false"`
	HasStoryboard      bool                    `gorm:"column:has_storyboard;default:false"`
	Server             constants.BeatmapServer `gorm:"column:server;default:0"`
	DownloadServer     constants.BeatmapServer `gorm:"column:download_server;default:0"`
	TopicId            *int                    `gorm:"column:topic_id"`
	CreatorId          *int                    `gorm:"column:creator_id"`
	Available          bool                    `gorm:"column:available;default:true"`
	Enhanced           bool                    `gorm:"column:enhanced;default:false"`
	Explicit           bool                    `gorm:"column:explicit;default:false"`
	CreatedAt          time.Time               `gorm:"column:submission_date;autoCreateTime"`
	ApprovedAt         *time.Time              `gorm:"column:approved_date"`
	ApprovedBy         *int                    `gorm:"column:approved_by"`
	LastUpdate         time.Time               `gorm:"column:last_updated;autoCreateTime"`
	AddedAt            *time.Time              `gorm:"column:added_at;autoCreateTime"`
	TotalPlaycount     int64                   `gorm:"column:total_playcount;default:0;->"`
	MaxDiff            float64                 `gorm:"column:max_diff;default:0.0;->"`
	RatingAverage      float64                 `gorm:"column:rating_average;default:0.0;->"`
	RatingCount        int                     `gorm:"column:rating_count;default:0;->"`
	FavouriteCount     int                     `gorm:"column:favourite_count;default:0;->"`
	OszFilesize        int                     `gorm:"column:osz_filesize;default:0"`
	OszFilesizeNovideo int                     `gorm:"column:osz_filesize_novideo;default:0"`
	LanguageId         int                     `gorm:"column:language_id;default:1"`
	GenreId            int                     `gorm:"column:genre_id;default:1"`
	StarPriority       int                     `gorm:"column:star_priority;default:0"`
	Offset             int                     `gorm:"column:offset;default:0"`
	MetaHash           *string                 `gorm:"column:meta_hash"`
	InfoHash           *string                 `gorm:"column:info_hash"`
	BodyHash           *string                 `gorm:"column:body_hash"`
	Search             string                  `gorm:"column:search;type:tsvector;->"`
}

func (Beatmapset) TableName() string {
	return "beatmapsets"
}

type Beatmap struct {
	Id               int            `gorm:"column:id;primaryKey;autoIncrement"`
	SetId            int            `gorm:"column:set_id"`
	Mode             constants.Mode `gorm:"column:mode;default:0"`
	Checksum         string         `gorm:"column:md5"`
	Status           int            `gorm:"column:status;default:2"` // TODO: Use status constant
	Version          string         `gorm:"column:version"`
	Filename         string         `gorm:"column:filename"`
	CreatedAt        time.Time      `gorm:"column:submission_date;autoCreateTime"`
	LastUpdate       time.Time      `gorm:"column:last_updated;autoCreateTime"`
	Playcount        int64          `gorm:"column:playcount;default:0"`
	Passcount        int64          `gorm:"column:passcount;default:0"`
	TotalLength      int            `gorm:"column:total_length"`
	DrainLength      int            `gorm:"column:drain_length;default:0"`
	CountNormal      int            `gorm:"column:count_normal;default:0"`
	CountSlider      int            `gorm:"column:count_slider;default:0"`
	CountSpinner     int            `gorm:"column:count_spinner;default:0"`
	MaxCombo         int            `gorm:"column:max_combo"`
	BPM              float64        `gorm:"column:bpm;default:0.0"`
	CS               float64        `gorm:"column:cs;default:0.0"`
	AR               float64        `gorm:"column:ar;default:0.0"`
	OD               float64        `gorm:"column:od;default:0.0"`
	HP               float64        `gorm:"column:hp;default:0.0"`
	Diff             float64        `gorm:"column:diff;default:0.0"`
	DiffEyup         float64        `gorm:"column:diff_eyup;default:0.0"`
	SliderMultiplier float64        `gorm:"column:slider_multiplier;default:0.0"`
	Search           string         `gorm:"column:search;type:tsvector;->"`
}

func (Beatmap) TableName() string {
	return "beatmaps"
}

type BeatmapCollaboration struct {
	UserId               int       `gorm:"column:user_id;primaryKey"`
	BeatmapId            int       `gorm:"column:beatmap_id;primaryKey"`
	IsBeatmapAuthor      bool      `gorm:"column:is_beatmap_author;default:false"`
	AllowResourceUpdates bool      `gorm:"column:allow_resource_updates;default:false"`
	CreatedAt            time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (BeatmapCollaboration) TableName() string {
	return "beatmap_collaboration"
}

type BeatmapCollaborationRequest struct {
	Id                   int       `gorm:"column:id;primaryKey;autoIncrement"`
	UserId               int       `gorm:"column:user_id"`
	TargetId             int       `gorm:"column:target_id"`
	BeatmapId            int       `gorm:"column:beatmap_id"`
	AllowResourceUpdates bool      `gorm:"column:allow_resource_updates;default:false"`
	CreatedAt            time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (BeatmapCollaborationRequest) TableName() string {
	return "beatmap_collaboration_requests"
}

type BeatmapCollaborationBlacklist struct {
	UserId    int       `gorm:"column:user_id;primaryKey"`
	TargetId  int       `gorm:"column:target_id;primaryKey"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (BeatmapCollaborationBlacklist) TableName() string {
	return "beatmap_collaboration_blacklist"
}

type BeatmapNomination struct {
	UserId int       `gorm:"column:user_id;primaryKey"`
	SetId  int       `gorm:"column:set_id;primaryKey"`
	Time   time.Time `gorm:"column:time;autoCreateTime"`
}

func (BeatmapNomination) TableName() string {
	return "beatmap_nominations"
}

type BeatmapModding struct {
	Id       int       `gorm:"column:id;primaryKey;autoIncrement"`
	TargetId int       `gorm:"column:target_id"`
	SenderId int       `gorm:"column:sender_id"`
	SetId    int       `gorm:"column:set_id"`
	PostId   int       `gorm:"column:post_id"`
	Amount   int       `gorm:"column:amount;default:0"`
	Time     time.Time `gorm:"column:time;autoCreateTime"`
}

func (BeatmapModding) TableName() string {
	return "beatmap_modding"
}

type BeatmapPack struct {
	Id           int       `gorm:"column:id;primaryKey;autoIncrement"`
	Name         string    `gorm:"column:name"`
	Category     string    `gorm:"column:category"`
	DownloadLink string    `gorm:"column:download_link"`
	Description  string    `gorm:"column:description;default:"`
	CreatorId    int       `gorm:"column:creator_id"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoCreateTime"`
}

func (BeatmapPack) TableName() string {
	return "beatmap_packs"
}

type BeatmapPackEntry struct {
	PackId       int       `gorm:"column:pack_id;primaryKey"`
	BeatmapsetId int       `gorm:"column:beatmapset_id;primaryKey"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (BeatmapPackEntry) TableName() string {
	return "beatmap_pack_entries"
}

type BeatmapPlays struct {
	UserId      int    `gorm:"column:user_id;primaryKey"`
	BeatmapId   int    `gorm:"column:beatmap_id;primaryKey"`
	SetId       int    `gorm:"column:set_id"`
	Count       int    `gorm:"column:count"`
	BeatmapFile string `gorm:"column:beatmap_file"`
}

func (BeatmapPlays) TableName() string {
	return "plays"
}

type BeatmapFavourite struct {
	UserId    int       `gorm:"column:user_id;primaryKey"`
	SetId     int       `gorm:"column:set_id;primaryKey"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (BeatmapFavourite) TableName() string {
	return "favourites"
}

type BeatmapRating struct {
	UserId      int    `gorm:"column:user_id;primaryKey"`
	SetId       int    `gorm:"column:set_id"`
	MapChecksum string `gorm:"column:map_checksum;primaryKey"`
	Rating      int    `gorm:"column:rating"`
}

func (BeatmapRating) TableName() string {
	return "ratings"
}

type BeatmapComment struct {
	Id         int            `gorm:"column:id;primaryKey;autoIncrement"`
	TargetId   int            `gorm:"column:target_id"`
	TargetType string         `gorm:"column:target_type"`
	UserId     int            `gorm:"column:user_id"`
	Mode       constants.Mode `gorm:"column:mode;default:0"`
	Time       time.Time      `gorm:"column:time;autoCreateTime"`
	Comment    string         `gorm:"column:comment"`
	Format     *string        `gorm:"column:format"`
	Color      *string        `gorm:"column:color"`
}

func (BeatmapComment) TableName() string {
	return "comments"
}

type BeatmapMirror struct {
	Url      string                  `gorm:"column:url;primaryKey"`
	Server   constants.BeatmapServer `gorm:"column:server"`
	Type     int                     `gorm:"column:type"`
	Priority int                     `gorm:"column:priority;default:0"`
}

func (BeatmapMirror) TableName() string {
	return "resource_mirrors"
}
