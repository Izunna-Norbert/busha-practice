package repository

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"sort"
	"time"

	"github.com/Izunna-Norbert/busha-practice/initializers"
	"github.com/Izunna-Norbert/busha-practice/models"
)

type Response struct {
	Results *[]Movies `json:"results"`
}

type Movies struct {
	ID            int      `json:"episode_id"`
	Title         string   `json:"title"`
	OpeningCrawl  string   `json:"opening_crawl"`
	Director      string   `json:"director"`
	Producer      string   `json:"producer"`
	ReleaseDate   string   `json:"release_date"`
	CommentsCount int      `json:"comments_count"`
	URL           string   `json:"url"`
	Characters    []string `json:"characters"`
}

type MoviesStore interface {
	FetchMovies() ([]Movies, error)
}

type CommentResponse struct {
	ID        uint64    `json:"id"`
	MovieID   string    `json:"movie_id"`
	Comment   string    `json:"comment"`
	Name      string    `json:"name"`
	ClientIP  string    `json:"client_ip"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var ctx = context.Background()
var commentModel = new(models.CommentModel)

func FetchMovies() (*[]Movies, error) {

	cache := initializers.RedisClient.Get(ctx, "movies").Val()

	if cache != "" {
		var movies []Movies
		err := json.Unmarshal([]byte(cache), &movies)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		sort.Slice(movies, func(i, j int) bool {
			layout := "2006-01-02"
			t1, _ := time.Parse(layout, movies[i].ReleaseDate)
			t2, _ := time.Parse(layout, movies[j].ReleaseDate)
			return t1.Before(t2)
		})
		return &movies, nil
	} else {
		response, err := http.Get("https://swapi.dev/api/films/")
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		var result Response
		err = json.Unmarshal(body, &result)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		movies := *result.Results

		sort.Slice(movies, func(i, j int) bool {
			layout := "2006-01-02"
			t1, _ := time.Parse(layout, movies[i].ReleaseDate)
			t2, _ := time.Parse(layout, movies[j].ReleaseDate)
			return t1.Before(t2)
		})

		// fetch comments count
		for i, movie := range movies {
			count, err := commentModel.GetCommentsCount(path.Base(movie.URL))
			if err != nil {
				log.Println(err.Error())
				return nil, err
			}
			movies[i].CommentsCount = count
		}

		moviesJson, err := json.Marshal(movies)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		err = initializers.RedisClient.Set(ctx, "movies", moviesJson, 24*time.Hour).Err()
		if err != nil {
			log.Println(err.Error())
			// return nil, err
		}

		return &movies, nil
	}
}

func FetchMovie(id string) (*Movies, error) {
	cache := initializers.RedisClient.Get(ctx, "movie:"+id).Val()

	if cache != "" {
		var movie Movies
		err := json.Unmarshal([]byte(cache), &movie)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		return &movie, nil
	} else {
		response, err := http.Get("https://swapi.dev/api/films/" + id)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		var movie Movies
		err = json.Unmarshal(body, &movie)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		movie.CommentsCount, err = commentModel.GetCommentsCount(path.Base(movie.URL))

		movieJson, err := json.Marshal(movie)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		err = initializers.RedisClient.Set(ctx, "movie:"+id, movieJson, 24*time.Hour).Err()
		if err != nil {
			log.Println(err.Error())
			// return nil, err
		}

		return &movie, nil
	}
}

func FetchMovieComments(id string) (*[]models.Comment, error) {
	comments, err := commentModel.GetComments(id)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &comments, nil
}

func CreateMovieComment(comment *models.Comment) (CommentResponse, error) {

	err := commentModel.CreateComment(comment)

	if err != nil {
		log.Println(err.Error())
		return CommentResponse{}, err
	}

	err = initializers.RedisClient.Del(ctx, "movies").Err()
	if err != nil {
		log.Println(err.Error())
	}
	err = initializers.RedisClient.Del(ctx, "movie:"+comment.IdentifierID).Err()
	if err != nil {
		log.Println(err.Error())
	}

	return CommentResponse{
		ID:        comment.ID,
		MovieID:   comment.IdentifierID,
		Comment:   comment.Comment,
		ClientIP:  comment.ClientIP,
		Name:      comment.Name,
		UpdatedAt: comment.UpdatedAt,
		CreatedAt: comment.CreatedAt,
	}, nil
}
