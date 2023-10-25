package event

import (
	"context"
	"sync"
	"time"

	"github.com/wesleycremonini/go-event-based-ddd/internal/domain"
	"go.uber.org/zap"
)

var WG sync.WaitGroup

type orderEventService struct {
	companies domain.CompanyRepository
	customers domain.CustomerRepository
}

func NewOrderEventService(cfgs ...any) orderEventService {
	svc := orderEventService{}

	for _, cfg := range cfgs {
		switch c := cfg.(type) {
		case domain.CompanyRepository:
			svc.companies = c
		case domain.CustomerRepository:
			svc.customers = c
		}
	}

	return svc
}

func (svc orderEventService) Execute(input any) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	WG.Add(1)
	defer WG.Done()

	event, ok := input.(OrderEventDTO)
	if !ok {
		zap.L().Error("error parsing event")
		return
	}

	err := event.Validate()
	if err != nil {
		zap.L().Error("error validating event", zap.Error(err))
		return
	}

	_ = ctx
	// HANDLE THE EVENT
}
