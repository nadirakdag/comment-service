package comment

import (
	"context"
	"fmt"
)

// Store - this interface all the methods
// that our service needs in order to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// Comment - a representation of the comment
// structure of our service
type Comment struct {
	Id     string
	Slug   string
	Body   string
	Author string
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
		fmt.Errorf("error while retrieving comment, err: %v", err)
		return Comment{}, err
	}

	return cmt, nil
}
