package dtos

// QuestionDto is a struct that represents the data transfer object for questions.
type QuestionDto struct {
	Question string `json:"question" validate:"required,min=5,max=1000"`
}
