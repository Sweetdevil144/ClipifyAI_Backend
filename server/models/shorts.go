package models

type Short struct {
	ID             int64   `gorm:"column:id;primaryKey" json:"id"`
	UserID         int64   `gorm:"column:user_id;not null" json:"user_id"`
	Src            string  `gorm:"column:src;not null" json:"src"`
	StartTime      float64 `gorm:"column:start_time;not null" json:"start_time"`
	EndTime        float64 `gorm:"column:end_time;not null" json:"end_time"`
	Description    *string `gorm:"column:description" json:"description"`
	QualityRanking *int    `gorm:"column:quality_ranking" json:"quality_ranking"`
	User           User    `gorm:"foreignKey:UserID" json:"-"`
}
