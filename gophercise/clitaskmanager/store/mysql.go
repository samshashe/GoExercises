package store

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var GlobalMySQLDB *sql.DB

type DBTaskStore struct {
	db *sql.DB
}

func NewDBTaskStore() TaskStore {
	return &DBTaskStore{
		db: GlobalMySQLDB,
	}
}

func NewMySQLDB(dsn string, dbName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn+"?parseTime=true")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("USE " + dbName)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Task ( ID INT, Name VARCHAR(255),  Completed BOOLEAN, CreatedDate DATETIME)")
	if err != nil {
		panic(err)
	}

	return db, db.Ping()
}

func (store *DBTaskStore) GetAll() ([]Task, error) {
	rows, err := store.db.Query(
		`
		SELECT ID,Name,Completed,CreatedDate
		FROM Task
		ORDER BY CreatedDate
		`)

	if err != nil {
		return nil, err
	}

	tasks := []Task{}
	for rows.Next() {
		task := Task{}
		err := rows.Scan(
			&task.ID,
			&task.Name,
			&task.Completed,
			&task.CreatedDate,
		)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (store *DBTaskStore) Add(task Task) error {
	_, err := store.db.Exec(
		`
		REPLACE INTO Task
			(ID,Name,Completed,CreatedDate)
		VALUES
			(?, ?, ?, ?)
		`,
		task.ID,
		task.Name,
		task.Completed,
		task.CreatedDate,
	)
	return err
}

func (store *DBTaskStore) ToggleCompleted(name string) error {
	_, err := store.db.Exec(
		`
		UPDATE Task
		SET Completed = !Completed
		WHERE Name = ?
		`,
		name,
	)
	return err
}
