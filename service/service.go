package service

import (
	"atous/db"

	"github.com/gin-gonic/gin"
)

func New(r *gin.Engine, db *db.DB) {
	initServiceUser(r, db)
}
