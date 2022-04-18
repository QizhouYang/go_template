package repository

import (
	"go_template/pkg/database"
	"go_template/pkg/model"
	"go_template/pkg/repository/impl"
)

func NewGreetRepository() impl.GreetRepository {
	return &greetRepository{}
}

type greetRepository struct {
}

func (g greetRepository) Get(name string) (model.Test, error) {
	var test model.Test

	err := database.DB.Where("name = ?", name).First(&test).Error

	return test, err
}

func (g greetRepository) List() ([]model.Test, error) {
	var tests []model.Test
	result := database.DB.Find(&tests)
	// 行数
	//result.RowsAffected
	return tests, result.Error
}

func (g greetRepository) Save(test *model.Test) error {

	if database.DB.NewRecord(test) {
		err := database.DB.Create(test).Error
		return err
	} else {
		err := database.DB.Create(&test).Error
		return err
	}
}

func (g greetRepository) Delete(name string) error {

	err := database.DB.Where("name = ?", name).Delete(&model.Test{}).Error

	return err
}

func (g greetRepository) Page(num, size int, projectName string) (int, []model.Test, error) {
	return 0, []model.Test{}, nil
}
