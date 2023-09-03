package migrations

import (
	"github.com/google/uuid"
	"time"
)

type Recipes struct {
	RecipeID                 uuid.UUID  `gorm:"uuid;default:uuid_generate_v4();primary_key;column:recipe_id;" json:"recipe_id"`
	RecipeName               string     `gorm:"type:varchar(255);not null" json:"recipe_name"`
	RecipeDescription        string     `gorm:"type:text;not null" json:"recipe_description"`
	RecipeImage              string     `gorm:"type:text;not null" json:"recipe_image"`
	RecipePreparationTime    string     `gorm:"type:varchar(120);not null" json:"recipe_preparation_time"`
	RecipeCookingTime        string     `gorm:"type:varchar(120);not null" json:"recipe_cooking_time"`
	RecipePortionSuggestions string     `gorm:"type:varchar(120);not null" json:"recipe_portion_suggestions"`
	RecipeRating             string     `gorm:"type:varchar(120);not null" json:"recipe_rating"`
	CreatedAt                time.Time  `gorm:"default:null" json:"created_at"`
	UpdatedAt                *time.Time `gorm:"default:null" json:"updated_at"`
	DeletedAt                *time.Time `gorm:"default:null" json:"deleted_at"`
	CategoryId               uuid.UUID  `gorm:"foreignKey;type:char(36);" json:"category_id"`
	Category                 Categories `gorm:"foreignKey:CategoryID;references:category_id"`
}
