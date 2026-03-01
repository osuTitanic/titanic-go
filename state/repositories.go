package state

import (
	"github.com/osuTitanic/common-go/database"
	"github.com/osuTitanic/common-go/repositories"
	"gorm.io/gorm"
)

type Repositories struct {
	// Users
	Users             database.IUserRepository
	Stats             database.IStatsRepository
	Relationships     database.IRelationshipRepository
	Badges            database.IBadgeRepository
	Names             database.INameRepository
	Infringements     database.IInfringementRepository
	Reports           database.IReportRepository
	Verifications     database.IVerificationRepository
	Groups            database.IGroupRepository
	GroupEntries      database.IGroupEntryRepository
	UserPermissions   database.IUserPermissionRepository
	GroupPermissions  database.IGroupPermissionRepository
	Notifications     database.INotificationRepository
	Achievements      database.IAchievementRepository
	BeatmapFavourites database.IBeatmapFavouriteRepository

	// Beatmaps
	// TODO: ...

	// Rankings
	// TODO: ...

	// Forums
	// TODO: ...
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Users:             repositories.NewUserRepository(db),
		Stats:             repositories.NewStatsRepository(db),
		Relationships:     repositories.NewRelationshipRepository(db),
		Badges:            repositories.NewBadgeRepository(db),
		Names:             repositories.NewNameRepository(db),
		Infringements:     repositories.NewInfringementRepository(db),
		Reports:           repositories.NewReportRepository(db),
		Verifications:     repositories.NewVerificationRepository(db),
		Groups:            repositories.NewGroupRepository(db),
		GroupEntries:      repositories.NewGroupEntryRepository(db),
		UserPermissions:   repositories.NewUserPermissionRepository(db),
		GroupPermissions:  repositories.NewGroupPermissionRepository(db),
		Notifications:     repositories.NewNotificationRepository(db),
		Achievements:      repositories.NewAchievementRepository(db),
		BeatmapFavourites: repositories.NewBeatmapFavouriteRepository(db),
	}
}
