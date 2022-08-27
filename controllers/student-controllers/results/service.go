package resultsStudent

import (
	model "github.com/j2eevip/gin-restful-example/models"
)

type Service interface {
	ResultsStudentService() (*[]model.EntityStudent, string)
}

type service struct {
	repository Repository
}

func NewServiceResults(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ResultsStudentService() (*[]model.EntityStudent, string) {

	resultCreateStudent, errCreateStudent := s.repository.ResultsStudentRepository()

	return resultCreateStudent, errCreateStudent
}
