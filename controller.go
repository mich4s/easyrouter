package router

import (
	"net/http"

	"github.com/jinzhu/gorm"
)

type Controller struct {
	DB *gorm.DB
}

func (c *Controller) NotFound(r *http.Request) (string, error) {
	return "Route not found", nil
}
