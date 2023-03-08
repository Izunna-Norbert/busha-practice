package forms

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type CommentForm struct{}

type CreateCommentForm struct {
	Comment string `form:"comment" json:"comment" binding:"required,min=1,max=500"`
	Name    string `form:"name" json:"name" binding:"required,min=1,max=50"`
}

func (c *CommentForm) Validate(data []byte) error {
	var comment CreateCommentForm
	err := json.Unmarshal(data, &comment)
	if err != nil {
		return err
	}

	validate := validator.New()
	err = validate.Struct(comment)
	if err != nil {
		return err
	}
	return nil
}

// Comment
func (c *CommentForm) Comment(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		message = "Comment is required"
	case "min":
		message = "Comment must be at least 1 character"
	case "max":
		message = "Comment must be less than 500 characters"
	}
	return
}

// Name
func (c *CommentForm) Name(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		message = "Name is required"
	case "min":
		message = "Name must be at least 1 character"
	case "max":
		message = "Name must be less than 50 characters"
	}
	return
}

//Create
func (c *CommentForm) Create(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:
		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Comment":
				return c.Comment(err.Tag())
			case "Name":
				return c.Name(err.Tag())
			}
		}
	default:
		return "Invalid request"
	}
	return "Something went wrong, please try again later"

}
