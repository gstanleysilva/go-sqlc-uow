// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: courses.sql

package db

import (
	"context"
	"database/sql"
)

const createCourse = `-- name: CreateCourse :exec
insert into courses (id, name, description, category_id, price)
VALUES (?, ?, ?, ?, ?)
`

type CreateCourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	CategoryID  string
	Price       float64
}

func (q *Queries) CreateCourse(ctx context.Context, arg CreateCourseParams) error {
	_, err := q.db.ExecContext(ctx, createCourse,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.CategoryID,
		arg.Price,
	)
	return err
}

const deleteCourse = `-- name: DeleteCourse :exec
DELETE
FROM courses
WHERE id = ?
`

func (q *Queries) DeleteCourse(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteCourse, id)
	return err
}

const getCourse = `-- name: GetCourse :one
SELECT id, category_id, name, description, price
FROM courses
WHERE id = ?
`

func (q *Queries) GetCourse(ctx context.Context, id string) (Course, error) {
	row := q.db.QueryRowContext(ctx, getCourse, id)
	var i Course
	err := row.Scan(
		&i.ID,
		&i.CategoryID,
		&i.Name,
		&i.Description,
		&i.Price,
	)
	return i, err
}

const listCourses = `-- name: ListCourses :many
SELECT id, category_id, name, description, price
FROM courses
ORDER BY id
`

func (q *Queries) ListCourses(ctx context.Context) ([]Course, error) {
	rows, err := q.db.QueryContext(ctx, listCourses)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Course
	for rows.Next() {
		var i Course
		if err := rows.Scan(
			&i.ID,
			&i.CategoryID,
			&i.Name,
			&i.Description,
			&i.Price,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCourse = `-- name: UpdateCourse :exec
UPDATE courses
SET name = ?,
    description = ?,
    category_id = ?,
    price = ?
WHERE id = ?
`

type UpdateCourseParams struct {
	Name        string
	Description sql.NullString
	CategoryID  string
	Price       float64
	ID          string
}

func (q *Queries) UpdateCourse(ctx context.Context, arg UpdateCourseParams) error {
	_, err := q.db.ExecContext(ctx, updateCourse,
		arg.Name,
		arg.Description,
		arg.CategoryID,
		arg.Price,
		arg.ID,
	)
	return err
}