package usecase

import (
	"wb-test-task-2022/internal/domain"
)

type UserGradeUseCase struct {
	repo domain.UserGradeRepository
}

func NewUserGradeUseCase(repo domain.UserGradeRepository) *UserGradeUseCase {
	return &UserGradeUseCase{repo: repo}
}

func (u *UserGradeUseCase) GetById(id string) (*domain.UserGrade, error) {
	return u.repo.GetById(id)
}

func (u *UserGradeUseCase) Save(userGrade *domain.UserGrade) {
	u.repo.Save(userGrade)
}
