package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	u "gocomments/utils"
)

var validate *validator.Validate

type User struct {
	Uid       uuid.UUID `json:"uid"`
	FirstName string    `json:"first_name"  validate:"required"`
	LastName  string    `json:"last_name"  validate:"required"`
	Image     string    `json:"image"`
}

type Comment struct {
	Uid         uuid.UUID `json:"uid"`
	Body        string    `json:"body" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
	Owner       User      `json:"owner" validate:"required"`
	Rate        uint8     `json:"rate"  validate:"required,gte=0,lte=5"`
	ContentType string    `json:"content_type" validate:"required,max=20"`
}

func (comment *Comment) Validate() (map[string]interface{}, bool) {
	validate = validator.New()
	err := validate.Struct(comment)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		responseBody := map[string]string{"error": validationErrors.Error()}
		resp := u.Message()
		resp["error"] = responseBody

		return resp, false
	}
	resp := u.Message()
	return resp, true
}

func (comment *Comment) Create() map[string] interface{} {
	if resp, ok := comment.Validate(); !ok {
		return resp
	}
	comment.Uid = uuid.New()
	GetDB().Create(comment)
	response := u.Message()
	response["data"] = comment

	return response
}

func GetComments(cType string) []*Comment {
	comments := make([]*Comment, 0)
	GetDB().Where("content_type = ?", cType).Order("created_at desc").Find(&comments)

	return comments
}