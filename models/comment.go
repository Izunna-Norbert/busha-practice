package models

import (
	"time"

	"github.com/Izunna-Norbert/busha-practice/initializers"
	// "gorm.io/gorm"
)

type Comment struct {
	ID           uint64    `json:"id"`
	Comment      string    `json:"comment"`
	IdentifierID string    `json:"identifier_id"`
	Name         string    `json:"name"`
	ClientIP     string    `json:"client_ip"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CommentStore interface {
	CreateComment(comment *Comment) error
	GetComments(movieID int) ([]Comment, error)
	GetCommentsCount(movieID int) (int, error)
}

type CommentModel struct{}

func (cm *CommentModel) CreateComment(comment *Comment) error {
	db := initializers.DB
	err := db.Create(comment).Error
	if err != nil {
		return err
	}
	return nil
}

func (cm *CommentModel) GetComments(movieID string) ([]Comment, error) {
	db := initializers.DB
	var comments []Comment
	
	//order by created_at
	err := db.Where("identifier_id = ?", movieID).Order("created_at DESC").Find(&comments).Error
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (cm *CommentModel) GetCommentsCount(movieID string) (int, error) {
	db := initializers.DB
	var comments []Comment
	err := db.Where("identifier_id = ?", movieID).Find(&comments).Error
	if err != nil {
		return 0, err
	}
	return len(comments), nil
}
