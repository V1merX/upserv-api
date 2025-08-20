package domain

type User struct {
	ID           uint64  `json:"id"`
	Login        string  `json:"login"`
	Avatar       string  `json:"avatar"`
	SteamID64    string  `json:"steamid64"`
	Balance      float64 `json:"balance"`
	Status       string  `json:"status"`
	APIKey       *string `json:"api_key"`
	IP           *string `json:"ip"`
	ReferralCode *string `json:"referral_code"`
	UpdatedAt    int64   `json:"updated_at"`
	CreatedAt    int64   `json:"created_at"`
}
