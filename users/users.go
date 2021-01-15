package users

import (
	"context"
	"fmt"
	"time"

	"github.com/jgolang/config"
	"github.com/jhuygens/db/users"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Users implements
type Users struct{}

// Save doc ...
func (u Users) Save(usr users.User) error {
	user := User{
		Email:            usr.Email,
		Password:         usr.Password,
		AppName:          usr.AppName,
		ClientID:         usr.ClientID,
		CurrentSecretKey: usr.CurrentSecretKey,
		Token:            usr.Token,
		RefreshToken:     usr.RefreshToken,
		CreatedAt:        time.Now().Format(time.RFC1123Z),
	}
	for _, url := range usr.RedirectURLs {
		user.RedirectURLs = append(
			user.RedirectURLs,
			RedirectURL{
				URL:         url.URL,
				Description: url.Description,
				CreatedAt:   time.Now().Format(time.RFC1123Z),
			},
		)
	}
	collection := client.Database(config.GetString("database.name")).Collection(config.GetString("database.collections.users"))
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	fmt.Println("Inserted user with ID: ", insertResult.InsertedID)
	return nil
}

// Get doc ..
func (u Users) Get(userID int64) (*users.User, error) {
	return nil, nil
}

// GetAll doc ...
func (u Users) GetAll() ([]users.User, error) {
	return nil, nil
}

// GetRedirectURLs ...
func (u Users) GetRedirectURLs(userID int64) ([]users.RedirectURL, error) {
	return nil, nil
}

// Delete doc
func (u Users) Delete(userID int64) error {
	return nil
}

// UpdateCurrentSecretKey doc
func (u Users) UpdateCurrentSecretKey(clientID, secretKey string) error {
	collection := client.Database(config.GetString("database.name")).Collection(config.GetString("database.collections.users"))
	result, err := collection.UpdateOne(
		context.Background(),
		bson.D{primitive.E{Key: "client_id", Value: clientID}},
		bson.D{primitive.E{Key: "current_secret_key", Value: secretKey}},
	)
	if err != nil {
		return err
	}
	fmt.Printf("Updated %v user!\n", result.ModifiedCount)
	return nil
}

// GetByEmail doc ...
func (u Users) GetByEmail(email string) (*users.User, error) {
	collection := client.Database(config.GetString("database.name")).Collection(config.GetString("database.collections.users"))
	filter := bson.D{primitive.E{Key: "email", Value: email}}
	var usr User
	err := collection.FindOne(context.TODO(), filter).Decode(&usr)
	if err != nil {
		return nil, err
	}
	user := users.User{
		Email:            usr.Email,
		Password:         usr.Password,
		AppName:          usr.AppName,
		ClientID:         usr.ClientID,
		CurrentSecretKey: usr.CurrentSecretKey,
		Token:            usr.Token,
		RefreshToken:     usr.RefreshToken,
		CreatedAt:        usr.CreatedAt,
	}
	for _, url := range usr.RedirectURLs {
		user.RedirectURLs = append(
			user.RedirectURLs,
			users.RedirectURL{
				URL:         url.URL,
				Description: url.Description,
				CreatedAt:   url.CreatedAt,
			},
		)
	}
	return &user, nil
}

// UpdateToken doc ...
func (u Users) UpdateToken(clientID, token, refreshToken string) error {
	collection := client.Database(config.GetString("database.name")).Collection(config.GetString("database.collections.users"))
	result, err := collection.UpdateOne(
		context.Background(),
		bson.D{primitive.E{Key: "client_id", Value: clientID}},
		bson.D{primitive.E{Key: "token", Value: token}, primitive.E{Key: "refresh_token", Value: refreshToken}},
	)
	if err != nil {
		return err
	}
	fmt.Printf("Updated %v user!\n", result.ModifiedCount)
	return nil
}

// GetByClientID doc ...
func (u Users) GetByClientID(clientID string) (*users.User, error) {
	collection := client.Database(config.GetString("database.name")).Collection(config.GetString("database.collections.users"))
	filter := bson.D{primitive.E{Key: "client_id", Value: clientID}}
	var usr User
	err := collection.FindOne(context.TODO(), filter).Decode(&usr)
	if err != nil {
		return nil, err
	}
	user := users.User{
		Email:            usr.Email,
		Password:         usr.Password,
		AppName:          usr.AppName,
		ClientID:         usr.ClientID,
		CurrentSecretKey: usr.CurrentSecretKey,
		Token:            usr.Token,
		RefreshToken:     usr.RefreshToken,
		CreatedAt:        usr.CreatedAt,
	}
	for _, url := range usr.RedirectURLs {
		user.RedirectURLs = append(
			user.RedirectURLs,
			users.RedirectURL{
				URL:         url.URL,
				Description: url.Description,
				CreatedAt:   url.CreatedAt,
			},
		)
	}
	return &user, nil
}
