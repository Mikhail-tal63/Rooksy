package auth

import "go.mongodb.org/mongo-driver/mongo"

type AuthRepository struct {
	db *mongo.Collection
}

func (r *AuthRepository) NewAuthRepository(db *mongo.Database) *AuthRepository {
	return &AuthRepository{
		db: db.Collection("user"),
	}
}
