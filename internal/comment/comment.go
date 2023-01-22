package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
)

type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

type Store interface {
	QueryComment(ctx context.Context, id string) (Comment, error)
	InsertComment(ctx context.Context, cmt Comment) (Comment, error)
	DeleteComment(ctx context.Context, id string) error
	UpdateComment(ctx context.Context, id string, cmt Comment) (Comment, error)
}

type Service struct {
	Store Store
}

func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	cmt, err := s.Store.QueryComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	comment, err := s.Store.InsertComment(ctx, cmt)
	if err != nil {
		return Comment{}, err
	}
	return comment, nil
}

func (s *Service) UpdateComment(ctx context.Context, ID string, uComment Comment) (Comment, error) {
	comment, err := s.Store.UpdateComment(ctx, ID, uComment)
	if err != nil {
		return Comment{}, err
	}
	return comment, nil
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return s.Store.DeleteComment(ctx, id)
}
