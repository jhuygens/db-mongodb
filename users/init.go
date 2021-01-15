package users

import (
	mongodb "github.com/jhuygens/db-mongodb"
	"github.com/jhuygens/db/users"
	"go.mongodb.org/mongo-driver/mongo"
)

var client mongo.Client

func init() {
	client = mongodb.NewClient()
	crs := NewUsers()
	users.Register(crs)
}
