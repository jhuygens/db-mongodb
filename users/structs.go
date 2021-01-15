package users

import "go.mongodb.org/mongo-driver/bson/primitive"

// User db info
type User struct {
	ID               primitive.ObjectID `json:"_id,omitempty"`
	Email            string             `json:"email,omitempty"`
	Password         string             `json:"password,omitempty"`
	AppName          string             `json:"app_name,omitempty"`
	ClientID         string             `json:"client_id,omitempty"`
	CurrentSecretKey string             `json:"current_secret_key,omitempty"`
	Token            string             `json:"token,omitempty"`
	RefreshToken     string             `json:"refresh_token,omitempty"`
	RedirectURLs     []RedirectURL      `json:"redirect_urls,omitempty"`
	CreatedAt        string             `json:"created_at,omitempty"`
}

// RedirectURL of the user app
type RedirectURL struct {
	ID          primitive.ObjectID `json:"_id,omitempty"`
	URL         string             `json:"url,omitempty"`
	Description string             `json:"description,omitempty"`
	CreatedAt   string             `json:"created_at,omitempty"`
}
