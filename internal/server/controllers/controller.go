package controllers

import (
	"github.com/KoruptTinker/korupt-monitor/internal/models"
	"github.com/KoruptTinker/korupt-monitor/internal/services"
)

type Controller struct {
	Services services.External
	DB       models.Mongo
}
