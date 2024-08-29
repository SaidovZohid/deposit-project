package models

type UserModelResp struct {
	Id          int64   `json:"id"`
	FullName    *string `json:"full_name"`
	Email       string  `json:"email"`
	PhoneNumber *string `json:"phone_number"`
	Balance     float64 `json:"balance"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
}
