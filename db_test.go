package main

import (
	"fmt"
	"github.com/axieinfinity/susanoo/log/logger/zerolog"
	"github.com/axieinfinity/susanoo/postgresql/gorm"
	"github.com/axieinfinity/susanoo/x/viper"
	zerolog2 "github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	gorm2 "gorm.io/gorm"
	"testing"
)

func initDB() (*gorm2.DB, error) {
	loggerProvider, err := zerolog.NewExportPipeline(
		zerolog.Options.Level(zerolog2.DebugLevel),
		zerolog.Options.OutputPath("/dev/stdout"),
	)
	if err != nil {
		return nil, err
	}

	v, err := viper.NewViperFrom("./")
	if err != nil {
		return nil, err
	}

	return gorm.New(
		gorm.Options.LoadConfig(v.Sub("gorm")),
		gorm.Options.Logger(loggerProvider.Logger("GORM")),
	)
}

func TestWriteToGamesTable(t *testing.T) {
	db, err := initDB()
	assert.NoError(t, err)

	// Write query: add a game
	game := &Game{
		Name:    "Axie Infinity",
		AdminID: 1,
	}
	err = db.Create(game).Error
	assert.NoError(t, err)

	fmt.Printf("Write to games table: %+v\n", game)
}

func TestReadFromGamesTableWithPreload(t *testing.T) {
	db, err := initDB()
	assert.NoError(t, err)

	// Read query: get a game preload with user
	var game Game
	err = db.
		Table("games").
		Preload("Admin").
		Where("name = ?", "Axie Infinity").
		First(&game).Error
	assert.NoError(t, err)
	assert.Equal(t, "Axie Infinity", game.Name)
	assert.Equal(t, 1, game.AdminID)
	assert.NotNil(t, game.Admin)
	assert.Equal(t, "Alice", game.Admin.Name)

	fmt.Printf("Read from games table with preload: %+v, admin:%+v\n", game, game.Admin)
}

func TestWriteToUsersTable(t *testing.T) {
	db, err := initDB()
	assert.NoError(t, err)

	// Write query: add a user
	user := &User{
		Name: "Dave",
	}
	err = db.Create(user).Error
	assert.NoError(t, err)

	fmt.Printf("Write to users table: %+v\n", user)
}

func TestReadFromUsersTable(t *testing.T) {
	db, err := initDB()
	assert.NoError(t, err)

	// Read query: get a user
	var user User
	err = db.First(&user).Error
	assert.NoError(t, err)
	assert.Equal(t, "Alice", user.Name)

	fmt.Printf("Read from users table: %+v\n", user)
}

type User struct {
	ID   int `gorm:"primary_key"`
	Name string
}

type Game struct {
	ID      int `gorm:"primary_key"`
	Name    string
	AdminID int
	Admin   *User
}
