package main

import "social-media-go/internal/models"

type Handler struct {
	users *models.UserModel
}

func NewHandler(dbModel *models.DBModel) *Handler {
	return &Handler{
		users: &dbModel.User,
	}
}
