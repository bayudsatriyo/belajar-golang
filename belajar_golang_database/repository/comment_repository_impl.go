package repository

import (
	"belajar_golang_database/entity"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{
		DB: db,
	}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "INSERT INTO comment(email, comment) VALUES (?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, comment.Email, comment.Comment)

	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return comment, err
	}

	comment.Id = int32(id)
	return comment, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, commentId int32) (entity.Comment, error) {
	script := "SELECT id, email, comment FROM comment WHERE id = ?"
	rows, err := repository.DB.QueryContext(ctx, script, commentId)
	comment := entity.Comment{}

	if err != nil {
		return entity.Comment{}, err
	}

	defer rows.Close()

	if rows.Next() {
		// ada 
		err := rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		if err != nil {
			return entity.Comment{}, err
		}
		return comment, nil
	} else {
		// tidak ada
		return comment, errors.New("Id " + strconv.Itoa(int(commentId)) + "is Not Found")
	} 
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "SELECT id, email, comment FROM comment"
	rows, err := repository.DB.QueryContext(ctx, script)

	if err != nil {
		return []entity.Comment{}, err
	}

	defer rows.Close()

	var comment []entity.Comment

	for rows.Next() {
		commentTemp := entity.Comment{}
		rows.Scan(&commentTemp.Id, &commentTemp.Email, &commentTemp.Comment)
		// if err != nil {
		// 	return []entity.Comment{}, err
		// }
		comment = append(comment, commentTemp)
	}

	return comment, nil
}