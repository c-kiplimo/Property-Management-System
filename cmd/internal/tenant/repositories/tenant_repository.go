package repositories

import (
	"database/sql"
	"tenant-management/cmd/internal/tenant/domain"
)

type tenantRepository struct {
	db *sql.DB
}

func NewTenantRepository(db *sql.DB) TenantRepository {
	return &tenantRepository{db: db}

}

func (t *tenantRepository) Save(tenant domain.Tenant) error {
	query := `
		INSERT INTO tenants (id, email, mobile_number, names, unit_id)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := t.db.Exec(
		query,
		tenant.ID,
		tenant.Email,
		tenant.MobileNumber,
		tenant.Names,
		tenant.UnitID,
	)

	return err
}

func (t *tenantRepository) FindByTenantId(tenantId string) (domain.Tenant, error) {
	query := `
		SELECT id, email, mobile_number, names, unit_id
		FROM tenants
		WHERE id = $1
	`

	var tenant domain.Tenant

	err := t.db.QueryRow(query, tenantId).Scan(
		&tenant.ID,
		&tenant.Email,
		&tenant.MobileNumber,
		&tenant.Names,
		&tenant.UnitID,
	)

	if err != nil {
		return domain.Tenant{}, err
	}

	return tenant, nil
}
