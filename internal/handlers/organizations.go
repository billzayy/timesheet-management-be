package handlers

import (
	"github.com/billzayy/timesheet-management-be/internal/services"
)

type OrganizeHandler struct {
	service services.OrganizeService
}

func NewOrganizeHandler(s services.OrganizeService) *OrganizeHandler {
	return &OrganizeHandler{s}
}
