package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
	"wb-test-task-2022/internal/backup"
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
	w.Header().Set("Content-Type", "application/json")

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
	filename := backup.GenerateBackupFilePath("csv.gz", time.Now())
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	w.Header().Set("Content-Type", "application/gzip")

	userGrades := h.userGradeUseCase.List()

	if err := backup.CreateDump(w, userGrades); err != nil {
		errorAsJSON(w, http.StatusInternalServerError, err)
		return
	}
}
