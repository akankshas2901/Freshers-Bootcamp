package models

import "time"

type Customer struct {
    ID            string    `gorm:"primaryKey"`
    LastOrderTime time.Time
}
