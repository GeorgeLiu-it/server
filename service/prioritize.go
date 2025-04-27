package service

import (
	"server/global"
	"server/model/database"
)

type PrioritizeService struct {
}

func (PrioritizeService *PrioritizeService) PrioritizeList() ([]database.SiebelItems, error) {
	var siebelPrioritize []database.SiebelItems
	if err := global.DB.Find(&siebelPrioritize).Error; err != nil {
		return nil, err
	}
	return siebelPrioritize, nil
}
