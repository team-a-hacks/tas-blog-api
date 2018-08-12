package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/team-a-hacks/tas-blog-api/article"
)

type articleRepository struct {
	Conn *gorm.DB
}

// NewArticleRepository mount article repository
func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{
		Conn: db,
	}
}

// ArticleRepository repository interface
type ArticleRepository interface {
	List() ([]*article.Article, error)
}

func (m *articleRepository) List() ([]*article.Article, error) {
	var articles = []*article.Article{}
	err := m.Conn.Model(&articles).Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}
