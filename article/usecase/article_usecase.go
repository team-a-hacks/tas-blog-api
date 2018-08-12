package usecase

import (
	model "github.com/team-a-hacks/tas-blog-api/article"
	"github.com/team-a-hacks/tas-blog-api/article/repository"
)

type articleUsecase struct {
	articleRepo repository.ArticleRepository
}

// NewArticleUsecase mount article usecase
func NewArticleUsecase(article repository.ArticleRepository) ArticleUsecase {
	return &articleUsecase{
		articleRepo: article,
	}
}

// ArticleUsecase article usecase interface
type ArticleUsecase interface {
	List() ([]*model.ArticleData, error)
}

// List list article
func (a *articleUsecase) List() ([]*model.ArticleData, error) {
	articles, err := a.articleRepo.List()
	if err != nil {
		return nil, err
	}
	var articleData = []*model.ArticleData{}
	for _, a := range articles {
		article := formatArticleData(a)
		articleData = append(articleData, article)
	}
	return articleData, nil
}

func formatArticleData(a *model.Article) *model.ArticleData {
	var articleData = &model.ArticleData{
		ID:        a.ID,
		Content:   a.Content,
		CreatedAt: a.CreatedAt,
	}
	return articleData
}
