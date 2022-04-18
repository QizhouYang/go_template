package service

import (
	"go_template/pkg/dto"
	"go_template/pkg/repository"
)

type SystemSettingService interface {
	ListByTab(tabName string) (dto.SystemSettingResult, error)
}

type systemSettingService struct {
	systemSettingRepo repository.SystemSettingRepository
}

func NewSystemSettingService() SystemSettingService {
	return &systemSettingService{
		systemSettingRepo: repository.NewSystemSettingRepository(),
	}
}

func (s systemSettingService) ListByTab(tabName string) (dto.SystemSettingResult, error) {
	var systemSettingResult dto.SystemSettingResult
	vars := make(map[string]string)
	mos, err := s.systemSettingRepo.ListByTab(tabName)
	if err != nil {
		return systemSettingResult, err
	}
	for _, mo := range mos {
		vars[mo.Key] = mo.Value
	}
	if len(mos) > 0 {
		systemSettingResult.Tab = tabName
	}
	systemSettingResult.Vars = vars
	return systemSettingResult, err
}
