package postgres

import (
	"context"
	"office/internal/types"

	"github.com/jmoiron/sqlx"
)

type Repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) CreateOffice(ctx context.Context, data types.OfficeMake) error {
	query := `
		INSERT INTO office (uuid, name, address, created_at)
		VALUES(:uuid, :name, :address, :created_at)
	`
	_, err := r.db.NamedExecContext(ctx, query, data)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetOffices(ctx context.Context) (types.OfficeList, error) {
	query := `
		SELECT uuid, name, address, created_at
		FROM office
	`

	var result types.OfficeList

	if err := r.db.SelectContext(ctx, &result, query); err != nil {
		return result, err
	}
	return result, nil
}
