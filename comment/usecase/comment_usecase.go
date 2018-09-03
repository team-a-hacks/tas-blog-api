package usecase

import (
	"database/sql"

	uuid "github.com/satori/go.uuid"
	commentModel "github.com/team-a-hacks/tas-blog-api/comment"
	commentRepo "github.com/team-a-hacks/tas-blog-api/comment/repository"
)

type commentUsecase struct {
	commentRepo commentRepo.CommentRepository
}

// NewCommentUsecase mount comment usecase
func NewCommentUsecase(comment commentRepo.CommentRepository) CommentUsecase {
	return &commentUsecase{
		commentRepo: comment,
	}
}

// CommentUsecase comment usecase interface
type CommentUsecase interface {
	Create(payload *commentModel.Payload) (*commentModel.ResponseData, error)
}

func (c *commentUsecase) Create(payload *commentModel.Payload) (*commentModel.ResponseData, error) {
	comment := commentModel.Comment{
		ID:        uuid.NewV4(),
		Content:   payload.Content,
		AuthorID:  payload.AuthorID,
		ArticleID: payload.ArticleID,
		CreatedBy: payload.CreatedBy,
		UpdatedBy: payload.CreatedBy,
	}

	if payload.Author == "" {
		comment.Author = sql.NullString{String: payload.Author, Valid: false}
	} else {
		comment.Author = sql.NullString{String: payload.Author, Valid: true}
	}
	resComment, err := c.commentRepo.Create(&comment)
	if err != nil {
		return nil, err
	}

	res := commentModel.ResponseData{
		ID:        resComment.ID,
		Content:   resComment.Content,
		ArticleID: resComment.ArticleID,
		PostDate:  resComment.CreatedAt.Format("2006-01-02"),
	}
	if !resComment.Author.Valid {
		res.Author = "匿名"
	} else {
		res.Author = resComment.Author.String
	}
	return &res, nil
}
