package v1

import (
	"blog-sevice/pkg/app"
	"blog-sevice/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

func (a Article) Create(c *gin.Context) {}
func (a Article) List(c *gin.Context)   {}
func (a Article) Update(c *gin.Context) {}
func (a Article) Delete(c *gin.Context) {}
func (a Article) GET(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}
