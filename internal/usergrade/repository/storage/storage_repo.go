package storage

import (
	"sync"

	"wb-test-task-2022/internal/domain"
)

type Repo struct {
	data *sync.Map
}

func NewUserGradeRepo(data *sync.Map) *Repo {
	return &Repo{data: data}
}

func (r *Repo) GetById(id string) (*domain.UserGrade, error) {
	val, ok := r.data.Load(id)
	if !ok {
		return nil, NotFoundError
	}
	return val.(*domain.UserGrade), nil
}

func (r *Repo) Save(userGrade *domain.UserGrade) {
	actual, ok := r.data.LoadOrStore(userGrade.UserId, userGrade)
	if ok {
		actual.(*domain.UserGrade).Update(userGrade)
	}
}

func (r *Repo) List() []domain.UserGrade {
	var userGrades []domain.UserGrade

	r.data.Range(func(key, value any) bool {
		userGrades = append(userGrades, *value.(*domain.UserGrade))
		return true
	})

	return userGrades
}
