package models

type Content struct {
	ID        int    `gorm:"column:content_id" json:"content_id"`
	Title     string `gorm:"column:title"      json:"title"`
	Summary   string `gorm:"column:summary"    json:"summary"`
	Content   string `gorm:"column:content"    json:"content"`
	CreatedAt string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt string `gorm:"column:updated_at" json:"updated_at"`
	User      int    `gorm:"column:user"       json:"user"`
}

func (Content) TableName() string {
	return "content_content"
}
