package menu

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		db: db,
	}
}

func (r Repository) Save(menu Menu) (err error) {
	query := `INSERT INTO menus(name, category, description, price, created_at, updated_at)
				VALUES ($1, $2, $3, $4, $5, $6)`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(menu.Name, menu.Category, menu.Desc, menu.Price, menu.CreatedAt, menu.UpdatedAt)
	return
}

func (r Repository) GetAll() (menus []Menu, err error) {
	query := `SELECT id, name, category, description, price, created_at, updated_at FROM menus`

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
			&menu.Desc, &menu.Price,
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
	query := `SELECT id, name, category, description, price, created_at, updated_at FROM menus WHERE id = $1`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	err = row.Scan(
		&menu.Id, &menu.Name, &menu.Category,
		&menu.Desc, &menu.Price,
		&menu.CreatedAt, &menu.UpdatedAt,
	)

	return
}
