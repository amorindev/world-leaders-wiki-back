package handler

import (
	"context"
	"io"
	"net/http"

	sharedC "github.com/amorindev/go-tmpl/pkg/shared/api/core"
	sharedD "github.com/amorindev/go-tmpl/pkg/shared/domain"
)

// UploadAvatar handles uploading an avatar or banner image for a leader.
// It validates parameters, checks the file type, and delegates the upload to the service layer.
func (h Handler) UploadAvatar(w http.ResponseWriter, r *http.Request) {
	leaderID := r.PathValue("leaderId")
	imgType := r.PathValue("type") // avatar || banner

	if leaderID == "" || imgType == "" {
		sharedC.RespondError(w, sharedD.NewAppError(sharedD.ErrCodeInvalidParams, "missing parameters"))
		return
	}

	allowedTypes := map[string]bool{
		"avatar": true,
		"banner": true,
	}

	if !allowedTypes[imgType] {
		sharedC.RespondError(w, sharedD.NewAppError(sharedD.ErrCodeInvalidParams, "invalid image type"))
		return
	}

	err := r.ParseMultipartForm(20 << 20) // 20MB
	if err != nil {
		sharedC.RespondError(w, sharedD.NewAppError(sharedD.ErrCodeInvalidParams, "invalid form"))
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		sharedC.RespondError(w, sharedD.NewAppError(sharedD.ErrCodeInvalidParams, "error retrieving file"))
		return
	}

	defer file.Close()

	// Proper Content-Type detection
	fileBuffer := make([]byte, 512)
	_, err = file.Read(fileBuffer)
	if err != nil {
		msg := "failed to read uploaded file for content-type detection"
		sharedC.RespondError(w, sharedD.NewAppError(sharedD.ErrCodeInternalServerError, msg))
		return
	}
	fileType := http.DetectContentType(fileBuffer)

	// Reset file reader position
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		msg := "failed to reset file pointer after reading"
		sharedC.RespondError(w, sharedD.NewAppError(sharedD.ErrCodeInternalServerError, msg))
		return
	}

	allowedMime := map[string]bool{
		"image/jpeg":    true,
		"image/png":     true,
		"image/gif":     true,
		"image/webp":    true,
		"image/avif":    true,
		"image/svg+xml": true,
	}

	if !allowedMime[fileType] {
		sharedC.RespondError(w, sharedD.NewAppError(sharedD.ErrCodeInternalServerError, "invalid image format"))
		return
	}

	err = h.LeaderSrv.UploadLeaderImage(context.Background(), leaderID, imgType, header.Filename, file, fileType)
	if err != nil {
		sharedC.RespondError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
