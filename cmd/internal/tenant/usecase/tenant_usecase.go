package usecase

import (
	"database/sql"
	"errors"
	"tenant-management/cmd/internal/tenant/domain"
	"tenant-management/cmd/internal/tenant/repositories"
)

type TenantUseCase struct {
	repo repositories.TenantRepository
}

func NewTenantUseCase(repo repositories.TenantRepository) *TenantUseCase {
	return &TenantUseCase{repo: repo}
}

// CreateTenant creates a new tenant if it doesn't already exist
func (t *TenantUseCase) CreateTenant(tenant domain.Tenant) error {
	_, err := t.repo.FindByTenantId(tenant.ID)
	if err == nil {
		return errors.New("tenant already exists")
	}

	// Optional: handle only "not found" error and return others
	if !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	return t.repo.Save(tenant)
}

// FindTenantById retrieves tenant details by ID
func (t *TenantUseCase) FindTenantById(id string) (domain.Tenant, error) {
	return t.repo.FindByTenantId(id)
}
