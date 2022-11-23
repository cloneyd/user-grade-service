package v1

import (
	"compress/gzip"
	"encoding/csv"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"wb-test-task-2022/internal/domain"
	"wb-test-task-2022/internal/usergrade/repository/storage"
)

type UserGradeHandlers struct {
	userGradeUseCase domain.UserGradeUseCase
}

func NewUserGradeHandlers(userGradeUseCase domain.UserGradeUseCase) *UserGradeHandlers {
	return &UserGradeHandlers{userGradeUseCase: userGradeUseCase}
}

func (h *UserGradeHandlers) Get(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")

	userGrade, err := h.userGradeUseCase.GetById(userId)
	if err != nil {
		if errors.Is(err, storage.NotFoundError) {
			errorAsJSON(w, http.StatusNotFound, err)
			return
		}

		errorAsJSON(w, http.StatusInternalServerError, err)
		return
	}

	err = json.NewEncoder(w).Encode(userGrade)
	if err != nil {
		errorAsJSON(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func (h *UserGradeHandlers) Set(w http.ResponseWriter, r *http.Request) {
	var userGrade domain.UserGrade

	if err := json.NewDecoder(r.Body).Decode(&userGrade); err != nil {
		errorAsJSON(w, http.StatusBadRequest, err)
		return
	}

	validate := validator.New()
	if err := validate.Struct(userGrade); err != nil {
		errorAsJSON(w, http.StatusBadRequest, err)
		return
	}

	if err := h.userGradeUseCase.Save(&userGrade); err != nil {
		errorAsJSON(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *UserGradeHandlers) Backup(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Disposition", "attachment; filename=backup.csv.gz")
	w.Header().Set("Content-Type", "application/gzip")

	zipWriter := gzip.NewWriter(w)
	csvWriter := csv.NewWriter(zipWriter)

	userGrades := h.userGradeUseCase.List()
	for _, userGrade := range userGrades {
		err := csvWriter.Write(userGrade.Fields())
		if err != nil {
			errorAsJSON(w, http.StatusInternalServerError, err)
			return
		}
	}

	csvWriter.Flush()
	if err := zipWriter.Flush(); err != nil {
		log.Println(err)
	}
	if err := zipWriter.Close(); err != nil {
		log.Println(err)
	}
}
