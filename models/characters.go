package models

import (
	"time"

	"github.com/Izunna-Norbert/busha-practice/initializers"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Character struct {
	gorm.Model
	//unique index for name
	Name      string    `json:"name" gorm:"uniqueIndex"`
	Height    float64    `json:"height"`
	Mass      float64    `json:"mass"`
	HairColor string    `json:"hair_color"`
	SkinColor string    `json:"skin_color"`
	EyeColor  string    `json:"eye_color"`
	BirthYear string    `json:"birth_year"`
	Gender    string    `json:"gender"`
	Created   time.Time `json:"created"`
	Edited    time.Time `json:"edited"`
}

type CharacterModel struct{
}

// sort queries
const (
	ASC  = "name asc"
	DESC = "name desc"
)

//add a new character to the database
func (ch *CharacterModel) CreateCharacter(character *Character) (cha Character, err error) {
	db := initializers.DB
	err = db.Create(&character).Error
	if err != nil {
		return cha, err
	}
	return *character, nil
}

func (ch *CharacterModel) CreateCharacters(characters []Character) (cha []Character, err error) {
	db := initializers.DB
	columns := []string{"name", "height", "mass", "hair_color", "skin_color", "eye_color", "birth_year", "gender", "created", "edited"}
	err = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		DoUpdates: clause.AssignmentColumns(columns),
	}).Create(&characters).Error
	if err != nil {
		return cha, err
	}
	return characters, nil
}

func (ch *CharacterModel) ListCharacters(p Pagination) (*Pagination, error) {
	db := initializers.DB
	var characters []*Character
	if p.Filter != nil {
		db = db.Where(p.Filter)
	}
	db.Scopes(paginate(characters, &p, db)).Find(&characters)
	p.Rows = characters

	return &p, nil
}

func (ch *CharacterModel) ListSearchCharacters(p Pagination, search string) (*Pagination, error) {
	db := initializers.DB
	var characters []*Character
	db.Scopes(paginate(characters, &p, db)).Where("name LIKE ?", "%"+search+"%").Find(&characters)
	p.Rows = characters

	return &p, nil
}

func (ch *CharacterModel) ListFilteredCharacters(p Pagination, filter string, filterBy string) (*Pagination, error) {
	db := initializers.DB

	var where string
	if filter == "gender" {
		where = "gender = ?"
	} 
	if filter == "name" {
		where = "name = ?"
	}

	var characters []*Character


	db.Where(where, filterBy).Find(&characters).Scopes(paginate(characters, &p, db))
	p.Rows = characters

	return &p, nil
}
