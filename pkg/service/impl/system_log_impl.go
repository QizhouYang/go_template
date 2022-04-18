package impl

import (
	"go_template/pkg/controller/condition"
	"go_template/pkg/controller/page"
	"go_template/pkg/dto"
)

type SystemLogService interface {
	Create(creation dto.SystemLogCreate) error
	Page(num, size int, conditions condition.Conditions) (*page.Page, error)
}
