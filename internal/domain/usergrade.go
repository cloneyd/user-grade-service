package domain

import (
	"github.com/nats-io/stan.go"
	"strconv"
)

type UserGrade struct {
	UserId        string `json:"user_id" validate:"required"`
	PostpaidLimit int    `json:"postpaid_limit"`
	Spp           int    `json:"spp"`
	ShippingFee   int    `json:"shipping_fee"`
	ReturnFee     int    `json:"return_fee"`
}

type UserGradeRepository interface {
	GetById(id string) (*UserGrade, error)
	Save(userGrade *UserGrade)
	List() []UserGrade
}

type UserGradeUseCase interface {
	GetById(id string) (*UserGrade, error)
	Save(userGrade *UserGrade) error
	List() []UserGrade
}

type UserGradePublisher interface {
	Publish(userGrade *UserGrade) error
}

type UserGradeSubscriber interface {
	Subscribe() (stan.Subscription, error)
}

func (ug *UserGrade) Update(new *UserGrade) {
	if new.PostpaidLimit != 0 {
		ug.PostpaidLimit = new.PostpaidLimit
	}

	if new.Spp != 0 {
		ug.Spp = new.Spp
	}

	if new.ShippingFee != 0 {
		ug.ShippingFee = new.ShippingFee
	}

	if ug.ReturnFee != 0 {
		ug.ReturnFee = new.ReturnFee
	}
}

func (ug *UserGrade) Values() []string {
	return []string{
		ug.UserId,
		strconv.Itoa(ug.PostpaidLimit),
		strconv.Itoa(ug.Spp),
		strconv.Itoa(ug.ShippingFee),
		strconv.Itoa(ug.ReturnFee),
	}
}
