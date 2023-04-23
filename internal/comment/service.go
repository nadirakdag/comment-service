package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrNotImplemented  = errors.New("not implemented")
)

type CommentService interface {
	GetComment(context.Context, string) (Comment, error)
	UpdateComment(context.Context, string, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
	CreateComment(context.Context, Comment) (Comment, error)
}

// Service - is the struct on which all our
// logic will be built on top of
type Service struct {
	Store Store
}

// NewService - returns a new pointer to a new
// service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(context context.Context, id string) (Comment, error) {
	fmt.Println("retrieving a comment ")

	cmt, err := s.Store.GetComment(context, id)
	if err != nil {
		fmt.Printf("error while retrieving comment, err: %v \n", err)
		return Comment{}, ErrFetchingComment
	}

	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, id string, cmt Comment) (Comment, error) {
	fmt.Println("updating comment")

	updatedComment, err := s.Store.UpdateComment(ctx, id, cmt)
	if err != nil {
		return Comment{}, err
	}

	return updatedComment, nil
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return s.Store.DeleteComment(ctx, id)
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	fmt.Println("creating comment")

	insertedComment, err := s.Store.CreateComment(ctx, cmt)
	if err != nil {
		return Comment{}, err
	}

	return insertedComment, nil
}
