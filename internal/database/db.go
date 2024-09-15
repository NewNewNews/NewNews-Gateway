package database

import (
	"context"

	"github.com/NewNewNews/NewNews-Gateway/internal/models"
)

type Database struct {
	client *PrismaClient
}

func New(databaseURL string) (*Database, error) {
	client := NewClient()
	if err := client.Connect(); err != nil {
		return nil, err
	}

	return &Database{client: client}, nil
}

func (db *Database) Disconnect() error {
	return db.client.Disconnect()
}

func (db *Database) CreateUser(ctx context.Context, user *models.User) error {
	_, err := db.client.User.CreateOne(
		User.Email.Set(user.Email),
		User.HashedPassword.Set(user.HashedPassword),
		User.IsAdmin.Set(user.IsAdmin),
	).Exec(ctx)
	return err
}

func (db *Database) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := db.client.User.FindUnique(
		User.Email.Equals(email),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &models.User{
		ID:             user.ID,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		IsAdmin:        user.IsAdmin,
	}, nil
}

func (db *Database) CreateLog(ctx context.Context, log *models.Log) error {
	_, err := db.client.Log.CreateOne(
		Log.UserID.Set(log.UserID),
		Log.Action.Set(log.Action),
		Log.Timestamp.Set(log.Timestamp),
	).Exec(ctx)
	return err
}
