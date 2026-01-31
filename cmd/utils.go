package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

func loadTemplates(router *gin.Engine) error {
	tmpl, err := template.New("").ParseGlob("templates/*.html")
	if err != nil {
		return err
	}

	router.SetHTMLTemplate(tmpl)
	return nil
}
