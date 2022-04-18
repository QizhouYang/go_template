package repository

import (
	"go_template/pkg/database"
	"go_template/pkg/model"
)

type SystemSettingRepository interface {
	ListByTab(tabName string) ([]model.SystemSetting, error)
}

func NewSystemSettingRepository() SystemSettingRepository {
	return &systemSettingRepository{}
}

type systemSettingRepository struct {
}

func (s systemSettingRepository) ListByTab(tabName string) ([]model.SystemSetting, error) {
	var systemSettings []model.SystemSetting
	err := database.DB.Where("tab = ?", tabName).Find(&systemSettings).Error
	return systemSettings, err
}
