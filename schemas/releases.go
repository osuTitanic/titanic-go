package schemas

import (
	"encoding/json"
	"time"
)

type Release struct {
	Name        string          `gorm:"column:name;primaryKey"`
	Version     int             `gorm:"column:version"`
	Description string          `gorm:"column:description;default:"`
	Category    string          `gorm:"column:category;default:Uncategorized"`
	KnownBugs   *string         `gorm:"column:known_bugs"`
	Supported   bool            `gorm:"column:supported;default:true"`
	Preview     bool            `gorm:"column:preview;default:false"`
	Downloads   []string        `gorm:"column:downloads;type:text[];default:'{}'"`
	Screenshots []string        `gorm:"column:screenshots;type:text[];default:'{}'"`
	Hashes      json.RawMessage `gorm:"column:hashes;type:jsonb;default:'[]'"`
	CreatedAt   time.Time       `gorm:"column:created_at;autoCreateTime"`
}

func (Release) TableName() string {
	return "releases_titanic"
}

type ModdedRelease struct {
	Name            string    `gorm:"column:name;primaryKey"`
	Description     string    `gorm:"column:description"`
	CreatorId       int       `gorm:"column:creator_id"`
	TopicId         int       `gorm:"column:topic_id"`
	ClientVersion   int       `gorm:"column:client_version"`
	ClientExtension string    `gorm:"column:client_extension"`
	CreatedAt       time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (ModdedRelease) TableName() string {
	return "releases_modding"
}

type ModdedReleaseEntries struct {
	Id          int       `gorm:"column:id;primaryKey;autoIncrement"`
	ModName     string    `gorm:"column:mod_name"`
	Version     string    `gorm:"column:version"`
	Stream      string    `gorm:"column:stream"`
	Checksum    string    `gorm:"column:checksum"`
	DownloadUrl *string   `gorm:"column:download_url"`
	UpdateUrl   *string   `gorm:"column:update_url"`
	PostId      *int      `gorm:"column:post_id"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (ModdedReleaseEntries) TableName() string {
	return "releases_modding_entries"
}

type ModdedReleaseChangelog struct {
	Id        int       `gorm:"column:id;primaryKey;autoIncrement"`
	EntryId   int       `gorm:"column:entry_id"`
	Text      string    `gorm:"column:text"`
	Type      string    `gorm:"column:type"`
	Branch    string    `gorm:"column:branch"`
	Author    string    `gorm:"column:author"`
	AuthorId  *int      `gorm:"column:author_id"`
	Area      *string   `gorm:"column:area"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (ModdedReleaseChangelog) TableName() string {
	return "releases_modding_changelog"
}

type ExtraRelease struct {
	Name        string `gorm:"column:name;primaryKey"`
	Description string `gorm:"column:description"`
	Download    string `gorm:"column:download"`
	Filename    string `gorm:"column:filename"`
	Md5         string `gorm:"column:md5"`
}

func (ExtraRelease) TableName() string {
	return "releases_extra"
}

type ReleasesOfficial struct {
	Id         int       `gorm:"column:id;primaryKey;autoIncrement"`
	Version    int       `gorm:"column:version"`
	Stream     string    `gorm:"column:stream"`
	Subversion int       `gorm:"column:subversion"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (ReleasesOfficial) TableName() string {
	return "releases_official"
}

type ReleasesOfficialEntries struct {
	ReleaseId int `gorm:"column:release_id;primaryKey"`
	FileId    int `gorm:"column:file_id;primaryKey"`
}

func (ReleasesOfficialEntries) TableName() string {
	return "releases_official_entries"
}

type ReleaseFiles struct {
	Id          int       `gorm:"column:id;primaryKey;autoIncrement"`
	Filename    string    `gorm:"column:filename"`
	FileVersion int       `gorm:"column:file_version"`
	FileHash    string    `gorm:"column:file_hash"`
	Filesize    int       `gorm:"column:filesize"`
	PatchId     *string   `gorm:"column:patch_id"`
	UrlFull     string    `gorm:"column:url_full"`
	UrlPatch    *string   `gorm:"column:url_patch"`
	Timestamp   time.Time `gorm:"column:timestamp;autoCreateTime"`
}

func (ReleaseFiles) TableName() string {
	return "releases_official_files"
}

type ReleaseChangelog struct {
	Id        int       `gorm:"column:id;primaryKey;autoIncrement"`
	Text      string    `gorm:"column:text"`
	Type      string    `gorm:"column:type"`
	Branch    string    `gorm:"column:branch"`
	Author    string    `gorm:"column:author"`
	Area      *string   `gorm:"column:area"`
	CreatedAt time.Time `gorm:"column:created_at;default:func.current_date()"`
}

func (ReleaseChangelog) TableName() string {
	return "releases_official_changelog"
}
