package database

import "server/global"

// Snow incident list
type SiebelItems struct {
	global.MODEL
	EID             uint     `json:"employee_id"` // Avaya handle id
	Employee        Employee `json:"-" gorm:"foreignKey:EID"`
	SiebelNumber    string   `json:"siebel_number"`    // Snow incident number
	Severity        uint     `json:"severity"`         // Snow incident severity high, medium ...
	Status          string   `json:"status"`           // Snow incident status pending, working in progress ...
	Type            string   `json:"type"`             // Snow incident CA, Consultation ...
	CustomerSegment string   `json:"customer_segment"` // Snow incident white glove, global support ...
}
