package menu

import (
	"database/sql"
	"time"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		db: db,
	}
}

func (r Repository) Save(menu Menu) (err error) {
	query := `INSERT INTO menus(name, category, description, price, image_url, created_at, updated_at)
				VALUES ($1, $2, $3, $4, $5, $6, $7)`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(menu.Name, menu.Category, menu.Desc, menu.Price, menu.ImageUrl, menu.CreatedAt, menu.UpdatedAt)
	return
}

func (r Repository) GetAll() (menus []Menu, err error) {
	query := `SELECT id, name, category, description, price, image_url, created_at, updated_at FROM menus`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		menu := Menu{}

		err = rows.Scan(
			&menu.Id, &menu.Name, &menu.Category,
			&menu.Desc, &menu.Price, &menu.ImageUrl,
			&menu.CreatedAt, &menu.UpdatedAt,
		)
		if err != nil {
			return
		}

		menus = append(menus, menu)
	}

	return
}

func (r Repository) GetById(id int) (menu Menu, err error) {
	query := `SELECT id, name, category, description, price, image_url, created_at, updated_at FROM menus WHERE id = $1`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	err = row.Scan(
		&menu.Id, &menu.Name, &menu.Category,
		&menu.Desc, &menu.Price, &menu.ImageUrl,
		&menu.CreatedAt, &menu.UpdatedAt,
	)

	return
}

func (r Repository) Update(req Menu) (err error) {
	query := `UPDATE menus SET name=$1, category=$2, description=$3, price=$4, image_url=$5, updated_at=$6 WHERE id=$7`

	// Set prepare statement
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		req.Name, req.Category,
		req.Desc, req.Price,
		req.ImageUrl, time.Now(), req.Id,
	)

	return
}

func (r Repository) Delete(id int) (err error) {
	query := `DELETE FROM menus WHERE id=$1`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec(id)
	return
}
