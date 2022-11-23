package usecase

import (
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
	if err := u.pub.Publish(userGrade); err != nil {
		return err
	}

	u.repo.Save(userGrade)

	return nil
}

func (u *UserGradeUseCase) List() []domain.UserGrade {
	return u.repo.List()
}
