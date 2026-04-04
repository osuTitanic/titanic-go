package state

import (
	"github.com/osuTitanic/titanic-go/internal/repositories"
	"gorm.io/gorm"
)

type Repositories struct {
	// Users
	Users             *repositories.UserRepository
	Stats             *repositories.StatsRepository
	Relationships     *repositories.RelationshipRepository
	Badges            *repositories.BadgeRepository
	Names             *repositories.NameRepository
	Infringements     *repositories.InfringementRepository
	Reports           *repositories.ReportRepository
	Verifications     *repositories.VerificationRepository
	Groups            *repositories.GroupRepository
	GroupEntries      *repositories.GroupEntryRepository
	UserPermissions   *repositories.UserPermissionRepository
	GroupPermissions  *repositories.GroupPermissionRepository
	Notifications     *repositories.NotificationRepository
	Achievements      *repositories.AchievementRepository
	BeatmapFavourites *repositories.BeatmapFavouriteRepository

	// Beatmaps
	Beatmaps    *repositories.BeatmapRepository
	Beatmapsets *repositories.BeatmapsetRepository

	// Rankings
	Scores *repositories.ScoreRepository

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
		Beatmaps:          repositories.NewBeatmapRepository(db),
		Beatmapsets:       repositories.NewBeatmapsetRepository(db),
		Scores:            repositories.NewScoreRepository(db),
	}
}
