package kolog

import (
	"go_template/pkg/dto"
	"go_template/pkg/service"
	"go_template/pkg/util/logger"
)

func Save(name, operation, operationInfo string) {
	lS := service.NewSystemLogService()
	logInfo := dto.SystemLogCreate{
		Name:          name,
		Operation:     operation,
		OperationInfo: operationInfo,
	}
	if err := lS.Create(logInfo); err != nil {
		logger.Log.Errorf("save system logs failed, error: %s", err.Error())
	}
}
