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
func NewAuthorMysqlRepo(conn string) (AuthorMysqlRepo, error) {
	db, err := sql.Open("mysql", conn)
	if err != nil {
		return AuthorMysqlRepo{}, err
	}
	return AuthorMysqlRepo{DB: db}, nil
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

	dbrow, err := a.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer dbrow.Close()
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
	query := "Delete from  Author where true=true"
	for _, v := range ids {
		query += fmt.Sprintf(" AND `id`='%s'", v)
	}

	_, err := a.DB.Exec(query)

	return err
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
	tx, err := a.DB.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	_, err = tx.Exec(query)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()

		return nil, err
	}
	return authors, nil
}
