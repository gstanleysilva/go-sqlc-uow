// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: reports.sql

package db

import (
	"context"
	"database/sql"
)

const getCoursesWithCategories = `-- name: GetCoursesWithCategories :many
SELECT  c.id as CourseID,
        c.name as CourseName, 
        c.description as CourseDescription, 
        c.price as CoursePrice, 
        ct.id as CategoryID,
        ct.name as CategoryName,
        ct.description as CategoryDescription
FROM courses c
LEFT JOIN categories ct ON c.category_id = ct.id
GROUP BY c.id
`

type GetCoursesWithCategoriesRow struct {
	Courseid            string
	Coursename          string
	Coursedescription   sql.NullString
	Courseprice         float64
	Categoryid          sql.NullString
	Categoryname        sql.NullString
	Categorydescription sql.NullString
}

func (q *Queries) GetCoursesWithCategories(ctx context.Context) ([]GetCoursesWithCategoriesRow, error) {
	rows, err := q.db.QueryContext(ctx, getCoursesWithCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCoursesWithCategoriesRow
	for rows.Next() {
		var i GetCoursesWithCategoriesRow
		if err := rows.Scan(
			&i.Courseid,
			&i.Coursename,
			&i.Coursedescription,
			&i.Courseprice,
			&i.Categoryid,
			&i.Categoryname,
			&i.Categorydescription,
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