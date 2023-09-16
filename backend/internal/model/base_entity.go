package model

import (
	"time"
)

type BaseEntity struct {
	AccountID string    `db:"account_id"`
	ID        uint      `gorm:"primarykey" db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
