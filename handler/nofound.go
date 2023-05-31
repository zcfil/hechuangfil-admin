package handler

import (
	jwt "hechuangfil-admin/pkg/jwtauth"
	"github.com/gin-gonic/gin"
	"log"
)

func NoFound(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	log.Printf("NoRoute claims: %#v\n", claims)
	c.JSON(404, gin.H{
		"code":    "NOT_FOUND",
		"message": "not found",
	})
}
