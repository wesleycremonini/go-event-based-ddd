package main

import (
	"fmt"
	"net/http"

	"github.com/wesleycremonini/go-event-based-ddd/internal/response"
	"go.uber.org/zap"
)

func (app *application) status(w http.ResponseWriter, r *http.Request) {
	zap.L().Info("health-check")
	data := map[string]string{
		"status": "OK",
	}

	err := response.Success(w, http.StatusOK, "OK", data)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) notFound(w http.ResponseWriter, r *http.Request) {
	var data []string
	response.Error(w, http.StatusNotFound, "Not found", data)
}

func (app *application) methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("The %s method is not supported for this resource", r.Method)
	var data []string
	response.Error(w, http.StatusNotFound, message, data)
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	zap.L().Error(err.Error())
	message := "The server encountered a problem and could not process your request"
	var data []string
	response.Error(w, http.StatusInternalServerError, message, data)
}
