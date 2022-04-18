package service

import (
	"fmt"
	"go_template/pkg/environment"
	"go_template/pkg/repository"
	repositoryImpl "go_template/pkg/repository/impl"
	"go_template/pkg/service/impl"
	self_exception "go_template/pkg/util/exception"
)

// NewGreetService returns a service backed with a "db" based on "env".
func NewGreetService(env environment.Env) impl.GreetService {
	service := &greetService{prefix: "Hello", greetRepo: repository.NewGreetRepository()}

	switch env {
	case environment.PROD:
		return service
	case environment.DEV:
		return &greeterWithLogging{service}
	default:
		panic("unknown environment")
	}
}

type greetService struct {
	prefix    string
	greetRepo repositoryImpl.GreetRepository
}

func (s *greetService) Say(name string) (string, error) {

	//create_test := model.Test{Name: "create_test"}

	getTest, err := s.greetRepo.Get(name)
	if err != nil {
		return "", err
	}

	/**if err := database.DB.Table("test").First(&test, 1).Error; err != nil {
		return "", err
	}**/

	result := s.prefix + " " + fmt.Sprint(getTest.Id)
	return result, nil
}

func (s *greetService) Delete(name string) error {
	if name == "" && len(name) <= 0 {
		return self_exception.GetErrorInfo("invalid info!")
	}
	return s.greetRepo.Delete(name)
}

type greeterWithLogging struct {
	*greetService
}

func (s *greeterWithLogging) Say(input string) (string, error) {
	result, err := s.greetService.Say(input)
	fmt.Printf("result: %s\nerror: %v\n", result, err)
	return result, err
}
