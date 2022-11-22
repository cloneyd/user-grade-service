package domain

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
}

type UserGradeUseCase interface {
	GetById(id string) (*UserGrade, error)
	Save(userGrade *UserGrade)
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
