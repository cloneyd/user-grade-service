package backup

import (
	"compress/gzip"
	"encoding/csv"
	"io"
	"strconv"

	"wb-test-task-2022/internal/domain"
)

var fields = map[string]int{
	"UserId":        0,
	"PostpaidLimit": 1,
	"Spp":           2,
	"ShippingFee":   3,
	"ReturnFee":     4,
}

func CreateDump(w io.Writer, userGrades []domain.UserGrade) error {
	zipWriter := gzip.NewWriter(w)
	csvWriter := csv.NewWriter(zipWriter)

	//if err := csvWriter.Write(fields); err != nil {
	//	return err
	//}

	for _, userGrade := range userGrades {
		if err := csvWriter.Write(userGrade.Values()); err != nil {
			return err
		}
	}

	csvWriter.Flush()
	if err := zipWriter.Flush(); err != nil {
		return err
	}
	if err := zipWriter.Close(); err != nil {
		return err
	}

	return nil
}

func ReadDump(w io.Reader) ([]domain.UserGrade, error) {
	csvReader := csv.NewReader(w)
	var userGrades []domain.UserGrade

	lines, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, line := range lines {
		postpaidLimit, _ := strconv.Atoi(line[fields["PostpaidLimit"]])
		spp, _ := strconv.Atoi(line[fields["Spp"]])
		shippingFee, _ := strconv.Atoi(line[fields["ShippingFee"]])
		returnFee, _ := strconv.Atoi(line[fields["ReturnFee"]])

		userGrade := domain.UserGrade{
			UserId:        line[fields["UserId"]],
			PostpaidLimit: postpaidLimit,
			Spp:           spp,
			ShippingFee:   shippingFee,
			ReturnFee:     returnFee,
		}

		userGrades = append(userGrades, userGrade)
	}

	return userGrades, nil
}
