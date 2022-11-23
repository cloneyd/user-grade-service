package backup

import (
	"compress/gzip"
	"encoding/csv"
	"io"

	"wb-test-task-2022/internal/domain"
)

var fields = []string{"UserId", "PostpaidLimit", "Spp", "ShippingFee", "ReturnFee"}

func CreateDump(w io.Writer, userGrades []domain.UserGrade) error {
	zipWriter := gzip.NewWriter(w)
	csvWriter := csv.NewWriter(zipWriter)

	if err := csvWriter.Write(fields); err != nil {
		return err
	}

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
