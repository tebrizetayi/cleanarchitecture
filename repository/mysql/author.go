package mysql

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/tebrizetayi/cleanarchitecture/domain/model"
)

type AuthorMysqlRepo struct {
	DB *sql.DB
}

//root:pass1@tcp(127.0.0.1:3306)/tuts
func NewAuthorMysqlRepo(db *sql.DB) AuthorMysqlRepo {

	return AuthorMysqlRepo{DB: db}
}

func (a *AuthorMysqlRepo) Create(Authors []model.Author) ([]model.Author, error) {

	result := []model.Author{}
	query := "Insert into `Author` (`Id`,`Name`) Values "
	for i, v := range Authors {
		v.ID = uuid.New()

		if i > 0 {
			query += ","
		}
		query += fmt.Sprintf("('%s','%s')\n", v.ID, v.Name)
		result = append(result, v)
	}

	var tx *sql.Tx
	var err error

	if tx, err = a.DB.BeginTx(context.Background(), nil); err != nil {
		return nil, err
	}

	var dbRow *sql.Rows
	if dbRow, err = tx.Query(query); err != nil {
		tx.Rollback()
		return nil, err
	}
	defer dbRow.Close()

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}

	return result, nil
}

func (a *AuthorMysqlRepo) GetAll() ([]model.Author, error) {

	query := "Select `Id`,`Name` From `Author` where true=true Limit 10"
	dbrow, err := a.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer dbrow.Close()

	result := []model.Author{}
	for dbrow.Next() {
		var newAuthor model.Author
		err = dbrow.Scan(&newAuthor.ID, &newAuthor.Name)
		if err != nil {
			return nil, err
		}
		result = append(result, newAuthor)
	}
	return result, nil
}

func (a *AuthorMysqlRepo) Delete(ids []uuid.UUID) error {
	query := "Delete from Author where true=true"
	for _, v := range ids {
		query += fmt.Sprintf(" AND `id`='%s'", v)
	}

	var tx *sql.Tx
	var err error
	if tx, err = a.DB.BeginTx(context.Background(), nil); err != nil {
		return err
	}

	if _, err = tx.Exec(query); err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (a *AuthorMysqlRepo) GetByIds(ids []uuid.UUID) ([]model.Author, error) {

	query := "Select Id, Name from  Author where true=true"
	for _, v := range ids {
		query += fmt.Sprintf(" AND Id='%s'", v)
	}

	dbrow, err := a.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer dbrow.Close()

	result := []model.Author{}
	for dbrow.Next() {
		var newAuthor model.Author
		err = dbrow.Scan(&newAuthor.ID, &newAuthor.Name)
		if err != nil {
			return nil, err
		}
		result = append(result, newAuthor)
	}
	return result, nil
}

func (a *AuthorMysqlRepo) Reset() {

}

func (a *AuthorMysqlRepo) Update(authors []model.Author) ([]model.Author, error) {
	if authors == nil || len(authors) == 0 {
		return []model.Author{}, nil
	}

	updatePhrase := "UPDATE Author SET `Name`='%s' WHERE `Id`='%s';"
	query := ""
	for _, author := range authors {
		query += fmt.Sprintf(updatePhrase, author.Name, author.ID)
	}

	var tx *sql.Tx
	var err error

	if tx, err = a.DB.BeginTx(context.Background(), nil); err != nil {
		return nil, err
	}

	if _, err = tx.Exec(query); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}

	return authors, nil
}
