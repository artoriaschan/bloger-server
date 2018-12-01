package controller

import (
	"html/template"
	"log"
	"net/http"
)

func IndexTemplate(writer http.ResponseWriter, request *http.Request) {
	indexTemplate, err := template.ParseFiles("./index.html")
	if err != nil {
		panic(err)
	}
	templateErr := indexTemplate.Execute(writer, nil)
	if templateErr != nil {
		panic(templateErr)
	}
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
}
