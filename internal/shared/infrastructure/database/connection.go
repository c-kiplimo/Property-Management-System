package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// Import all domain models for migration
	documentDomain "tenant-management/internal/domains/document/domain/entities"
	financialDomain "tenant-management/internal/domains/financial/domain/entities"
	identityDomain "tenant-management/internal/domains/identity/domain/entities"
	leasingDomain "tenant-management/internal/domains/leasing/domain/entities"
	maintenanceDomain "tenant-management/internal/domains/maintenance/domain/entities"
	propertyDomain "tenant-management/internal/domains/property/domain/entities"
	tenantDomain "tenant-management/internal/domains/tenant/domain/entities"
)

func Initialize(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate all models
	if err := db.AutoMigrate(
		&identityDomain.Landlord{},
		&propertyDomain.Property{},
		&propertyDomain.Unit{},
		&tenantDomain.Tenant{},
		&leasingDomain.Lease{},
		&financialDomain.Payment{},
		&financialDomain.Expense{},
		&maintenanceDomain.MaintenanceRequest{},
		&documentDomain.Document{},
	); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database connected and migrated successfully")
	return db
}
