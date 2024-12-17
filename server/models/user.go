package models

type User struct {
	ID           int64   `gorm:"column:id;primaryKey" json:"id"`
	Username     string  `gorm:"column:username;unique;not null" json:"username"`
	Name         string  `gorm:"column:name;not null" json:"name"`
	Password     string  `gorm:"column:password;not null" json:"-"`
	YTApiKey     *string `gorm:"column:yt_api_key" json:"yt_api_key"`
	GeminiApiKey *string `gorm:"column:gemini_api_key" json:"gemini_api_key"`
	OpenAIApiKey *string `gorm:"column:openai_api_key" json:"openai_api_key"`
	ClaudeApiKey *string `gorm:"column:claude_api_key" json:"claude_api_key"`
	Shorts       []Short `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"shorts"`
}
