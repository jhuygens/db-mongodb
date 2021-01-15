package users

// User db info
type User struct {
	ID               int64         `json:"_id"`
	Email            string        `json:"email"`
	Password         string        `json:"password"`
	AppName          string        `json:"app_name"`
	ClientID         string        `json:"client_id"`
	CurrentSecretKey string        `json:"current_secret_key"`
	Token            string        `json:"token"`
	RefreshToken     string        `json:"refresh_token"`
	RedirectURLs     []RedirectURL `json:"redirect_urls"`
	CreatedAt        string        `json:"created_at"`
}

// RedirectURL of the user app
type RedirectURL struct {
	ID          string `json:"_id"`
	URL         string `json:"url"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}
