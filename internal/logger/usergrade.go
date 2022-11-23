package logger

import (
	"log"

	"wb-test-task-2022/internal/domain"
)

func LogUserGrade(userGrade domain.UserGrade) {
	log.Printf(`
UserId: %s
PostpaidLimit: %d
Spp: %d
ShippingFee: %d
ReturnFee: %d`,
		userGrade.UserId,
		userGrade.PostpaidLimit,
		userGrade.Spp,
		userGrade.ShippingFee,
		userGrade.ReturnFee,
	)
}
