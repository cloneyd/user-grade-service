package storage

import (
	"sync"

	"wb-test-task-2022/internal/domain"
)

type UserGradeRepository struct {
	data *sync.Map
}

func NewUserGradeRepo(data *sync.Map) *UserGradeRepository {
	return &UserGradeRepository{data: data}
}

func (repo *UserGradeRepository) GetById(id string) (*domain.UserGrade, error) {
	val, ok := repo.data.Load(id)
	if !ok {
		return nil, NotFoundError
	}
	return val.(*domain.UserGrade), nil
}

func (repo *UserGradeRepository) Save(userGrade *domain.UserGrade) {
	actual, ok := repo.data.LoadOrStore(userGrade.UserId, userGrade)
	if ok {
		actual.(*domain.UserGrade).Update(userGrade)
	}
}

func (repo *UserGradeRepository) List() []domain.UserGrade {
	var userGrades []domain.UserGrade

	repo.data.Range(func(key, value any) bool {
		userGrades = append(userGrades, *value.(*domain.UserGrade))
		return true
	})

	return userGrades
}
