//go:build integration

package db

import (
	"context"
	"testing"

	"github.com/TonyPath/go-rest-api/internal/comment"
	"github.com/stretchr/testify/require"
)

func TestCommentDatabase(t *testing.T) {
	t.Run("test create comment", func(t *testing.T) {

		db, err := NewDatabase()
		require.NoError(t, err)

		cmt, err := db.InsertComment(context.TODO(), comment.Comment{
			Slug:   "slug",
			Author: "author",
			Body:   "body",
		})
		require.NoError(t, err)

		newCmt, err := db.QueryComment(context.TODO(), cmt.ID)
		require.NoError(t, err)
		require.Equal(t, "slug", newCmt.Slug)

	})

	t.Run("test delete comment", func(t *testing.T) {

		db, err := NewDatabase()
		require.NoError(t, err)

		cmt, err := db.InsertComment(context.TODO(), comment.Comment{
			Slug:   "slug",
			Author: "author",
			Body:   "body",
		})
		require.NoError(t, err)

		err = db.DeleteComment(context.TODO(), cmt.ID)
		require.NoError(t, err)

		_, err = db.QueryComment(context.TODO(), cmt.ID)
		require.Error(t, err)
	})
}
