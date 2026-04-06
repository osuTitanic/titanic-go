package repositories

import (
	"github.com/osuTitanic/titanic-go/internal/schemas"
	"gorm.io/gorm"
)

type TopicRepository struct {
	db *gorm.DB
}

func NewTopicRepository(db *gorm.DB) *TopicRepository {
	return &TopicRepository{db: db}
}

func (r *TopicRepository) Create(topic *schemas.ForumTopic) error {
	return r.db.Create(topic).Error
}

func (r *TopicRepository) Delete(topic *schemas.ForumTopic) error {
	return r.db.Delete(topic).Error
}

func (r *TopicRepository) Update(updates *schemas.ForumTopic, columns ...string) (int64, error) {
	return CommonUpdate(r.db, updates, columns...)
}

func (r *TopicRepository) ById(id int, preload ...string) (*schemas.ForumTopic, error) {
	var topic schemas.ForumTopic
	err := Preloaded(r.db, preload).Where("id = ?", id).First(&topic).Error
	if err != nil {
		return nil, err
	}
	return &topic, nil
}

func (r *TopicRepository) ManyById(ids []int, preload ...string) ([]*schemas.ForumTopic, error) {
	if len(ids) == 0 {
		return []*schemas.ForumTopic{}, nil
	}

	var topics []*schemas.ForumTopic
	err := Preloaded(r.db, preload).Where("id IN ?", ids).Find(&topics).Error
	return topics, err
}
