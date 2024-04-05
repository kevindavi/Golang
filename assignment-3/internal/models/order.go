package models

import "time"

type Order struct {
	OrderID      int       `json:"orderId"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `json:"items"`
}
