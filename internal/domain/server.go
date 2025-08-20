package domain

type Server struct {
	ID           uint64  `gorm:"primaryKey" json:"id"`
	UserID       uint64  `gorm:"type:bigint;not null;index" json:"user_id"`
	ProjectID    *uint64 `gorm:"type:bigint;default:null;index" json:"project_id"`
	GameID       uint64  `gorm:"type:bigint;not null;index" json:"game_id"`
	Port         int     `gorm:"type:int;not null" json:"port"`
	QueryPort    *int    `gorm:"type:int;default:null" json:"query_port"`
	Status       string  `gorm:"type:varchar(20);default:default;not null;index" json:"status"`
	Name         *string `gorm:"type:varchar(150);default:null" json:"name"`
	Region       *string `gorm:"type:varchar(4);default:null" json:"region"`
	Color        uint8   `gorm:"type:int;default:0" json:"color"`
	Tags         *string `gorm:"type:json;default:null" json:"tags"`
	IP           string  `gorm:"type:varchar(40);not null" json:"ip"`
	Likes        uint    `gorm:"->" json:"likes"`
	Dislikes     uint    `gorm:"->" json:"dislikes"`
	ModeID       uint64  `gorm:"type:int;not null" json:"mode_id"`
	ModeName     string  `gorm:"->" json:"mode_name"`
	AggressiveAd bool    `gorm:"->" json:"aggressive_ad"`
	UpdatedAt    int64   `json:"updated_at"`
	CreatedAt    int64   `json:"created_at"`
}
