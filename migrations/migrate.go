package main

import (
	"strconv"

	"github.com/Izunna-Norbert/busha-practice/initializers"
	"github.com/Izunna-Norbert/busha-practice/models"
)

func init() {
	initializers.LoadEnvironmentVariables()
	initializers.ConnectDb()
}

func convertCharacterColumnMassFromStringToFloat() {
	var characters []models.Character
	initializers.DB.Find(&characters)

	for _, character := range characters {
		newMass := strconv.FormatFloat(character.Mass, 'f', 2, 64)
		// convert string to float
		newMassFloat, _ := strconv.ParseFloat(newMass, 64)

		initializers.DB.Model(&character).Update("mass", newMassFloat)
	}
}

func dropAllTables() {
	initializers.DB.Migrator().DropTable(&models.Comment{})
}

func main() {
	dropAllTables()
	initializers.DB.AutoMigrate(&models.Comment{}, &models.Character{})
}
