package dtos

// POST /v1/upload/avatar
// Request: multipart/form-data with file field
// No request struct needed - handled by c.FormFile("file")

// UploadAvatarResponse is returned after successful avatar upload
type UploadAvatarResponse struct {
	AvatarURL string `json:"avatar_url"`
}
