package event

import (
	"errors"
)

type OrderEventDTO struct {
	Type    string `json:"type"`
	Token string `json:"-"`
}

func (o *OrderEventDTO) Validate() error {
	if o.Type != "EVENT" {
		return errors.New("invalid event type")
	}

	return nil
}
