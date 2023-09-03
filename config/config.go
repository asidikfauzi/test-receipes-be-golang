package config

import (
	"fmt"
	"github.com/asidikfauzi/test-recipes-be-golang/config/migrations"
	"github.com/asidikfauzi/test-recipes-be-golang/repository/domain"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	db               *gorm.DB
	categorySeeder   domain.CategoryMigration
	ingredientSeeder domain.IngredientMigration
	recipeSeeder     domain.RecipeMigration
}

func NewConfig(
	conn *gorm.DB,
	cs domain.CategoryMigration,
	is domain.IngredientMigration,
	rs domain.RecipeMigration) domain.Config {
	return &Config{
		db:               conn,
		categorySeeder:   cs,
		ingredientSeeder: is,
		recipeSeeder:     rs,
	}
}

func GetEnv(key string) string {
	var appConfig map[string]string
	appConfig, err := godotenv.Read()

	if err != nil {
		fmt.Println("Error reading .env file")
	}

	return appConfig[key]
}

func Open() (*gorm.DB, error) {
	postgresCredentials := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai",
		GetEnv("DB_HOST"),
		GetEnv("DB_USER"),
		GetEnv("DB_PASSWORD"),
		GetEnv("DB_NAME"),
		GetEnv("DB_PORT"),
		GetEnv("SSL_MODE"),
	)

	conn, err := gorm.Open(postgres.Open(postgresCredentials), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	return conn, nil
}

func (c *Config) InitDB() (string, error) {
	conn, _ := Open()

	c.db = conn
	fmt.Println("Database connection successful")

	msg := "Connection database successfully"
	c.InitMigrate()
	c.InitSeeder()

	return msg, nil

}

func (c *Config) InitMigrate() interface{} {
	c.db.AutoMigrate(&migrations.LogActivity{})
	c.db.AutoMigrate(&migrations.Categories{})
	c.db.AutoMigrate(&migrations.Ingredients{})
	c.db.AutoMigrate(&migrations.Recipes{})
	c.db.AutoMigrate(&migrations.RecipesToIngredients{})

	message := "SUCCESSFULLY ADD ALL TABLE MIGRATION"
	fmt.Println(message)

	return message
}

func (c *Config) InitSeeder() interface{} {
	c.categorySeeder.UpCategorySeeder()
	c.ingredientSeeder.UpIngredientSeeder()
	c.recipeSeeder.UpRecipeSeeder()

	message := "SUCCESSFULLY ADD ALL TABLE SEEDER"
	fmt.Println(message)

	return message
}
