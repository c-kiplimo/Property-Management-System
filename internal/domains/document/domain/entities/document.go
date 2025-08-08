package entities

import (
	"time"
)

type EntityType string

const (
	EntityTypeProperty    EntityType = "property"
	EntityTypeTenant      EntityType = "tenant"
	EntityTypeLease       EntityType = "leasing"
	EntityTypeMaintenance EntityType = "maintenance"
)

type Document struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	EntityType  EntityType `json:"entity_type" gorm:"not null"`
	EntityID    uint       `json:"entity_id" gorm:"not null"`
	FileName    string     `json:"file_name" gorm:"not null"`
	FilePath    string     `json:"file_path" gorm:"not null"`
	FileSize    int64      `json:"file_size"`
	MimeType    string     `json:"mime_type"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func NewDocument(entityType EntityType, entityID uint, fileName, filePath string) *Document {
	return &Document{
		EntityType: entityType,
		EntityID:   entityID,
		FileName:   fileName,
		FilePath:   filePath,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

func (d *Document) UpdateMetadata(fileSize int64, mimeType, description string) {
	d.FileSize = fileSize
	d.MimeType = mimeType
	d.Description = description
	d.UpdatedAt = time.Now()
}
