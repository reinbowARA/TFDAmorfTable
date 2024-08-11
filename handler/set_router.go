package handler

import (
	"text/template"

	"github.com/gin-gonic/gin"
)

func (s *Server) SetRouter(router *gin.Engine) {
	router.Static("/static", "./web/static")
	
	router.SetFuncMap(template.FuncMap{
		"add": func(a int) int {
			a++
			return a
		},
	})

}
