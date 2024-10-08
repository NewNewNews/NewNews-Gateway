package database

import (
	"context"
	"fmt"

	"github.com/NewNewNews/NewNews-Gateway/internal/models"
	"github.com/NewNewNews/NewNews-Gateway/prisma/db"
)

type Database struct {
	client *db.PrismaClient
}

func New(databaseURL string) (*Database, error) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
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

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (d *Database) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := d.client.User.FindUnique(
		db.User.Email.Equals(email),
	).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email %s: %w", email, err)
	}
	return &models.User{
		ID:             user.ID,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		IsAdmin:        user.IsAdmin,
	}, nil
}

func (d *Database) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	user, err := d.client.User.FindUnique(
		db.User.ID.Equals(id),
	).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email %s: %w", id, err)
	}
	return &models.User{
		ID:             user.ID,
		Name:           user.Name,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		IsAdmin:        user.IsAdmin,
	}, nil
}

func (d *Database) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	users, err := d.client.User.FindMany().Exec(ctx)
	if err != nil {
		return nil, err
	}

	var allUsers []*models.User
	for _, user := range users {
		allUsers = append(allUsers, &models.User{
			ID:             user.ID,
			Email:          user.Email,
			HashedPassword: user.HashedPassword,
			IsAdmin:        user.IsAdmin,
		})
	}
	return allUsers, nil
}

func (d *Database) UpdateUserByEmail(ctx context.Context, email string, updatedUser *models.User) error {
	_, err := d.client.User.FindUnique(
		db.User.Email.Equals(email),
	).Update(
		db.User.Email.Set(updatedUser.Email),
		db.User.HashedPassword.Set(updatedUser.HashedPassword),
		db.User.Name.Set(updatedUser.Name),
		db.User.IsAdmin.Set(updatedUser.IsAdmin),
	).Exec(ctx)

	return err
}

func (d *Database) DeleteUser(ctx context.Context, email string) error {
	_, err := d.client.User.FindUnique(
		db.User.Email.Equals(email),
	).Delete().Exec(ctx)

	return err
}

func (d *Database) CreateLog(ctx context.Context, log *models.Log) error {
	_, err := d.client.Log.CreateOne(
		db.Log.User.Link(db.User.ID.Equals(log.UserID)),
		db.Log.Action.Set(log.Action),
		db.Log.Timestamp.Set(log.Timestamp),
	).Exec(ctx)

	if err != nil {
		return fmt.Errorf("failed to create log for userID %s: %w", log.UserID, err)
	}

	return nil
}
