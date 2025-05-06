package entities

import "time"

type Provider struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement;"`
	UserID     uint64    `gorm:"not null;"`
	Provider   string    `gorm:"type:varchar(255);not null;uniqueIndex:idx_provider_provider_id"` // e.g., Google, Facebook
	ProviderID string    `gorm:"type:varchar(50);not null;uniqueIndex:idx_provider_provider_id"`  // Google ID, Facebook ID, etc.
	CreatedAt  time.Time `gorm:"autoCreateTime;not null;"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime;not null;"`
	User       *User     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
