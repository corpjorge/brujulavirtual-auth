package repositories

import (
	"brujulavirtual-auth/src/auth/domain/models"
	"brujulavirtual-auth/src/auth/domain/ports"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type Mongo struct {
	collection *mongo.Collection
}

func Auth() ports.Repository {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("brujulavirtual").Collection("users")
	return &Mongo{
		collection: collection,
	}
}

func (r *Mongo) Validate(auth models.Auth) (models.Auth, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"user": auth.User, "password": auth.Password}

	var result models.Auth
	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Default().Println("No documents found for the given filter")
			return models.Auth{}, errors.New("no user found")
		}
		log.Default().Printf("Error finding document: %v\n", err)
		return models.Auth{}, err
	}

	return result, nil
}
