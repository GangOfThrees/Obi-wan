package dtos

import "github.com/GangOfThrees/Obi-wan/kenobi/internal/models"

type CreateBusinessDto struct {
	Name                string `json:"name" validate:"required,min=3"`
	Description         string `json:"description"`
	Address             string `json:"address"`
	City                string `json:"city"`
	State               string `json:"state"`
	Zip                 string `json:"zip"`
	Phone               string `json:"phone" validate:"e164"`
	Website             string `json:"website"`
	Email               string `json:"email" validate:"required,email"`
	Password            string `json:"password"`
	PreferredBotService string `json:"preferredBotService" validate:"required,oneof=chatgpt custom-llama"`
}

func (c *CreateBusinessDto) ToBusiness() *models.Business {
	return &models.Business{
		Name:                c.Name,
		Description:         c.Description,
		Address:             c.Address,
		City:                c.City,
		State:               c.State,
		Zip:                 c.Zip,
		Phone:               c.Phone,
		Website:             c.Website,
		Email:               c.Email,
		PreferredBotService: c.PreferredBotService,
	}
}
