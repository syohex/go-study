package models

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type Person struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	IPAddress string `json:"ip_address"`
}

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./db/person.db")
	if err != nil {
		return nil
	}

	DB = db
	return nil
}

func GetPersons(count int) ([]*Person, error) {
	query := fmt.Sprintf("SELECT id, first_name, last_name, email, ip_address from people LIMIT %d", count)
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var ret []*Person
	for rows.Next() {
		p := Person{}
		if err := rows.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Email, &p.IPAddress); err != nil {
			return nil, err
		}

		ret = append(ret, &p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ret, nil
}

func GetPersonByID(id string) (*Person, error) {
	stmt, err := DB.Prepare("SELECT id, first_name, last_name, email, ip_address from people WHERE id = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	p := Person{}
	if err := stmt.QueryRow(id).Scan(&p.ID, &p.FirstName, &p.LastName, &p.Email, &p.IPAddress); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &p, nil
}

func AddPerson(p *Person) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO people (first_name, last_name, email, ip_address) VALUES (?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(p.FirstName, p.LastName, p.Email, p.IPAddress); err != nil {
		tx.Rollback()
		return nil
	}

	tx.Commit()

	return nil
}

func UpdatePerson(p *Person, id int) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("UPDATE people SET first_name = ? last_name = ? email = ? ip_address = ? WHERE id = ?")
	if err != nil {
		tx.Rollback()
		return nil
	}

	defer stmt.Close()

	if _, err := stmt.Exec(p.FirstName, p.LastName, p.Email, p.IPAddress, id); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func DeletePerson(id int) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := DB.Prepare("DELETE from people where id = ?")
	if err != nil {
		tx.Rollback()
		return err
	}

	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
