package main

import (
	"net/http"

	"github.com/alexedwards/flow"
	"github.com/wesleycremonini/go-event-based-ddd/internal/event"
	"github.com/wesleycremonini/go-event-based-ddd/internal/request"
	"github.com/wesleycremonini/go-event-based-ddd/internal/response"
	"github.com/wesleycremonini/go-event-based-ddd/internal/workforce"
)

func (app *application) handleEvent(w http.ResponseWriter, r *http.Request) {
	var DTO event.OrderEventDTO
	err := request.DecodeJSON(w, r, &DTO)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	DTO.Token = flow.Param(r.Context(), "token")

	app.eventQueue.Enqueue(workforce.Job{
		Service: app.eventService,
		Input:   DTO,
	})

	response.Success(w, http.StatusOK, "", nil)
}
