package usecase

import (
	"encoding/json"
	"wb-test-task-2022/internal/domain"
)

type UserGradeUseCase struct {
	pub  domain.UserGradePublisher
	repo domain.UserGradeRepository
}

func NewUserGradeUseCase(pub domain.UserGradePublisher, repo domain.UserGradeRepository) *UserGradeUseCase {
	return &UserGradeUseCase{pub: pub, repo: repo}
}

func (u *UserGradeUseCase) GetById(id string) (*domain.UserGrade, error) {
	return u.repo.GetById(id)
}

func (u *UserGradeUseCase) Save(userGrade *domain.UserGrade) error {
	body, err := json.Marshal(userGrade)
	if err != nil {
		return err
	}
	if err = u.pub.Publish(body); err != nil {
		return err
	}
	u.repo.Save(userGrade)

	return nil
}

func (u *UserGradeUseCase) List() []domain.UserGrade {
	return u.repo.List()
}
