package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	model "github.com/team-a-hacks/tas-blog-api/comment"
	"github.com/team-a-hacks/tas-blog-api/status"
)

type commentRepository struct {
	Conn *gorm.DB
}

// NewCommentRepository mount comment repository
func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{
		Conn: db,
	}
}

// CommentRepository comment repository interface
type CommentRepository interface {
	ListByArticleID(id uuid.UUID) ([]*model.Comment, error)
	Create(comment *model.Comment) (*model.Comment, error)
}

func (m *commentRepository) ListByArticleID(id uuid.UUID) ([]*model.Comment, error) {
	comments := []*model.Comment{}
	err := m.Conn.Model(&comments).Where("article_id = ?", id).Find(&comments).Error
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (m *commentRepository) Create(comment *model.Comment) (*model.Comment, error) {
	err := m.Conn.Create(comment).Error
	if err != nil {
		return nil, errors.Wrap(status.ErrInternalServer, err.Error())
	}
	return comment, nil
}
