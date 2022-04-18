package service

import (
	"go_template/pkg/controller/condition"
	"go_template/pkg/controller/page"
	"go_template/pkg/database"
	"go_template/pkg/dto"
	"go_template/pkg/model"
	"go_template/pkg/service/impl"
)

type systemLogService struct{}

func NewSystemLogService() impl.SystemLogService {
	return &systemLogService{}
}

func (s systemLogService) Create(creation dto.SystemLogCreate) error {
	log := model.SystemLog{
		Name:          creation.Name,
		Operation:     creation.Operation,
		OperationInfo: creation.OperationInfo,
	}

	if database.DB.NewRecord(log) {
		return database.DB.Create(&log).Error
	} else {
		return database.DB.Save(&log).Error
	}
}

func (u systemLogService) Page(num, size int, conditions condition.Conditions) (*page.Page, error) {
	var (
		p         page.Page
		logOfDTOs []dto.SystemLog
		mos       []model.SystemLog
	)
	d := database.DB.Model(model.SystemLog{})
	/**if err := dbUtil.WithConditions(&d, model.SystemLog{}, conditions); err != nil {
		return nil, err
	}*/
	if err := d.
		Count(&p.Total).
		Order("created_at DESC").
		Offset((num - 1) * size).
		Limit(size).
		Find(&mos).Error; err != nil {
		return nil, err
	}

	for _, mo := range mos {
		logOfDTOs = append(logOfDTOs, dto.SystemLog{SystemLog: mo})
	}
	p.Items = logOfDTOs
	return &p, nil
}
