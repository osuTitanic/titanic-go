package repositories

import (
	"github.com/osuTitanic/common-go/schemas"
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

func (r *NotificationRepository) Update(id int64, updates map[string]interface{}) error {
	return r.db.Model(&schemas.Notification{}).Where("id = ?", id).Updates(updates).Error
}

func (r *NotificationRepository) Delete(notification *schemas.Notification) error {
	return r.db.Delete(notification).Error
}

func (r *NotificationRepository) ById(id int64) (*schemas.Notification, error) {
	var notification schemas.Notification
	err := r.db.Where("id = ?", id).First(&notification).Error
	if err != nil {
		return nil, err
	}
	return &notification, nil
}

func (r *NotificationRepository) ManyByUserId(userId int) ([]*schemas.Notification, error) {
	var notifications []*schemas.Notification
	err := r.db.Where("user_id = ?", userId).Order("time DESC").Find(&notifications).Error
	return notifications, err
}

func (r *NotificationRepository) CountByUserId(userId int) (int, error) {
	var count int64
	err := r.db.Model(&schemas.Notification{}).Where("user_id = ?", userId).Count(&count).Error
	return int(count), err
}
