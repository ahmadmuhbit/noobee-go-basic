package auth

import "database/sql"

func NewRepository(db *sql.DB) Repository {
	return Repository{
		db: db,
	}
}

type Repository struct {
	db *sql.DB
}

func (r Repository) Create(auth Auth) (err error) {
	query := `
		INSERT INTO auth (email, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4) 
		`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(auth.Email, auth.Password, auth.CreatedAt, auth.UpdatedAt)
	return
}

func (r Repository) GetByEmail(email string) (auth Auth, err error) {
	query := `
		SELECT
		id, email, password, created_at, updated_at
		FROM auth
		WHERE email = $1
		`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()
	row := stmt.QueryRow(email)

	err = row.Scan(
		&auth.Id, &auth.Email, &auth.Password, &auth.CreatedAt, &auth.UpdatedAt,
	)

	return
}
