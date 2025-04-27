package database

import "server/global"

// Snow incident list
type SnowItems struct {
	global.MODEL
	EID             uint     `json:"employee_id"` // Avaya handle id
	Employee        Employee `json:"-" gorm:"foreignKey:EID"`
	IncidentNumber  string   `json:"incident_number"`  // Snow incident number
	Priority        uint     `json:"priority"`         // Snow incident priority high, medium ...
	Status          string   `json:"status"`           // Snow incident status pending, working in progress ...
	Type            string   `json:"type"`             // Snow incident CA, Consultation ...
	CustomerSegment string   `json:"customer_segment"` // Snow incident white glove, global support ...
}
