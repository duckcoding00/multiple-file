package handler

import (
	"fmt"
	"net/http"

	"github.com/duckcoding00/multiple-file/internal/service"
	"github.com/duckcoding00/multiple-file/lib/utils"
)

type FileHandler struct {
	s service.Service
}

func (h *FileHandler) Upload(w http.ResponseWriter, r *http.Request) {
	const maxSize = 2 << 20 // 2MB
	r.Body = http.MaxBytesReader(w, r.Body, maxSize)

	// Parse the multipart form
	if err := r.ParseMultipartForm(maxSize); err != nil {
		if err.Error() == "http: request body too large" {
			utils.WriteErr(w, http.StatusBadRequest, fmt.Errorf("file too large, max size is %d MB", maxSize/(1<<20)))
			return
		}
		utils.WriteErr(w, http.StatusBadRequest, fmt.Errorf("failed to parse request: %w", err))
		return
	}

	files := r.MultipartForm.File["files"]
	if len(files) == 0 {
		utils.WriteErr(w, http.StatusBadRequest, fmt.Errorf("please upload at least one file"))
		return
	}

	if err := h.s.File.Upload(files); err != nil {
		utils.WriteErr(w, http.StatusInternalServerError, fmt.Errorf("upload failed: %w", err))
		return
	}

	utils.WriteOk(w, http.StatusCreated, "Files successfully uploaded")
}
