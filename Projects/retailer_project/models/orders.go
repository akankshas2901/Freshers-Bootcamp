package models

import "time"

type Order struct {
    ID         string    `gorm:"primaryKey"`
    ProductID  string
    CustomerID string
    Quantity   int
    Status     string
    CreatedAt  time.Time
}
