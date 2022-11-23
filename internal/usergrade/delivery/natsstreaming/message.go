package natsstreaming

import (
	"github.com/google/uuid"

	"wb-test-task-2022/internal/domain"
)

type UserGradeMessage struct {
	PublisherUUID uuid.UUID         `json:"publisher_uuid"`
	Payload       *domain.UserGrade `json:"payload"`
}

func NewUserGradeMessage(publisherUUID uuid.UUID, payload *domain.UserGrade) *UserGradeMessage {
	return &UserGradeMessage{PublisherUUID: publisherUUID, Payload: payload}
}
