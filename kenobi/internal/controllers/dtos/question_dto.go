package dtos

// QuestionDto is a struct that represents the data transfer object for questions.
type QuestionDto struct {
	Question string `json:"question" validate:"required,max=1000"`
}
