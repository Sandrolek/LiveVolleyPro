package dto

type CreateActionRateDTO struct {
	HelpText  string `json:"help_text" binding:"required"`
	Signature string `json:"signature" binding:"required"`
	ActionID  int    `json:"action_id" binding:"required"`
}

type UpdateActionRateDTO struct {
	HelpText  *string `json:"help_text,omitempty"`
	Signature *string `json:"signature,omitempty"`
	ActionID  *int    `json:"action_id,omitempty"`
}
