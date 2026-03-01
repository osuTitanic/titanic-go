package database

import "github.com/osuTitanic/common-go/schemas"

type IUserRepository interface {
	Create(user *schemas.User) error
	Update(id int, updates map[string]interface{}) (int64, error)
	Delete(user *schemas.User) error

	ById(id int, preload ...string) (*schemas.User, error)
	ByName(name string, preload ...string) (*schemas.User, error)
	BySafeName(safeName string, preload ...string) (*schemas.User, error)
	ByEmail(email string, preload ...string) (*schemas.User, error)
	ByDiscordId(discordId int64, preload ...string) (*schemas.User, error)

	ManyById(userIds []int, preload ...string) ([]*schemas.User, error)
	ManyByName(names []string, preload ...string) ([]*schemas.User, error)
	ManyByRank(limit int, ascending bool, preload ...string) ([]*schemas.User, error)
	ManyByCreationDate(limit int, ascending bool, preload ...string) ([]*schemas.User, error)

	GetUsername(id int) (string, error)
	GetUserId(name string) (int, error)
	GetAvatarChecksum(id int) (string, error)
	GetCount() (int, error)
}

type IStatsRepository interface {
	Create(stats *schemas.Stats) error
	Update(userId int, mode int, updates map[string]interface{}) (int64, error)
	Delete(stats *schemas.Stats) error

	ByMode(userId int, mode int, preload ...string) (*schemas.Stats, error)
	ManyByUserId(userId int, preload ...string) ([]*schemas.Stats, error)
}

type IRelationshipRepository interface {
	Create(relationship *schemas.Relationship) error
	Update(userId int, targetId int, updates map[string]interface{}) (int64, error)
	Delete(relationship *schemas.Relationship) error

	ByUserAndTarget(userId int, targetId int, preload ...string) (*schemas.Relationship, error)
	ManyByUserId(userId int, preload ...string) ([]*schemas.Relationship, error)
	ManyByTargetId(targetId int, preload ...string) ([]*schemas.Relationship, error)
	CountByUserId(userId int) (int, error)
	CountByTargetId(targetId int) (int, error)
}

type IBadgeRepository interface {
	Create(badge *schemas.Badge) error
	Update(id int, updates map[string]interface{}) (int64, error)
	Delete(badge *schemas.Badge) error

	ById(id int, preload ...string) (*schemas.Badge, error)
	ManyByUserId(userId int, preload ...string) ([]*schemas.Badge, error)
}

type INameRepository interface {
	Create(name *schemas.Name) error
	Update(id int, updates map[string]interface{}) (int64, error)
	Delete(name *schemas.Name) error

	ById(id int, preload ...string) (*schemas.Name, error)
	ByName(name string, preload ...string) (*schemas.Name, error)
	ManyByUserId(userId int, preload ...string) ([]*schemas.Name, error)
}

type IInfringementRepository interface {
	Create(infringement *schemas.Infringement) error
	Update(id int, updates map[string]interface{}) (int64, error)
	Delete(infringement *schemas.Infringement) error

	ById(id int, preload ...string) (*schemas.Infringement, error)
	ManyByUserId(userId int, preload ...string) ([]*schemas.Infringement, error)
}

type IReportRepository interface {
	Create(report *schemas.Report) error
	Update(id int, updates map[string]interface{}) (int64, error)
	Delete(report *schemas.Report) error

	ById(id int, preload ...string) (*schemas.Report, error)
	ManyByTargetId(targetId int, preload ...string) ([]*schemas.Report, error)
	ManyBySenderId(senderId int, preload ...string) ([]*schemas.Report, error)
}

type IVerificationRepository interface {
	Create(verification *schemas.Verification) error
	Update(id int, updates map[string]interface{}) (int64, error)
	Delete(verification *schemas.Verification) error

	ById(id int, preload ...string) (*schemas.Verification, error)
	ByToken(token string, preload ...string) (*schemas.Verification, error)
	ManyByUserId(userId int, preload ...string) ([]*schemas.Verification, error)
	DeleteByToken(token string) error
}

type IGroupRepository interface {
	Create(group *schemas.Group) error
	Update(id int, updates map[string]interface{}) (int64, error)
	Delete(group *schemas.Group) error

	ById(id int, preload ...string) (*schemas.Group, error)
	Many(includeHidden bool, preload ...string) ([]*schemas.Group, error)
}

type IGroupEntryRepository interface {
	Create(entry *schemas.GroupEntry) error
	Update(userId int, groupId int, updates map[string]interface{}) (int64, error)
	Delete(entry *schemas.GroupEntry) error

	ByUserAndGroup(userId int, groupId int, preload ...string) (*schemas.GroupEntry, error)
	ManyByUserId(userId int, preload ...string) ([]*schemas.GroupEntry, error)
	ManyByGroupId(groupId int, preload ...string) ([]*schemas.GroupEntry, error)
}

type IUserPermissionRepository interface {
	Create(permission *schemas.UserPermission) error
	Update(id int, updates map[string]interface{}) (int64, error)
	Delete(permission *schemas.UserPermission) error

	ById(id int, preload ...string) (*schemas.UserPermission, error)
	ManyByUserId(userId int, preload ...string) ([]*schemas.UserPermission, error)
}

type IGroupPermissionRepository interface {
	Create(permission *schemas.GroupPermission) error
	Update(id int, updates map[string]interface{}) (int64, error)
	Delete(permission *schemas.GroupPermission) error

	ById(id int, preload ...string) (*schemas.GroupPermission, error)
	ManyByGroupId(groupId int, preload ...string) ([]*schemas.GroupPermission, error)
}

type INotificationRepository interface {
	Create(notification *schemas.Notification) error
	Update(id int64, updates map[string]interface{}) (int64, error)
	Delete(notification *schemas.Notification) error

	ById(id int64, preload ...string) (*schemas.Notification, error)
	ManyByUserId(userId int, preload ...string) ([]*schemas.Notification, error)
	CountByUserId(userId int) (int, error)
}

type IAchievementRepository interface {
	Create(achievement *schemas.Achievement) error
	Update(userId int, name string, updates map[string]interface{}) (int64, error)
	Delete(achievement *schemas.Achievement) error

	ManyByUserId(userId int, preload ...string) ([]*schemas.Achievement, error)
}

type IBeatmapFavouriteRepository interface {
	Create(favourite *schemas.BeatmapFavourite) error
	Update(userId int, setId int, updates map[string]interface{}) (int64, error)
	Delete(favourite *schemas.BeatmapFavourite) error

	ByUserAndSet(userId int, setId int, preload ...string) (*schemas.BeatmapFavourite, error)
	ManyByUserId(userId int, preload ...string) ([]*schemas.BeatmapFavourite, error)
	CountByUserId(userId int) (int, error)
	CountBySetId(setId int) (int, error)
}
