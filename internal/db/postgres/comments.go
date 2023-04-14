package postgres

import (
	"comment-service/internal/comment"
	"context"
	"database/sql"
	uuid "github.com/satori/go.uuid"
)

type CommentRow struct {
	Id     string
	Slug   sql.NullString
	Body   sql.NullString
	Author sql.NullString
}

func (c *CommentRow) projectToComment() comment.Comment {
	return comment.Comment{
		Id:     c.Id,
		Slug:   c.Slug.String,
		Body:   c.Body.String,
		Author: c.Author.String,
	}
}

func (d *Database) GetComment(ctx context.Context, id string) (comment.Comment, error) {

	commentRow := CommentRow{}

	row := d.Client.QueryRowContext(ctx, `SELECT id, slug, body, author FROM comments WHERE id=$1`, id)
	err := row.Scan(&commentRow.Id, &commentRow.Slug, &commentRow.Body, &commentRow.Author)
	if err != nil {
		return comment.Comment{}, err
	}

	return commentRow.projectToComment(), nil
}

func (d *Database) CreateComment(ctx context.Context, cmt comment.Comment) (comment.Comment, error) {
	cmt.Id = uuid.NewV4().String()

	postRow := CommentRow{
		Id:     cmt.Id,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
	}

	row, err := d.Client.NamedQueryContext(ctx, `INSERT INTO comments(id, slug, author, body) VALUES (:id, :slug, :author, :body)`, postRow)
	if err != nil {
		return comment.Comment{}, err
	}

	if err := row.Close(); err != nil {
		return comment.Comment{}, err
	}

	return cmt, nil
}

func (d *Database) DeleteComment(ctx context.Context, id string) error {

	_, err := d.Client.ExecContext(ctx, `DELETE FROM comments WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) UpdateComment(ctx context.Context, id string, cmt comment.Comment) (comment.Comment, error) {

	var row = CommentRow{
		Id:     id,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
	}
	updatedRow, err := d.Client.NamedQueryContext(ctx, `UPDATE comments SET slug = :slug, author = :author, body = :body WHERE id = :id`, row)
	if err != nil {
		return comment.Comment{}, err
	}

	if err := updatedRow.Close(); err != nil {
		return comment.Comment{}, err
	}

	return cmt, nil
}
