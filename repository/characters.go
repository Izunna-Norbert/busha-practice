package repository

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Izunna-Norbert/busha-practice/initializers"
	"github.com/Izunna-Norbert/busha-practice/models"
)

type SwapiCharactersResponse struct {
	Results *[]Characters `json:"results"`
}

type Characters struct {
	Name      string   `json:"name"`
	Height    string   `json:"height"`
	Mass      string   `json:"mass"`
	HairColor string   `json:"hair_color"`
	SkinColor string   `json:"skin_color"`
	EyeColor  string   `json:"eye_color"`
	BirthYear string   `json:"birth_year"`
	Gender    string   `json:"gender"`
	Created   string   `json:"created"`
	Edited    string   `json:"edited"`
}

type CharactersStore interface {
	FetchCharacters() ([]Characters, error)
}

var CharacterModel = new(models.CharacterModel)

func FetchCharacters(page string) (*[]Characters, error) {
	cache := initializers.RedisClient.Get(ctx, "characters:"+page).Val()

	// fetch paginated data from cache
	if cache != "" {
		var characters []Characters
		err := json.Unmarshal([]byte(cache), &characters)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		return &characters, nil
	}

	var characters []Characters
	var response SwapiCharactersResponse
	var url = "https://swapi.dev/api/people/?page=" + page
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	//check response status
	if resp.StatusCode != 200 {
		return nil, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	characters = append(characters, *response.Results...)

	if len(characters) > 0 {
		cache, err := json.Marshal(characters)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		//save to cache
		err = initializers.RedisClient.Set(ctx, "characters:"+page, cache, 24*time.Hour).Err()
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		var data []models.Character
		//save to database
		for _, character := range characters {
			created, err := time.Parse("2006-01-02", character.Created)
			if err != nil {
				created = time.Now()
			}
			edited, err := time.Parse("2006-01-02", character.Edited)
			if err != nil {
				edited = time.Now()
			}
			newHeight, err := strconv.ParseFloat(character.Height, 64)
			if err != nil {
				newHeight = 0
			}
			newMass, err := strconv.ParseFloat(character.Mass, 64)
			if err != nil {
				newMass = 0
			}
			data = append(data, models.Character{
				Name:      character.Name,
				Height:    newHeight,
				Mass:      newMass,
				HairColor: character.HairColor,
				SkinColor: character.SkinColor,
				EyeColor:  character.EyeColor,
				BirthYear: character.BirthYear,
				Gender:    character.Gender,
				Created:   created,
				Edited:    edited,
			})
		}
		CharacterModel.CreateCharacters(data)
	}

	return &characters, nil
}

func ListCharacters(p models.Pagination) (*models.Pagination, error) {
	return CharacterModel.ListCharacters(p)
}

func ListFilteredCharacters(p models.Pagination, filter string, filterBy string) (*models.Pagination, error) {
	return CharacterModel.ListFilteredCharacters(p, filter, filterBy)
}
