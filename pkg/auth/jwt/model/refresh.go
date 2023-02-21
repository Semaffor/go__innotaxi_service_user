package model

type RefreshInput struct {
	Token string `json:"token" binding:"required"`
}
