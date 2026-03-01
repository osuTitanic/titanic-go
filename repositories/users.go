package repositories

import (
	"strings"

	"github.com/osuTitanic/common-go/schemas"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *schemas.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) Update(id int, updates map[string]interface{}) error {
	return r.db.Model(&schemas.User{}).Where("id = ?", id).Updates(updates).Error
}

func (r *UserRepository) Delete(user *schemas.User) error {
	return r.db.Delete(user).Error
}

func (r *UserRepository) ById(id int) (*schemas.User, error) {
	var user schemas.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) ByName(name string) (*schemas.User, error) {
	var user schemas.User
	err := r.db.Where("name = ?", name).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) BySafeName(safeName string) (*schemas.User, error) {
	var user schemas.User
	err := r.db.Where("safe_name = ?", safeName).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) ByEmail(email string) (*schemas.User, error) {
	var user schemas.User
	err := r.db.Where("LOWER(email) = ?", strings.ToLower(email)).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) ByDiscordId(discordId int64) (*schemas.User, error) {
	var user schemas.User
	err := r.db.Where("discord_id = ?", discordId).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) ManyById(userIds []int) ([]*schemas.User, error) {
	if len(userIds) == 0 {
		return []*schemas.User{}, nil
	}

	var users []*schemas.User
	err := r.db.Where("id IN ?", userIds).Find(&users).Error
	return users, err
}

func (r *UserRepository) ManyByName(names []string) ([]*schemas.User, error) {
	if len(names) == 0 {
		return []*schemas.User{}, nil
	}

	var users []*schemas.User
	err := r.db.Where("name IN ?", names).Find(&users).Error
	return users, err
}

func (r *UserRepository) ManyByRank(limit int, ascending bool) ([]*schemas.User, error) {
	query := r.db.Model(&schemas.User{}).
		Joins("JOIN stats ON stats.id = users.id").
		Where("users.restricted = ?", false)

	if ascending {
		query = query.Order("stats.rank ASC")
	} else {
		query = query.Order("stats.rank DESC")
	}

	if limit > 0 {
		query = query.Limit(limit)
	}

	var users []*schemas.User
	err := query.Find(&users).Error
	return users, err
}

func (r *UserRepository) ManyByCreationDate(limit int, ascending bool) ([]*schemas.User, error) {
	query := r.db.Model(&schemas.User{}).
		Where("restricted = ?", false).
		Where("activated = ?", true)

	if ascending {
		query = query.Order("created_at ASC")
	} else {
		query = query.Order("created_at DESC")
	}

	if limit > 0 {
		query = query.Limit(limit)
	}

	var users []*schemas.User
	err := query.Find(&users).Error
	return users, err
}

func (r *UserRepository) GetUsername(id int) (string, error) {
	var username string
	err := r.db.Model(&schemas.User{}).
		Where("id = ?", id).
		Select("name").
		Scan(&username).Error
	return username, err
}

func (r *UserRepository) GetUserId(name string) (int, error) {
	var userId int
	err := r.db.Model(&schemas.User{}).
		Where("name = ?", name).
		Select("id").
		Scan(&userId).Error
	return userId, err
}

func (r *UserRepository) GetAvatarChecksum(id int) (string, error) {
	var checksum *string
	err := r.db.Model(&schemas.User{}).
		Where("id = ?", id).
		Select("avatar_hash").
		Scan(&checksum).Error
	if err != nil {
		return "", err
	}
	if checksum == nil {
		return "", nil
	}
	return *checksum, nil
}

func (r *UserRepository) GetCount() (int, error) {
	var count int64
	err := r.db.Model(&schemas.User{}).
		Where("restricted = ?", false).
		Count(&count).Error
	return int(count), err
}
