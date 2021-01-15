package users

import "go.mongodb.org/mongo-driver/bson/primitive"

// User db info
type User struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	Email            string             `bson:"email,omitempty"`
	Password         string             `bson:"password,omitempty"`
	AppName          string             `bson:"app_name,omitempty"`
	ClientID         string             `bson:"client_id,omitempty"`
	CurrentSecretKey string             `bson:"current_secret_key,omitempty"`
	Token            string             `bson:"token,omitempty"`
	RefreshToken     string             `bson:"refresh_token,omitempty"`
	RedirectURLs     []RedirectURL      `bson:"redirect_urls,omitempty"`
	CreatedAt        string             `bson:"created_at,omitempty"`
}

// RedirectURL of the user app
type RedirectURL struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	URL         string             `bson:"url,omitempty"`
	Description string             `bson:"description,omitempty"`
	CreatedAt   string             `bson:"created_at,omitempty"`
}
