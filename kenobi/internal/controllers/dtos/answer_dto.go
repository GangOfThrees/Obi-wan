package dtos

import "github.com/GangOfThrees/Obi-wan/internal/models"

// AnswerDto is a struct that represents the data transfer object for answers.
type AnswerDto struct {
	Answer string `json:"answer"`
}

// ToAnswerDto converts a models.Answer to an AnswerDto.
func ToAnswerDto(answer models.BotAnswer) AnswerDto {
	return AnswerDto{
		Answer: answer.Answer,
	}
}
