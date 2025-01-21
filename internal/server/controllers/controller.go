package controllers

import "github.com/KoruptTinker/korupt-monitor/internal/models"

type Controller struct {
	Services External
	DB       models.Mongo
}

type External struct{}
