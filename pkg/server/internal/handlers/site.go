package handlers

import (
	"context"
	"github.com/Takasakiii/ayanami/pkg/server/internal/templates"
	"github.com/gin-gonic/gin"
)

func IndexPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		background := context.Background()
		page := templates.IndexPage()
		_ = page.Render(background, c.Writer)
	}
}
