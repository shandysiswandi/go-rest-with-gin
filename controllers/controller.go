package controllers

import "github.com/gin-gonic/gin"

// Controller is base interface for inject
type Controller interface {
	Inject(e *gin.Engine)
}
