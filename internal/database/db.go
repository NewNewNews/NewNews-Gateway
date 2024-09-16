package database

import (
	"context"

	"github.com/NewNewNews/NewNews-Gateway/internal/models"
	"github.com/NewNewNews/NewNews-Gateway/prisma/db"
)

type Database struct {
	client *db.PrismaClient
}

func New(databaseURL string) (*Database, error) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return nil, err
	}

	return &Database{client: client}, nil
}

func (d *Database) Disconnect() error {
	return d.client.Prisma.Disconnect()
}

func (d *Database) CreateUser(ctx context.Context, user *models.User) error {
	_, err := d.client.User.CreateOne(
		db.User.Email.Set(user.Email),
		db.User.HashedPassword.Set(user.HashedPassword),
		db.User.Name.Set(user.Name),
		db.User.IsAdmin.Set(user.IsAdmin),
	).Exec(ctx)
	return err
}

func (d *Database) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := d.client.User.FindUnique(
		db.User.Email.Equals(email),
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

func (d *Database) CreateLog(ctx context.Context, log *models.Log) error {
	_, err := d.client.Log.CreateOne(
		db.Log.User.Link(db.User.ID.Equals(log.UserID)),
		db.Log.Action.Set(log.Action),
		db.Log.Timestamp.Set(log.Timestamp),
	).Exec(ctx)
	return err
}
