package impl

import "go_template/pkg/model"

type GreetRepository interface {
	Get(name string) (model.Test, error)
	List() ([]model.Test, error)
	Save(cluster *model.Test) error
	Delete(name string) error
	Page(num, size int, projectName string) (int, []model.Test, error)
}
