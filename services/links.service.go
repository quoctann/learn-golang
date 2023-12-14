package service

import (
	"sampe-go-graph/models"

	"gorm.io/gorm"
)

type LinkService interface {
	GetAll() []models.Link
}

type LinkServiceImpl struct {
	DB *gorm.DB
}

func (srv *LinkServiceImpl) GetAll() []models.Link {
	return []models.Link{}
}
