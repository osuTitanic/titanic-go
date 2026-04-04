package repositories

import (
	"github.com/osuTitanic/titanic-go/internal/schemas"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(post *schemas.ForumPost) error {
	return r.db.Create(post).Error
}

func (r *PostRepository) Delete(post *schemas.ForumPost) error {
	return r.db.Delete(post).Error
}

func (r *PostRepository) Update(updates *schemas.ForumPost, columns ...string) (int64, error) {
	if len(columns) == 0 {
		result := r.db.Save(updates)
		return result.RowsAffected, result.Error
	}
	result := r.db.Model(updates).Select(columns).Updates(updates)
	return result.RowsAffected, result.Error
}

func (r *PostRepository) UpdateByTopic(topicId int, updates *schemas.ForumPost, columns ...string) (int64, error) {
	if len(columns) == 0 {
		result := r.db.Model(&schemas.ForumPost{}).Where("topic_id = ?", topicId).Updates(updates)
		return result.RowsAffected, result.Error
	}
	result := r.db.Model(&schemas.ForumPost{}).Where("topic_id = ?", topicId).Select(columns).Updates(updates)
	return result.RowsAffected, result.Error
}

func (r *PostRepository) ById(id int64, preload ...string) (*schemas.ForumPost, error) {
	var post schemas.ForumPost
	err := Preloaded(r.db, preload).Where("id = ?", id).First(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) ManyById(ids []int64, preload ...string) ([]*schemas.ForumPost, error) {
	if len(ids) == 0 {
		return []*schemas.ForumPost{}, nil
	}

	var posts []*schemas.ForumPost
	err := Preloaded(r.db, preload).Where("id IN ?", ids).Find(&posts).Error
	return posts, err
}
