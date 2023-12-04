package employee

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		db: db,
	}
}

// Create implements RepositoryInterface
func (r Repository) Create(emp Employee) (err error) {
	query := `INSERT INTO employees(nip, name, address, created_at, updated_at)
				VALUES ($1, $2, $3, $4, $5)`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(emp.NIP, emp.Name, emp.Address, emp.CreatedAt, emp.UpdatedAt)

	return
}

// DeleteById implements RepositoryInterface
func (r Repository) DeleteById(id int) (err error) {
	query := `DELETE FROM employees WHERE id=$1`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	return
}

// GetAll implements RepositoryInterface
func (r Repository) GetAll() (employees []Employee, err error) {
	query := `SELECT id, nip, name, address, created_at, updated_at FROM employees`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return
	}

	for rows.Next() {
		emp := Employee{}
		err = rows.Scan(
			&emp.Id, &emp.NIP, &emp.Name, &emp.Address,
			&emp.CreatedAt, &emp.UpdatedAt,
		)
		if err != nil {
			return
		}

		employees = append(employees, emp)
	}

	return
}

// UpdateById implements RepositoryInterface
func (r Repository) UpdateById(emp Employee, id int) (err error) {
	query := `UPDATE employees SET nip=$1, name=$2, address=$3, updated_at=$4 WHERE id=$5`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(emp.NIP, emp.Name, emp.Address, emp.UpdatedAt, id)

	return
}

func (r Repository) GetById(id int) (emp Employee, err error) {
	query := `SELECT id, nip, name, address, created_at, updated_at FROM employees WHERE id = $1`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	err = row.Scan(
		&emp.Id, &emp.NIP, &emp.Name,
		&emp.Address, &emp.CreatedAt, &emp.UpdatedAt,
	)

	return
}
