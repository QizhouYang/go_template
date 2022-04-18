package dto

type SystemSettingResult struct {
	Vars map[string]string `json:"vars" validate:"required"`
	Tab  string            `json:"tab" validate:"required"`
}
