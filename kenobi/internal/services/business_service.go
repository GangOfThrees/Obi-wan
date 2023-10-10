package services

import (
	"context"

	"github.com/GangOfThrees/Obi-wan/kenobi/internal/models"
	"github.com/GangOfThrees/Obi-wan/kenobi/internal/repository"
)

func CreateBusiness(body models.Business) {
	panic("Not implemented")
}

func CreateUserWithCtx(ctx context.Context, body repository.CreateUserParams) (repository.ObiwanUser, error) {
	q, _ := repository.GetDbQueries()
	return q.CreateUser(ctx, body)
}

func GetUsersWithCtx(ctx context.Context) ([]repository.ObiwanUser, error) {
	q, _ := repository.GetDbQueries()
	return q.ListUsers(ctx)
}
