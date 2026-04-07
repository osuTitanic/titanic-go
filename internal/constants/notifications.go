package constants

type NotificationType int

const (
	NotificationTypeWelcome NotificationType = iota
	NotificationTypeAchievement
	NotificationTypeChat
	NotificationTypeForum
	NotificationTypeBeatmaps
	NotificationTypeSecurity
)
