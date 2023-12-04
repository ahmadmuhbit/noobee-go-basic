package transaction

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		db: db,
	}
}

// GetEmployeeById implements RepositoryInterface
func (r Repository) GetEmployeeById(id int) (emp Employee, err error) {
	query := `SELECT id, name FROM employees WHERE id=$1`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(
		&emp.Id, &emp.Name,
	)
	return
}

// GetMenuById implements RepositoryInterface
func (r Repository) GetMenuById(id int) (menu Menu, err error) {
	query := `SELECT id, name, price FROM menus where id=$1`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(
		&menu.Id, &menu.Name, &menu.Price,
	)
	return
}

// Create implements RepositoryInterface
func (r Repository) Create(trx Transaction) (err error) {
	query := `INSERT INTO transactions(menu_id, employee_id, quantity, total, created_at, updated_at)
				VALUES($1, $2, $3, $4, $5, $6)`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		trx.MenuId, trx.EmployeeId, trx.Quantity,
		trx.Total, trx.CreatedAt, trx.UpdatedAt,
	)

	return
}

// FindAll implements RepositoryInterface
func (r Repository) FindAll() (trxs []Transaction, err error) {
	query := `SELECT t.id, t.employee_id, t.menu_id, t.quantity, t.total, t.created_at, t.updated_at, 
				m.id, m.name, m.price, 
				e.id, e.name
				FROM transactions as t
				JOIN employees as e
					ON e.id = t.employee_id
				JOIN menus as m
					ON m.id = t.menu_id`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return
	}
	defer stmt.Close()

	for rows.Next() {
		trx := Transaction{}
		err = rows.Scan(
			&trx.Id, &trx.EmployeeId, &trx.MenuId, &trx.Quantity, &trx.Total, &trx.CreatedAt, &trx.UpdatedAt,
			&trx.Menu.Id, &trx.Menu.Name, &trx.Menu.Price,
			&trx.Employee.Id, &trx.Employee.Name,
		)

		if err != nil {
			return
		}

		trxs = append(trxs, trx)
	}

	return
}

// FindById implements RepositoryInterface
func (r Repository) FindById(id int) (trx Transaction, err error) {
	query := `SELECT t.id, t.employee_id, t.menu_id, t.quantity, t.total, t.created_at, t.updated_at, 
				m.id, m.name, m.price, 
				e.id, e.name
				FROM transactions as t
				JOIN employees as e
					ON e.id = t.employee_id
				JOIN menus as m
					ON m.id = t.menu_id
				WHERE t.id=$1`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(
		&trx.Id, &trx.EmployeeId, &trx.MenuId, &trx.Quantity, &trx.Total, &trx.CreatedAt, &trx.UpdatedAt,
		&trx.Menu.Id, &trx.Menu.Name, &trx.Menu.Price,
		&trx.Employee.Id, &trx.Employee.Name,
	)
	return
}
