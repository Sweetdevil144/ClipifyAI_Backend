package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User Model
type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Username       string             `bson:"username"`
	Name           string             `bson:"name"`
	Password       string             `bson:"password"`
	YTApiKey       string             `bson:"yt_api_key"`
	GeminiApiKey   string             `bson:"gemini_api_key"`
	OpenAIApiKey   string             `bson:"openai_api_key"`
	ClaudeApiKey   string             `bson:"claude_api_key"`
}