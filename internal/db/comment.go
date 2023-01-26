package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/TonyPath/go-rest-api/internal/comment"
	uuid "github.com/satori/go.uuid"
)

type CommentDB struct {
	ID     string         `db:"id"`
	Slug   sql.NullString `db:"slug"`
	Body   sql.NullString `db:"body"`
	Author sql.NullString `db:"author"`
}

func toComment(c CommentDB) comment.Comment {
	return comment.Comment{
		ID:     c.ID,
		Slug:   c.Slug.String,
		Body:   c.Body.String,
		Author: c.Author.String,
	}
}

func (d *Database) QueryComment(ctx context.Context, id string) (comment.Comment, error) {
	const q = `
	SELECT 
		id,
		slug,
		body,
		author
	FROM 
		comments
	WHERE
		id = $1`

	row := d.DBConn.QueryRowContext(ctx, q, id)

	var cmtDB CommentDB
	err := row.Scan(&cmtDB.ID, &cmtDB.Slug, &cmtDB.Body, &cmtDB.Author)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("error fetching comment by uuid: %w", err)
	}

	return toComment(cmtDB), nil
}

func (d *Database) InsertComment(ctx context.Context, cmt comment.Comment) (comment.Comment, error) {
	cmt.ID = uuid.NewV4().String()

	newCommentDB := CommentDB{
		ID:     cmt.ID,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
	}

	const q = `
	INSERT INTO comments
	(id, slug, body, author)
	VALUES
	(:id, :slug, :body, :author)
	`

	_, err := d.DBConn.NamedExecContext(ctx, q, newCommentDB)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("failed to insert comment: %w", err)
	}

	return cmt, nil
}

func (d *Database) DeleteComment(ctx context.Context, id string) error {
	const q = `DELETE FROM comments WHERE id = $1`

	_, err := d.DBConn.ExecContext(ctx, q, id)
	if err != nil {
		return fmt.Errorf("failed to delete comment: %w", err)
	}

	return nil
}

func (d *Database) UpdateComment(ctx context.Context, id string, cmt comment.Comment) (comment.Comment, error) {
	updateCommentDB := CommentDB{
		ID:     id,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
	}

	const q = `
	UPDATE comments SET
		slug = :slug,
		body = :body,
		author = :author
	WHERE
		id = :id
	`

	_, err := d.DBConn.NamedExecContext(ctx, q, updateCommentDB)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("failed to insert comment: %w", err)
	}

	return toComment(updateCommentDB), nil
}
