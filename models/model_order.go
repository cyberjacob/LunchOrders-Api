/*
 * Lunch Orders
 *
 * This API allows lunch ordering.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package models

import (
	"github.com/cuvva/ksuid"
)

type OrderStatus int

type Order struct {
	Id     ksuid.ID     `json:"id,omitempty" gorm:"type:blob"`
	From   string       `json:"from,omitempty"`
	Items  []OrderItems `json:"items,omitempty"`
	Status OrderStatus  `json:"status,omitempty"`

	Session      Session  `gorm:"foreignkey:SessionRefer"` // use UserRefer as foreign key
	SessionRefer ksuid.ID `gorm:"type:blob"`
}

const (
	Pending OrderStatus = iota + 1
	Rejected
	Canceled
	Accepted
	Complete
)

func (day OrderStatus) String() string {
	names := [...]string{
		"Pending",
		"Rejected",
		"Canceled",
		"Accepted",
		"Complete",
	}

	// prevent panicking in case of
	// `day` is out of range of Weekday
	if day < Pending || day > Complete {
		return "Unknown"
	}

	// return the name of a Weekday
	// constant from the names array
	// above.
	return names[day]
}
