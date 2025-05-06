package repositories

import "tenant-management/cmd/internal/tenant/domain"

type TenantRepository interface {
	Save(tenant domain.Tenant) error
	FindByTenantId(tenantId string) (domain.Tenant, error)
}
