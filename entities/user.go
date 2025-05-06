package entities

import "time"

type User struct {
	ID        uint64     `gorm:"primaryKey;autoIncrement;"`
	Name      string     `gorm:"type:varchar(255);not null;"`
	Email     string     `gorm:"type:varchar(255);unique;not null;"`
	Picture   string     `gorm:"type:varchar(255);"`
	CreatedAt time.Time  `gorm:"autoCreateTime;not null;"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime;not null;"`
	Providers []Provider `gorm:"foreignKey:UserID"`
}
