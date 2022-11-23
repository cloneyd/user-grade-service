package natsstreaming

import (
	"wb-test-task-2022/internal/domain"
)

type UserGradeMessage struct {
	PublisherId string            `json:"publisher_id"`
	Payload     *domain.UserGrade `json:"payload"`
}

func NewUserGradeMessage(PublisherId string, payload *domain.UserGrade) *UserGradeMessage {
	return &UserGradeMessage{PublisherId: PublisherId, Payload: payload}
}
