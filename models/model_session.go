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

type Session struct {
	Id ksuid.ID `json:"id,omitempty" gorm:"type:blob"`
	Organiser string `json:"organiser,omitempty"`
	From string `json:"from,omitempty"`
	Cutoff string `json:"cutoff,omitempty"`
	Payment string `json:"payment,omitempty"`
	Notes string `json:"notes,omitempty"`
	State string `json:"state,omitempty"`
}
