package repositories

import (
	"github.com/osuTitanic/common-go/internal/schemas"
	"gorm.io/gorm"
)

type NotificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
	return &NotificationRepository{db: db}
}

func (r *NotificationRepository) Create(notification *schemas.Notification) error {
	return r.db.Create(notification).Error
}

func (r *NotificationRepository) Delete(notification *schemas.Notification) error {
	return r.db.Delete(notification).Error
}

func (r *NotificationRepository) Update(updates *schemas.Notification, columns ...string) (int64, error) {
	if len(columns) == 0 {
		result := r.db.Save(updates)
		return result.RowsAffected, result.Error
	}
	result := r.db.Model(updates).Select(columns).Updates(updates)
	return result.RowsAffected, result.Error
}

func (r *NotificationRepository) ById(id int64, preload ...string) (*schemas.Notification, error) {
	var notification schemas.Notification
	err := Preloaded(r.db, preload).Where("id = ?", id).First(&notification).Error
	if err != nil {
		return nil, err
	}
	return &notification, nil
}

func (r *NotificationRepository) ManyByUserId(userId int, preload ...string) ([]*schemas.Notification, error) {
	var notifications []*schemas.Notification
	err := Preloaded(r.db, preload).Where("user_id = ?", userId).Order("time DESC").Find(&notifications).Error
	return notifications, err
}

func (r *NotificationRepository) CountByUserId(userId int) (int, error) {
	var count int64
	err := r.db.Model(&schemas.Notification{}).Where("user_id = ?", userId).Count(&count).Error
	return int(count), err
}
