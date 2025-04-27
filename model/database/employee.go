package database

import (
	"server/global"
)

// Avaya Employee table
type Employee struct {
	global.MODEL
	Handle string `json:"handle"` // Avaya handle
}
