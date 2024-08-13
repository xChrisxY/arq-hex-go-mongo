package domain

import (
	"time"
)

type Player struct {

	FirstName string `json:"first_name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	Age int `json:"age" binding:"required"`
	CreationTime time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

}

